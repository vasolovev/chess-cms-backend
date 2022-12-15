package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/vasolovev/ChessCMS/internal/usecase"
	"github.com/vasolovev/ChessCMS/pkg/logger"
)

type userRoutes struct {
	t usecase.User
	l logger.Interface
}

func newUserRoutes(handler *gin.RouterGroup, t usecase.User, l logger.Interface) {
	r := &userRoutes{t, l}

	h := handler.Group("/")
	{
		// h.GET("/getAll", r)
		// h.GET("/getByID", r.getByID)
		h.POST("/add", r.add)
	}
}

func (r *userRoutes) add(c *gin.Context) {
	// user := c.Query("id")
	// user, err := r.t.Create(c.Request.Context(), user)

	// if err != nil {
	// 	fmt.Errorf("TournamentHttp - getAll - r.t.Add: %w", err)
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"status": "failure",
	// 	})
	// 	return
	// }

	return
}
