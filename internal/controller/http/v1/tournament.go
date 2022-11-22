package v1

import (
	"fmt"
	"net/http"

	"github.com/vasolovev/ChessCMS/internal/usecase"
	"github.com/vasolovev/ChessCMS/pkg/logger"

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
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failure",
		})
		return
	}
	c.JSON(http.StatusOK, tournaments)
	return
}
func (r *tournamentRoutes) add(c *gin.Context) {
	id := c.Query("id")
	tournament, err := r.t.Add(c.Request.Context(), id)

	if err != nil {
		fmt.Errorf("TournamentHttp - getAll - r.t.Add: %w", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failure",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         tournament.ID,
		"name":       tournament.FullName,
		"createdBy":  tournament.CreatedBy,
		"startsAt":   tournament.StartsAt,
		"isFinished": tournament.IsFinished,
	})
	return
}
func (r *tournamentRoutes) getByID(c *gin.Context) {
	// Если в запросе
	id := c.Query("id")
	tournament, err := r.t.GetByID(c.Request.Context(), id)

	if err != nil {
		fmt.Errorf("TournamentHttp - getAll - r.t.GetAll: %w", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failure",
		})
		return
	}
	if tournament.NbPlayers == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failure",
		})
		return
	}
	c.JSON(http.StatusOK, tournament)
	return
}
