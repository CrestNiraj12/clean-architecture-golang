package routes

import (
	"clean-architecture/api/controllers"
	"clean-architecture/infrastructure"
	"clean-architecture/lib"
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
		api.GET("/posts", pr.postController.GetPosts)
		api.GET("/posts/:id", pr.postController.GetOnePost)
		api.POST("/posts", pr.postController.SavePost)
		api.PUT("/posts/:id", pr.postController.UpdatePost)
		api.DELETE("/posts/:id", pr.postController.DeletePost)
	}
}
