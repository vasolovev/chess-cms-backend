package usecase

import (
	"context"
	"fmt"
	"v1/internal/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserUseCase -.
type UserUseCase struct {
	repo UserRepo
}

// New -.
func NewUserUseCase(r UserRepo) *UserUseCase {
	return &UserUseCase{
		repo: r,
	}
}

func (uc *UserUseCase) Add(ctx context.Context, user entity.User) error {
	err := uc.repo.Create(ctx, user)
	if err != nil {
		return fmt.Errorf("UserUseCase - Add - uc.repo.Create: %w", err)
	}
	return err
}
func (uc *UserUseCase) GetByID(ctx context.Context, id primitive.ObjectID) (entity.User, error) {
	user, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return entity.User{}, fmt.Errorf("UserUseCase - GetByID - uc.repo.GetByID: %w", err)
	}
	return user, nil
}
func (uc *UserUseCase) GetAll(ctx context.Context) ([]entity.User, error) {
	users, err := uc.repo.GetAll(ctx)
	if err != nil {
		return []entity.User{}, fmt.Errorf("UserUseCase - GetByID - uc.repo.GetByID: %w", err)
	}
	return users, nil
}
func (uc *UserUseCase) Delete(ctx context.Context, id primitive.ObjectID) error {
	err := uc.repo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("UserUseCase - Delete -  uc.repo.Delete: %w", err)
	}
	return nil
}
