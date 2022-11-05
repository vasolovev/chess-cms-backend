package repo

import (
	"context"
	"fmt"
	"v1/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// TournamentRepo -.
type TournamentRepo struct {
	db *mongo.Collection
}

// New -.
func New(db *mongo.Database) *TournamentRepo {
	return &TournamentRepo{
		db: db.Collection(tournamentsCollection),
	}
}

// Сохранение в базу данных сведений о турнире
func (r *TournamentRepo) Store(ctx context.Context, tournament entity.Tournament) error {
	_, err := r.db.InsertOne(ctx, tournament)
	if err != nil {
		return fmt.Errorf("TournamentRepo - Store - r.db.InsertOne: %w", err)
	}
	return err
}
func (r *TournamentRepo) GetByID(ctx context.Context, id string) (entity.Tournament, error) {
	var tournament entity.Tournament

	err := r.db.FindOne(ctx, bson.M{"id": id}).Decode(&tournament)
	if err != nil {
		return entity.Tournament{}, fmt.Errorf("TournamentRepo - GetByID - r.db.FindOne: %w", err)
	}

	return tournament, err
}

func (r *TournamentRepo) GetAll(ctx context.Context) ([]entity.Tournament, error) {
	// Получение всех данных о турнирах в базе данных
	var tournaments []entity.Tournament
	result, err := r.db.Find(ctx, bson.M{})
	if err != nil {
		return []entity.Tournament{}, fmt.Errorf("TournamentRepo - GetAll - r.db.Find: %w", err)
	}

	// Декоидрование результата запроса в массив
	err = result.All(ctx, &tournaments)
	if err != nil {
		return []entity.Tournament{}, fmt.Errorf("TournamentRepo - GetAll - result.All: %w", err)
	}
	return tournaments, err
}
