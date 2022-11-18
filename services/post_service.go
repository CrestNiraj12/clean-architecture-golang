package services

import (
	"clean-architecture/lib"
	"clean-architecture/repository"
)

type PostService struct {
	logger     lib.Logger
	repository repository.PostRepository
}

func NewPostService(
	logger lib.Logger,
	repository repository.PostRepository,
) PostService {
	return PostService{
		logger:     logger,
		repository: repository,
	}
}
