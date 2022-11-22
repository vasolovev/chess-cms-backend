// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"v1/internal/usecase"
	"v1/pkg/logger"

	"github.com/gin-gonic/gin"
)

// NewRouter -.
// Swagger spec:
// @title       Go Clean Template API
// @description Using a translation service as an example
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(handler *gin.Engine, l logger.Interface, t usecase.Tournament, u usecase.User, li usecase.Lichess) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Routers
	h := handler.Group("/")
	{
		newTournamentRoutes(h, t, l)
	}
}
