package routes

import (
	"clean-architecture/api/controllers"
	"clean-architecture/api/responses"
	"clean-architecture/infrastructure"
	"clean-architecture/lib"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostRoutes struct {
	logger         lib.Logger
	handler        infrastructure.Router
	postController controllers.PostController
}

func NewPostRoutes(
	handler infrastructure.Router,
	logger lib.Logger,
	postController controllers.PostController,
) *PostRoutes {
	return &PostRoutes{
		handler:        handler,
		logger:         logger,
		postController: postController,
	}
}

func (pr PostRoutes) Setup() {
	pr.logger.Info("Setting up post routes")

	api := pr.handler.Group("/")
	{
		api.GET("/posts/:id", func(ctx *gin.Context) {
			responses.JSON(ctx, http.StatusOK, "Api to get post")
		})

		api.POST("/posts", func(ctx *gin.Context) {
			responses.JSON(ctx, http.StatusOK, "Api to create post")
		})
	}
}
