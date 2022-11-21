package v1

import (
	"fmt"
	"net/http"
	"v1/internal/usecase"
	"v1/pkg/logger"

	"github.com/gin-gonic/gin"
)

type tournamentRoutes struct {
	t usecase.Tournament
	l logger.Interface
}

func newTournamentRoutes(handler *gin.RouterGroup, t usecase.Tournament, l logger.Interface) {
	r := &tournamentRoutes{t, l}

	h := handler.Group("/")
	{
		h.GET("/getAll", r.getAll)
		h.GET("/getByID", r.getByID)
		h.POST("/add", r.add)
	}
}
func (r *tournamentRoutes) getAll(c *gin.Context) {
	tournaments, err := r.t.GetAll(c.Request.Context())

	if err != nil {
		fmt.Errorf("TournamentHttp - getAll - r.t.GetAll: %w", err)
		c.String(http.StatusBadRequest, "")
		return
	}
	c.JSON(http.StatusOK, tournaments)
}
func (r *tournamentRoutes) add(c *gin.Context) {
	id := c.Query("id")
	err := r.t.Add(c.Request.Context(), id)

	if err != nil {
		fmt.Errorf("TournamentHttp - getAll - r.t.Add: %w", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failure",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
func (r *tournamentRoutes) getByID(c *gin.Context) {
	// Если в запросе
	id := c.Query("id")
	tournament, err := r.t.GetByID(c.Request.Context(), id)

	if err != nil {
		fmt.Errorf("TournamentHttp - getAll - r.t.GetAll: %w", err)
		c.String(http.StatusBadRequest, "")
		return
	}
	if tournament.NbPlayers == 0 {
		c.String(http.StatusNotFound, "null")
		return
	}
	c.JSON(http.StatusOK, tournament)
}
