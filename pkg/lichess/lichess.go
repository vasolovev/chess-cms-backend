package lichess

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/vasolovev/ChessCMS/internal/entity"
)

func GetInfoAboutTournament(tournamentID string) (entity.Tournament, error) {
	resp, err := http.Get("https://lichess.org/api/tournament/" + tournamentID)
	if err != nil {
		return entity.Tournament{}, fmt.Errorf("Lichess - GetInfoAboutTournament - http.Get: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return entity.Tournament{}, fmt.Errorf("Lichess - GetInfoAboutTournament - io.ReadAll(resp.Body): %w", err)
	}

	var tournament entity.Tournament
	err = json.Unmarshal(body, &tournament)
	if err != nil {
		return entity.Tournament{}, fmt.Errorf("Lichess - GetInfoAboutTournament - json.Unmarshal: %w", err)
	}

	count := tournament.NbPlayers / 10
	if tournament.NbPlayers%10 != 0 {
		count = count + 1
	}
	for i := 2; i <= count; i++ {
		resp, err := http.Get("https://lichess.org/api/tournament/" + tournamentID + "?page=" + strconv.Itoa(i))
		if err != nil {
			return entity.Tournament{}, fmt.Errorf("Lichess - GetInfoAboutTournament - http.Get: %w", err)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return entity.Tournament{}, fmt.Errorf("Lichess - GetInfoAboutTournament - io.ReadAll: %w", err)
		}

		var pageTournament entity.Tournament
		err = json.Unmarshal(body, &pageTournament)
		if err != nil {
			return entity.Tournament{}, fmt.Errorf("Lichess - GetInfoAboutTournament - json.Unmarshal: %w", err)
		}
		tournament.Standing.Players = append(tournament.Standing.Players, pageTournament.Standing.Players...)
	}
	return tournament, nil
}
