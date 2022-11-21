package usecase

import (
	"context"
	"fmt"
	"v1/internal/entity"

	"github.com/xuri/excelize/v2"
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
	uc.repo.Create(ctx, tournament)

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

func (uc *TournamentUseCase) ExportToXlsx(ctx context.Context, tournaments []entity.Tournament) error {
	// Создаем Excel файл
	xlsx := excelize.NewFile()
	sheet := "Sheet1"
	// Индекс листа в Excel
	index := xlsx.NewSheet(sheet)
	// Индекс строки в Excel
	indexRow := 2

	// Перебираем все турниры
	for id := 0; id < len(tournaments); id++ {

		var addr string
		var err error

		// Записываем в каждый столбец дату турнира
		// Узнаем номер ячейки
		if addr, err = excelize.CoordinatesToCellName(id+2, 1); err != nil {
			fmt.Println(err)
		}

		// Записываем в ячейку дату турнира
		if err = xlsx.SetCellValue(sheet, addr, tournaments[id].StartsAt.String()); err != nil {
			fmt.Println(err)
		}

		// Перебираем всех игроков в цикле
		for idPlayer := 0; idPlayer < len(tournaments[id].Standing.Players); idPlayer++ {
			//Поиск уже записанных игроков
			rows, err := xlsx.GetRows("Sheet1")

			if err != nil {
				fmt.Println(err)
			}

			find := false
			indexPlayer := 2

			// Ищем есть ли никнейм игрока среди уже записанных в Excel
			for i := 0; i < len(rows) && !find; i++ {
				if rows[i][0] == tournaments[id].Standing.Players[idPlayer].Name {
					find = true
					indexPlayer = i + 1
				}
			}

			// Если не нашли никнейм, то записываем его в Excel
			if !find {
				// Узнаем номер ячейки
				if addr, err = excelize.CoordinatesToCellName(1, indexRow); err != nil {
					fmt.Println(err)
				}

				//Записываем в ячейку дату турнира
				if err = xlsx.SetCellValue(sheet, addr, tournaments[id].Standing.Players[idPlayer].Name); err != nil {
					fmt.Println(err)
				}
				indexPlayer = indexRow
				indexRow += 1
			}

			// Записываем посещение
			// Узнаем номер ячейки
			if addr, err = excelize.CoordinatesToCellName(2+id, indexPlayer); err != nil {
				fmt.Println(err)
			}

			// Тест
			if val, err := xlsx.GetCellValue(sheet, addr); val != "" {
				fmt.Println(val)
				fmt.Println(err)
			}

			// Записываем в ячейку количество шахматных партий за турнир
			if err = xlsx.SetCellValue(sheet, addr, len(tournaments[id].Standing.Players[idPlayer].Sheet.Scores)); err != nil {
				fmt.Println(err)
			}

		}
	}

	// Set active sheet of the workbook.
	xlsx.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := xlsx.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}

	return nil
}
