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
		Create(context.Context, entity.Tournament) error
		GetByID(context.Context, string) (entity.Tournament, error)
		GetAll(context.Context) ([]entity.Tournament, error)
	}

	// TournamentWebAPI -.
	TournamentWebAPI interface {
		RequestInfoTournament(context.Context, string) (entity.Tournament, error)
	}

	User interface {
		Add(context.Context, entity.User) error
		GetByID(context.Context, string) (entity.User, error)
		GetAll(context.Context) ([]entity.User, error)
	}

	UserRepo interface {
		Create(context.Context, entity.User) error
	}
	Lichess interface {
		Create(context.Context, entity.Lichess) error
	}
	LichessRepo interface {
		Create(context.Context, entity.Lichess) error
		GetAll(context.Context, entity.Lichess) error
		Update(context.Context, entity.Lichess) error
		Delete(context.Context, entity.Lichess) error
	}
)
