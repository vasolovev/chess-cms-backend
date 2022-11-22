package repo

import (
	"context"
	"fmt"

	"github.com/vasolovev/ChessCMS/internal/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r *UserRepo) GetAll(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	result, err := r.db.Find(ctx, bson.M{})
	if err != nil {
		return []entity.User{}, fmt.Errorf("TournamentRepo - GetAll - r.db.Find: %w", err)
	}

	// Декодирование результата запроса в массив
	err = result.All(ctx, &users)
	if err != nil {
		return []entity.User{}, fmt.Errorf("TournamentRepo - GetAll - result.All: %w", err)
	}
	return users, err
}

func (r *UserRepo) GetByID(ctx context.Context, id primitive.ObjectID) (entity.User, error) {
	var user entity.User

	err := r.db.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return entity.User{}, fmt.Errorf("UserRepo - GetByID - r.db.FindOne: %w", err)
	}

	return user, err
}

func (r *UserRepo) Update(ctx context.Context, user entity.User) error {
	updateQuery := bson.M{}

	if user.Name != "" {
		updateQuery["name"] = user.Name
	}

	if user.Surname != "" {
		updateQuery["surname"] = user.Surname
	}

	if user.Patronymic != "" {
		updateQuery["patronymic"] = user.Patronymic
	}

	if user.Email != "" {
		updateQuery["email"] = user.Email
	}

	if user.GroupNumber != "" {
		updateQuery["groupNumber"] = user.GroupNumber
	}
	if user.TelegramID != "" {
		updateQuery["telegramID"] = user.TelegramID
	}

	_, err := r.db.UpdateOne(ctx,
		bson.M{"_id": user.ID}, bson.M{"$set": updateQuery})

	return err
}

func (r *UserRepo) Delete(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.db.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("UserRepo - Delete - r.db.DeleteOne: %w", err)
	}
	return err
}
