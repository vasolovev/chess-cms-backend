package amqprpc

import (
	"github.com/vasolovev/ChessCMS/internal/usecase"
	"github.com/vasolovev/ChessCMS/pkg/rabbitmq/rmq_rpc/server"
)

// NewRouter -.
func NewRouter(t usecase.Tournament) map[string]server.CallHandler {
	routes := make(map[string]server.CallHandler)
	{
		newTournamentRoutes(routes, t)
	}

	return routes
}
