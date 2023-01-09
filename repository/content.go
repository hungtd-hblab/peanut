package repository

import "gorm.io/gorm"

type ContentRepo interface {
	CreateContent()
}

type contentRepo struct {
	DB *gorm.DB
}

func NewContentRepo(db *gorm.DB) ContentRepo {
	return &contentRepo{
		DB: db,
	}
}

func (r *contentRepo) CreateContent() {

}
