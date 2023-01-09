package controller

import (
	"peanut/domain"
	"peanut/pkg/response"
	"peanut/repository"
	"peanut/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ContentController struct {
	Usecase usecase.ContentUsecase
}

func NewContentController(db *gorm.DB) *ContentController {
	return &ContentController{
		Usecase: usecase.NewContentUsecase(repository.NewContentRepo(db)),
	}
}

func (c *ContentController) CreateContent(ctx *gin.Context) {
	req := domain.CreateContentReq{}
	if !bind(ctx, &req) {
		return
	}

	resp, err := c.Usecase.CreateContent(ctx, req)
	if checkError(ctx, err) {
		return
	}

	response.OK(ctx, resp)
}
