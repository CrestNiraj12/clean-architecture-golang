package repository

import (
	"clean-architecture/infrastructure"
	"clean-architecture/models"

	"gorm.io/gorm"
)

type PostRepository struct {
	db infrastructure.Database
}

func NewPostRepository(db infrastructure.Database) PostRepository {
	return PostRepository{db: db}
}

func (r PostRepository) WithTrx(trxHandle *gorm.DB) PostRepository {
	r.db.DB = trxHandle
	return r
}

func (pr PostRepository) FindAll() (posts []models.Post, err error) {
	return posts, pr.db.DB.Find(&posts).Error
}

func (pr PostRepository) GetOne(id uint) (post models.Post, err error) {
	return post, pr.db.DB.Where("id = ?", id).First(&post).Error
}

func (pr PostRepository) Save(post models.Post) (models.Post, error) {
	return post, pr.db.DB.Create(&post).Error
}

func (pr PostRepository) Update(post models.Post) (models.Post, error) {
	return post, pr.db.DB.Save(&post).Error
}

func (pr PostRepository) Delete(id uint) error {
	return pr.db.DB.Where("id = ?", id).Delete(&models.Post{}).Error
}
