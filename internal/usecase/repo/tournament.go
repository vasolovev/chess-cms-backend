package repo

import (
	"context"
	"fmt"

	"github.com/vasolovev/ChessCMS/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TournamentRepo -.
type TournamentRepo struct {
	db *mongo.Collection
}

// New -.
func NewTournamentRepo(db *mongo.Database) *TournamentRepo {
	return &TournamentRepo{
		db: db.Collection(tournamentsCollection),
	}
}

// Сохранение в базу данных сведений о турнире
func (r *TournamentRepo) Create(ctx context.Context, tournament entity.Tournament) error {
	_, err := r.db.InsertOne(ctx, tournament)
	if err != nil {
		return fmt.Errorf("TournamentRepo - Store - r.db.InsertOne: %w", err)
	}
	return err
}
func (r *TournamentRepo) GetByID(ctx context.Context, id string) (entity.Tournament, error) {
	var tournament entity.Tournament

	err := r.db.FindOne(ctx, bson.M{"_id": id}).Decode(&tournament)
	if err != nil {
		return entity.Tournament{}, fmt.Errorf("TournamentRepo - GetByID - r.db.FindOne: %w", err)
	}

	return tournament, err
}

func (r *TournamentRepo) GetAll(ctx context.Context) ([]entity.Tournament, error) {
	// Получение всех данных о турнирах в базе данных
	var tournaments []entity.Tournament
	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{"startAt", 1}})
	result, err := r.db.Find(ctx, filter, opts)
	if err != nil {
		return []entity.Tournament{}, fmt.Errorf("TournamentRepo - GetAll - r.db.Find: %w", err)
	}

	// Декодирование результата запроса в массив
	err = result.All(ctx, &tournaments)
	if err != nil {
		return []entity.Tournament{}, fmt.Errorf("TournamentRepo - GetAll - result.All: %w", err)
	}
	return tournaments, err
}
func (r *TournamentRepo) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.db.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("TournamentRepo - Delete - r.db.DeleteOne: %w", err)
	}
	return err
}
