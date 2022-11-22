package usecase

import (
	"context"
	"fmt"
	"v1/internal/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// LichessUseCase -.
type LichessUseCase struct {
	repo LichessRepo
}

// New -.
func NewLichessUseCase(r LichessRepo) *LichessUseCase {
	return &LichessUseCase{
		repo: r,
	}
}

func (uc *LichessUseCase) Create(ctx context.Context, lichess entity.Lichess) (primitive.ObjectID, error) {
	id, err := uc.repo.Create(ctx, lichess)
	if err != nil {
		return primitive.ObjectID{}, fmt.Errorf("UserUseCase - Add - uc.repo.Create: %w", err)
	}
	return id, err
}
func (uc *LichessUseCase) GetAll(ctx context.Context) ([]entity.Lichess, error) {
	lichess, err := uc.repo.GetAll(ctx)
	if err != nil {
		return []entity.Lichess{}, fmt.Errorf("LichessUseCase - GetAll - uc.repo.GetAll: %w", err)
	}

	return lichess, nil
}
func (uc *LichessUseCase) GetByID(ctx context.Context, id primitive.ObjectID) (entity.Lichess, error) {
	lichess, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return entity.Lichess{}, fmt.Errorf("LichessUseCase - GetByID - uc.repo.GetByID: %w", err)
	}

	return lichess, nil
}
func (uc *LichessUseCase) Update(ctx context.Context, lichess entity.Lichess) error {
	err := uc.repo.Update(ctx, lichess)
	if err != nil {
		return fmt.Errorf("LichessUseCase - GetByID - uc.repo.GetByID: %w", err)
	}
	return nil
}
func (uc *LichessUseCase) Delete(ctx context.Context, id primitive.ObjectID) error {
	err := uc.repo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("LichessUseCase - Delete -  uc.repo.Delete: %w", err)
	}
	return nil
}
