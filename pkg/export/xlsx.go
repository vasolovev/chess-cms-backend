package export

import (
	"fmt"
	"v1/internal/entity"

	"github.com/xuri/excelize/v2"
)

func ExportToXLSXformat(tournaments []entity.Tournament) error {
	// Создаем Excel файл
	xlsx := excelize.NewFile()

	size := tournaments[0].NbPlayers * 10
	excel := make([][]interface{}, size)
	for i := range excel {
		excel[i] = make([]interface{}, len(tournaments)+6)
	}

	// Лист 1 ФИО, группа, никнейм и кол-во турниров
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

					if tournaments[id].Perf.Key == "blitz" && count >= 5 {
						// excel[последняя запись в таблице][номер турнира по счету + номер стартовой колонки + 1]
						excel[i][id+1] = "+"
					} else if (tournaments[id].Perf.Key == "classical" || tournaments[id].Perf.Key == "rapid") && count >= 3 {
						excel[i][id+1] = "+"
					} else if tournaments[id].Perf.Key != "classical" && tournaments[id].Perf.Key != "rapid" && tournaments[id].Perf.Key != "blitz" {
						excel[i][id+1] = "?"
					}

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

				if tournaments[id].Perf.Key == "blitz" && count >= 5 {
					// excel[последняя запись в таблице][номер турнира по счету + номер стартовой колонки + 1]
					excel[lastRow][id+1] = "+"
				} else if (tournaments[id].Perf.Key == "classical" || tournaments[id].Perf.Key == "rapid") && count >= 3 {
					excel[lastRow][id+1] = "+"
				} else if tournaments[id].Perf.Key != "classical" && tournaments[id].Perf.Key != "rapid" && tournaments[id].Perf.Key != "blitz" {
					excel[lastRow][id+1] = "?"
				}

				lastRow += 1
			}

		}
	}

	// Добавляем ФИО игроков
	excel, err := importStudents(excel)
	if err != nil {
		return err
	}

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

	// Лист 2 ФИО, группа, никнейм и игры по дням недели

	// Последний индекс записанного никнейма
	lastRow = 1
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

	// Добавляем ФИО игроков
	excel, err = importStudents(excel)
	if err != nil {
		return err
	}

	sheet = "Sheet2"
	index = xlsx.NewSheet(sheet)

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

	// Save spreadsheet by the given path.
	if err := xlsx.SaveAs("Book1.xlsx"); err != nil {
		return err
	}

	return nil
}

func importStudents(excel [][]interface{}) ([][]interface{}, error) {
	// Открываем Excel файл
	f, err := excelize.OpenFile("Students.xlsx")
	if err != nil {
		return nil, err

	}

	rows, err := f.GetRows("Ответы на форму (1)")
	if err != nil {
		return nil, err

	}

	// В цикле
	for indexStudent := 1; indexStudent < len(rows); indexStudent++ {

		for indexTournament := 1; indexTournament < len(excel); indexTournament++ {

			if rows[indexStudent][5] == excel[indexTournament][0] {
				excel[indexTournament][len(excel[0])-1] = rows[indexStudent][4]
				excel[indexTournament][len(excel[0])-2] = rows[indexStudent][3]
				excel[indexTournament][len(excel[0])-3] = rows[indexStudent][2]
				excel[indexTournament][len(excel[0])-4] = rows[indexStudent][1]
			}
		}

	}
	return excel, nil
}
