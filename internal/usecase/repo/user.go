package repo

import (
	"context"
	"fmt"
	"v1/internal/entity"

	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepo -.
type UserRepo struct {
	db *mongo.Collection
}

// New -.
func NewUserRepo(db *mongo.Database) *UserRepo {
	return &UserRepo{
		db: db.Collection(userCollection),
	}
}

// Сохранение в базу данных сведений о турнире
func (r *UserRepo) Create(ctx context.Context, user entity.User) error {
	_, err := r.db.InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("TournamentRepo - Store - r.db.InsertOne: %w", err)
	}
	return err
}
