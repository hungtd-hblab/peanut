package usecase

import "peanut/repository"

type ContentUsecase interface {
}

type contentUsecase struct {
	ContentRepo repository.ContentRepo
}

func NewContentUsecase(r repository.ContentRepo) ContentUsecase {
	return &contentUsecase{
		ContentRepo: r,
	}
}
