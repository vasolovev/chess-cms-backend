package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"v1/config"
	v1 "v1/internal/controller/http/v1"
	"v1/internal/usecase"
	"v1/internal/usecase/repo"
	"v1/internal/usecase/webapi"
	"v1/pkg/httpserver"
	"v1/pkg/logger"
	"v1/pkg/mongodb"

	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)
	// Repository
	mongoClient, err := mongodb.NewClient(cfg.Mongo.URL, cfg.Mongo.User, cfg.Mongo.Password)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - mongodb.NewClient: %w", err))

		return
	}

	db := mongoClient.Database(cfg.Mongo.Name)

	// Use case
	tournamentUseCase := usecase.NewTournamentUseCase(
		repo.NewTournamentRepo(db),
		webapi.New(),
	)
	userUseCase := usecase.NewUserUseCase(
		repo.NewUserRepo(db),
	)
	lichessAccountUseCase := usecase.NewLichessUseCase(
		repo.NewLichessRepo(db),
	)

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l, tournamentUseCase, userUseCase, lichessAccountUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// gRPC Server
	// go func() {
	// 	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.GRPC.IP, cfg.GRPC.Port))
	// 	if err != nil {
	// 		l.Error(fmt.Errorf("app - Run - gRPC.Listen: %w", err))
	// 	}

	// 	serverOptions := []grpc.ServerOption{}

	// 	grpcServer := grpc.NewServer(serverOptions...)

	// 	err = grpcServer.Serve(listener)
	// }()

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

	if err := mongoClient.Disconnect(context.Background()); err != nil {
		l.Error(fmt.Errorf("app - Run - mongoClient.Disconnect: %w", err))
	}
}
