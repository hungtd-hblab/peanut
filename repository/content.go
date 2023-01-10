package repository

import (
	"fmt"
	"peanut/domain"

	"gorm.io/gorm"
)

type ContentRepo interface {
	CreateContent(cnt domain.Content) (*domain.Content, error)
}

type contentRepo struct {
	DB *gorm.DB
}

func NewContentRepo(db *gorm.DB) ContentRepo {
	return &contentRepo{
		DB: db,
	}
}

func (r *contentRepo) CreateContent(cnt domain.Content) (content *domain.Content, err error) {
	result := r.DB.Create(&cnt)
	if result.Error != nil {
		err = fmt.Errorf("[repo.User.CreateContent] failed: %w", result.Error)
		return nil, err
	}

	return &cnt, nil
}
