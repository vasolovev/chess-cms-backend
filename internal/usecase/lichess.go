package usecase

import (
	"context"
	"v1/internal/entity"
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

func (uc *LichessUseCase) Create(ctx context.Context, la entity.Lichess) error {
	return nil
}
