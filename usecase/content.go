package usecase

import (
	"context"
	"fmt"
	"peanut/config"
	"peanut/domain"
	"peanut/pkg/apierrors"
	"peanut/pkg/array"
	"peanut/repository"
	"strings"
	"time"
)

type ContentUsecase interface {
	CreateContent(ctx context.Context, req domain.CreateContentReq, filePath string) (domain.CreateContentResp, error)
	GetContents(ctx context.Context, req domain.GetContentsReq) (*domain.GetContentsResp, error)
}

type contentUsecase struct {
	ContentRepo repository.ContentRepo
}

func NewContentUsecase(r repository.ContentRepo) ContentUsecase {
	return &contentUsecase{
		ContentRepo: r,
	}
}

func (uc *contentUsecase) CreateContent(ctx context.Context, req domain.CreateContentReq, filePath string) (resp domain.CreateContentResp, err error) {
	content := domain.Content{
		Title:       req.Title,
		Description: req.Description,
		Resolution:  req.Resolution,
		AspectRatio: req.AspectRatio,
		Tag:         req.Tag,
		Category:    req.Category,
	}

	content.FilePath = filePath
	content.PlayTime, err = time.Parse("2006-01-02 15:04:05", req.PlayTime)
	if err != nil {
		err = apierrors.NewErrorf(apierrors.InternalError, err.Error())
		return
	}

	cnt, err := uc.ContentRepo.CreateContent(content)
	if err != nil {
		err = fmt.Errorf("[usecase.ContentUsecase.CreateContent] failed: %w", err)
		return
	}

	resp = domain.CreateContentResp{
		ID:          cnt.ID,
		Title:       cnt.Title,
		Description: cnt.Description,
		Resolution:  cnt.Resolution,
		AspectRatio: cnt.AspectRatio,
		Tag:         cnt.Tag,
		Category:    cnt.Category,
		File:        cnt.FilePath,
	}

	return
}

func (uc *contentUsecase) GetContents(ctx context.Context, req domain.GetContentsReq) (resp *domain.GetContentsResp, err error) {
	uc.setDefaultFilter(&req)

	contents, err := uc.ContentRepo.GetContents(req)
	if err != nil {
		err = fmt.Errorf("[usecase.ContentUsecase.GetContents] failed: %w", err)
		return
	}

	resp = &domain.GetContentsResp{
		Contents: contents,
	}

	return
}

func (uc *contentUsecase) setDefaultFilter(req *domain.GetContentsReq) {
	if req.Limit == 0 {
		req.Limit = config.DefaultLimit
	}
	if req.Offset == 0 {
		req.Offset = config.DefaultOffset
	}
	sortedCol := []string{"id", "title", "updated_at", "created_at"}
	if !array.IsIn(req.SortBy, sortedCol) {
		req.SortBy = config.DefaultSortBy
	}

	if strings.ToUpper(req.SortType) != "ASC" {
		req.SortType = config.DefaultSortType
	}
}
