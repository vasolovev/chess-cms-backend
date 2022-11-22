package repo

import (
	"context"
	"fmt"
	"v1/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// LichessRepo -.
type LichessRepo struct {
	db *mongo.Collection
}

func NewLichessRepo(db *mongo.Database) *LichessRepo {
	return &LichessRepo{
		db: db.Collection(lichessCollection),
	}
}

func (r *LichessRepo) Create(ctx context.Context, lichess entity.Lichess) (primitive.ObjectID, error) {
	res, err := r.db.InsertOne(ctx, lichess)
	if err != nil {
		return primitive.ObjectID{}, fmt.Errorf("LichessRepo - Create - r.db.InsertOne: %w", err)
	}

	return res.InsertedID.(primitive.ObjectID), err
}

func (r *LichessRepo) GetAll(ctx context.Context) ([]entity.Lichess, error) {
	// Получение всех данных о турнирах в базе данных
	var players []entity.Lichess
	result, err := r.db.Find(ctx, bson.M{})
	if err != nil {
		return []entity.Lichess{}, fmt.Errorf("LichessRepo - GetAll - r.db.Find: %w", err)
	}

	// Декодирование результата запроса в массив
	err = result.All(ctx, &players)
	if err != nil {
		return []entity.Lichess{}, fmt.Errorf("LichessRepo - GetAll - result.All: %w", err)
	}
	return players, err
}

func (r *LichessRepo) GetByID(ctx context.Context, id primitive.ObjectID) (entity.Lichess, error) {
	var lichess entity.Lichess

	err := r.db.FindOne(ctx, bson.M{"_id": id}).Decode(&lichess)
	if err != nil {
		return entity.Lichess{}, fmt.Errorf("TournamentRepo - GetByID - r.db.FindOne: %w", err)
	}

	return lichess, err
}

func (r *LichessRepo) Update(ctx context.Context, lichess entity.Lichess) error {
	_, err := r.db.UpdateOne(ctx, bson.M{"_id": lichess.ID}, bson.M{"$set": bson.M{"ban": lichess.Ban}})
	if err != nil {
		return fmt.Errorf("LichessRepo - Update - r.db.UpdateOne: %w", err)
	}
	return err
}

func (r *LichessRepo) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.db.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("LichessRepo - Delete - r.db.DeleteOne: %w", err)
	}
	return err
}
