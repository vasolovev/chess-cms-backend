// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"
	"v1/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	Tournament interface {
		Add(context.Context, string) error
		GetByID(context.Context, string) (entity.Tournament, error)
		GetAll(context.Context) ([]entity.Tournament, error)
	}

	// TournamentRepo -.
	TournamentRepo interface {
		Store(context.Context, entity.Tournament) error
		GetByID(context.Context, string) (entity.Tournament, error)
		GetAll(context.Context) ([]entity.Tournament, error)
	}

	// TournamentWebAPI -.
	TournamentWebAPI interface {
		RequestInfoTournament(context.Context, string) (entity.Tournament, error)
	}
)
