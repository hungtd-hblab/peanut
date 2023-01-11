package repository

import (
	"fmt"
	"peanut/domain"

	"gorm.io/gorm"
)

type ContentRepo interface {
	CreateContent(cnt domain.Content) (*domain.Content, error)
	GetContents(filter domain.GetContentsReq) ([]domain.Content, error)
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
		err = fmt.Errorf("[repo.Content.CreateContent] failed: %w", result.Error)
		return nil, err
	}

	return &cnt, nil
}

func (r *contentRepo) GetContents(filter domain.GetContentsReq) (contents []domain.Content, err error) {
	query := r.DB.Debug().
		Where("title LIKE ?", fmt.Sprintf("%%%v%%", filter.Title)).
		Where("category = ?", filter.Category)

	if filter.Tag != nil {
		query.Where("tag IN ?", filter.Tag)
	}

	query = query.Limit(filter.Limit).
		Offset(filter.Offset).
		Order(fmt.Sprintf("%s %s", filter.SortBy, filter.SortType))

	result := query.Find(&contents)

	if result.Error != nil {
		err = fmt.Errorf("[repo.Content.GetContents] failed: %w", result.Error)
		return nil, err
	}

	return
}
