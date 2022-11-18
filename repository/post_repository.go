package repository

import (
	"clean-architecture/infrastructure"

	"gorm.io/gorm"
)

type PostRepository struct {
	infrastructure.Database
}

func NewPostRepository(db infrastructure.Database) PostRepository {
	return PostRepository{db}
}

func (r PostRepository) WithTrx(trxHandle *gorm.DB) PostRepository {
	r.Database.DB = trxHandle
	return r
}
