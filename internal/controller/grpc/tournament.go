package grpc

import (
	"context"
	"fmt"

	"github.com/vasolovev/ChessCMS/internal/usecase"
	"github.com/vasolovev/ChessCMS/pkg/logger"
	pb "github.com/vasolovev/ChessCMS/protobuf"
)

type tournamentRoutes struct {
	t usecase.Tournament
	l logger.Interface
}

// AddProduct реализует ecommerce.AddProduct
func (r *tournamentRoutes) AddTournament(ctx context.Context, protobuf pb.TournamentID) (*pb.Tournament, error) {
	tournament, err := r.t.Add(ctx, protobuf.Id)
	if err != nil {
		return &pb.Tournament{}, fmt.Errorf("tournamentRoutes - AddTournament - r.t.Add: %w", err)
	}
	return &pb.Tournament{Id: tournament.ID,
		Name:       tournament.FullName,
		CreatedBy:  tournament.CreatedBy,
		IsFinished: tournament.IsFinished,
		StartsAt:   tournament.StartsAt.Unix()}, nil
}
