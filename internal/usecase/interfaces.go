// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/vasolovev/ChessCMS/internal/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	Tournament interface {
		Create(context.Context, string) (entity.Tournament, error)
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
		Create(context.Context, entity.User) error
		GetByID(context.Context, primitive.ObjectID) (entity.User, error)
		GetAll(context.Context) ([]entity.User, error)
		Delete(context.Context, primitive.ObjectID) error
	}

	UserRepo interface {
		Create(context.Context, entity.User) error
		GetAll(context.Context) ([]entity.User, error)
		GetByID(context.Context, primitive.ObjectID) (entity.User, error)
		Update(context.Context, entity.User) error
		Delete(context.Context, primitive.ObjectID) error
	}
	Lichess interface {
		Create(context.Context, entity.Lichess) (primitive.ObjectID, error)
		GetAll(context.Context) ([]entity.Lichess, error)
		GetByID(context.Context, primitive.ObjectID) (entity.Lichess, error)
		Update(context.Context, entity.Lichess) error
		Delete(context.Context, primitive.ObjectID) error
	}
	LichessRepo interface {
		Create(context.Context, entity.Lichess) (primitive.ObjectID, error)
		GetAll(context.Context) ([]entity.Lichess, error)
		GetByID(context.Context, primitive.ObjectID) (entity.Lichess, error)
		Update(context.Context, entity.Lichess) error
		Delete(context.Context, primitive.ObjectID) error
	}
)
