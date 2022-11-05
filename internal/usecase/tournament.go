package usecase

import (
	"context"
	"fmt"
	"v1/internal/entity"
)

// TournamentUseCase -.
type TournamentUseCase struct {
	repo   TournamentRepo
	webAPI TournamentWebAPI
}

// New -.
func New(r TournamentRepo, w TournamentWebAPI) *TournamentUseCase {
	return &TournamentUseCase{
		repo:   r,
		webAPI: w,
	}
}

func (uc *TournamentUseCase) Add(ctx context.Context, id string) error {
	// Проверяем есть ли запись о турнире в БД
	_, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		if err.Error() != "TournamentRepo - GetByID - r.db.FindOne: mongo: no documents in result" {
			return fmt.Errorf("TournamentUseCase - Add - repo.GetByID: %w", err)
		}

	}

	// Если записи нет, то
	// Запрашиваем у Lichess сведения о турнире
	tournament, err := uc.webAPI.RequestInfoTournament(ctx, id)
	if err != nil {
		return fmt.Errorf("TournamentUseCase - Add - uc.webAPI.RequestInfoTournament: %w", err)
	}

	// Если турнир не завершен, то не записываем в базу данных
	if !tournament.IsFinished {
		return fmt.Errorf("TournamentUseCase - Add - !tournament.IsFinished: %w", err)
	}

	// Запись в базу данных результатов турниров, которые еще не завершены
	uc.repo.Store(ctx, tournament)

	return nil
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
	return tournaments, nil
}
