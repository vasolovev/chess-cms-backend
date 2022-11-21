package usecase

import (
	"context"
	"fmt"
	"v1/internal/entity"
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
		return fmt.Errorf("TournamentRepo - Store - r.db.InsertOne: %w", err)
	}
	return err
}
func (uc *UserUseCase) GetByID(context.Context, string) (entity.User, error) {
	return entity.User{}, nil
}
func (uc *UserUseCase) GetAll(context.Context) ([]entity.User, error) {
	return []entity.User{}, nil
}
