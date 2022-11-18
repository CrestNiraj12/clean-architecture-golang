package controllers

import (
	"clean-architecture/lib"
	"clean-architecture/repository"
	"clean-architecture/services"
)

type PostController struct {
	service    services.PostService
	logger     lib.Logger
	repository repository.PostRepository
}

func NewPostController(
	postService services.PostService,
	logger lib.Logger,
	postRepository repository.PostRepository,
) PostController {
	return PostController{
		service:    postService,
		logger:     logger,
		repository: postRepository,
	}
}
