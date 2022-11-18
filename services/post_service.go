package services

import (
	"clean-architecture/lib"
	"clean-architecture/models"
	"clean-architecture/repository"

	"github.com/jinzhu/copier"
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

func (ps PostService) GetOnePost(id uint) (post models.Post, err error) {
	return ps.repository.GetOne(id)
}

func (ps PostService) GetAllPosts() ([]models.Post, error) {
	return ps.repository.FindAll()
}

func (ps PostService) SavePost(post models.Post) error {
	_, err := ps.repository.Save(post)
	return err
}

func (ps PostService) UpdatePost(id uint, post models.Post) error {
	postDB, err := ps.repository.GetOne(id)
	if err != nil {
		return err
	}

	copier.Copy(&postDB, &post)

	post.ID = id

	_, err = ps.repository.Update(post)

	return err
}

func (ps PostService) DeletePost(id uint) error {
	return ps.repository.Delete(id)
}
