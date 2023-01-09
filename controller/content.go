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

// SignUp godoc
//
//	@Summary		CreateContent
//	@Description	Create New Content
//	@Tags			content
//	@Accept			json
//	@Produce		json
//	@Param			param	formData	domain.CreateContentReq	true	"Create content request"
//	@Success		200	{object}	domain.Response{data=domain.CreateContentResp}
//	@Failure		400	{object}	domain.ErrorResponse
//	@Failure		500	{object}	domain.ErrorResponse
//	@Router			/v1/contents [post]
func (c *ContentController) CreateContent(ctx *gin.Context) {
	req := domain.CreateContentReq{}
	if !bind(ctx, &req) {
		return
	}

	// Handle uploaded file
	filePath, err := saveUploadedFile(ctx, req.File)
	if checkError(ctx, err) {
		return
	}

	resp, err := c.Usecase.CreateContent(ctx, req, filePath)
	if checkError(ctx, err) {
		return
	}

	response.OK(ctx, resp)
}
