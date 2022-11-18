package routes

import (
	"clean-architecture/api/responses"
	"clean-architecture/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestRoutes struct {
	handler infrastructure.Router
}

func NewTestRoutes(
	handler infrastructure.Router,
) *TestRoutes {
	return &TestRoutes{
		handler: handler,
	}
}

// Setup test routes
func (t *TestRoutes) Setup() {
	t.handler.GET("/test", func(ctx *gin.Context) {
		responses.JSON(ctx, http.StatusOK, gin.H{"message": "Test route working!"})
	})

}
