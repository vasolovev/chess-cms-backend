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
	tournamentUseCase := usecase.New(
		repo.New(db),
		webapi.New(),
	)

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l, tournamentUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

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
