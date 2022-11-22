package usecase

import (
	"context"
	"fmt"
	"v1/internal/entity"
	"v1/pkg/export"
)

// TournamentUseCase -.
type TournamentUseCase struct {
	repo   TournamentRepo
	webAPI TournamentWebAPI
}

// New -.
func NewTournamentUseCase(r TournamentRepo, w TournamentWebAPI) *TournamentUseCase {
	return &TournamentUseCase{
		repo:   r,
		webAPI: w,
	}
}

func (uc *TournamentUseCase) Add(ctx context.Context, id string) (entity.Tournament, error) {
	// Проверяем есть ли запись о турнире в БД
	res, err := uc.repo.GetByID(ctx, id)
	if err == nil {
		return res, fmt.Errorf("TournamentUseCase - Add - repo.GetByID: %w", err)
	}

	// Если записи нет, то
	// Запрашиваем у Lichess сведения о турнире
	tournament, err := uc.webAPI.RequestInfoTournament(ctx, id)
	if err != nil {
		return entity.Tournament{}, fmt.Errorf("TournamentUseCase - Add - uc.webAPI.RequestInfoTournament: %w", err)
	}

	// Если турнир не завершен, то не записываем в базу данных
	if !tournament.IsFinished {
		return entity.Tournament{}, fmt.Errorf("TournamentUseCase - Add - !tournament.IsFinished: %w", err)
	}

	// Запись в базу данных результатов турниров, которые еще не завершены
	err = uc.repo.Create(ctx, tournament)

	return tournament, err
}
func (uc *TournamentUseCase) GetByID(ctx context.Context, id string) (entity.Tournament, error) {
	tournament, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return entity.Tournament{}, fmt.Errorf("TournamentUseCase - GetByID - uc.repo.GetByID: %w", err)
	}

	return tournament, nil
}

func (uc *TournamentUseCase) GetAll(ctx context.Context) ([]entity.Tournament, error) {
	tournaments, err := uc.repo.GetAll(ctx)
	if err != nil {
		return []entity.Tournament{}, fmt.Errorf("TournamentUseCase - GetAll - uc.repo.GetAll: %w", err)
	}

	uc.ExportToXlsx(ctx, tournaments)
	return tournaments, nil
}

func (uc *TournamentUseCase) ExportToXlsx(ctx context.Context, tournaments []entity.Tournament) error {
	err := export.ExportToXLSXformat(tournaments)
	return err
}
