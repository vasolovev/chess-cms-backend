package export

import (
	"fmt"
	"v1/internal/entity"

	"github.com/xuri/excelize/v2"
)

func ExportToXLSXformat(tournaments []entity.Tournament) error {
	size := tournaments[0].NbPlayers * 10
	excel := make([][]interface{}, size)
	for i := range excel {
		excel[i] = make([]interface{}, len(tournaments)+1)
	}

	// Последний индекс записанного никнейма
	lastRow := 1
	// Перебор турниров
	for id := 0; id < len(tournaments); id++ {

		// Запись в самую первую строку даты турнира
		excel[0][id+1] = tournaments[id].StartsAt.String()

		// Перебираем всех игроков в цикле
		for idPlayer := 0; idPlayer < len(tournaments[id].Standing.Players); idPlayer++ {

			// Найден ли никнейм?
			find := false

			// Ищем есть ли никнейм игрока среди уже записанных
			for i := 1; excel[i][0] != nil && !find && i < size; i++ {
				if excel[i][0] == tournaments[id].Standing.Players[idPlayer].Name {
					find = true
				}

				if find == true {
					// Записываем в ячейку количество шахматных партий за турнир
					count := len(tournaments[id].Standing.Players[idPlayer].Sheet.Scores)
					excel[i][id+1] = count
					if i > lastRow {
						lastRow = i + 1
					}

				}

			}

			// Если не нашли никнейм, то записываем его
			if !find {
				excel[lastRow][0] = tournaments[id].Standing.Players[idPlayer].Name

				// Записываем в ячейку количество шахматных партий за турнир
				count := len(tournaments[id].Standing.Players[idPlayer].Sheet.Scores)
				excel[lastRow][id+1] = count

				lastRow += 1
			}

		}
	}

	// Создаем Excel файл
	xlsx := excelize.NewFile()
	sheet := "Sheet1"
	index := xlsx.NewSheet(sheet)

	for r, row := range excel {
		for c, col := range row {
			var addr string
			var err error

			// Создаем индекс ячейки Excel (B1 например)
			if addr, err = excelize.CoordinatesToCellName(c+1, r+1); err != nil {
				fmt.Println(err)
			}

			//Записываем в ячейку
			if err = xlsx.SetCellValue(sheet, addr, col); err != nil {
				fmt.Println(err)
			}
		}
		if r == lastRow {
			break
		}
	}

	// Set active sheet of the workbook.
	xlsx.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := xlsx.SaveAs("Book1.xlsx"); err != nil {
		return err
	}

	return nil
}
