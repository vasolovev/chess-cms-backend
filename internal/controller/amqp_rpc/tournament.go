package amqprpc

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/vasolovev/ChessCMS/internal/entity"
	"github.com/vasolovev/ChessCMS/internal/usecase"
	"github.com/vasolovev/ChessCMS/pkg/rabbitmq/rmq_rpc/server"
)

type tournamentRoutes struct {
	tournamentUseCase usecase.Tournament
}

func newTournamentRoutes(routes map[string]server.CallHandler, t usecase.Tournament) {
	r := &tournamentRoutes{t}
	{
		routes["getHistory"] = r.getHistory()
	}
}

type historyResponse struct {
	History []entity.Tournament `json:"history"`
}

func (r *tournamentRoutes) getHistory() server.CallHandler {
	return func(d *amqp.Delivery) (interface{}, error) {
		tournaments, err := r.tournamentUseCase.GetAll(context.Background())
		if err != nil {
			return nil, fmt.Errorf("amqp_rpc - tournamentRoutes - getHistory - r.tournamentUseCase.History: %w", err)
		}

		response := historyResponse{tournaments}

		return response, nil
	}
}
