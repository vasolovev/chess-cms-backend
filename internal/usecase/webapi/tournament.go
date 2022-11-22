package webapi

import (
	"fmt"

	"github.com/vasolovev/ChessCMS/internal/entity"
	"github.com/vasolovev/ChessCMS/pkg/lichess"

	"golang.org/x/net/context"
)

// TournamentWebAPI -.
type TournamentWebAPI struct {
}

// New -.
func New() *TournamentWebAPI {

	return &TournamentWebAPI{}
}

func (t *TournamentWebAPI) RequestInfoTournament(ctx context.Context, id string) (entity.Tournament, error) {
	resp, err := lichess.GetInfoAboutTournament(id)
	if err != nil {
		return entity.Tournament{}, fmt.Errorf("TournamentWebAPI - RequestInfoTournament - lichess.GetInfoAboutTournament: %w", err)
	}
	return resp, nil
}
