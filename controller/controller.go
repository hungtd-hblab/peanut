package controller

import (
	"mime/multipart"
	"os"
	"path/filepath"
	"peanut/pkg/apierrors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func bindJSON(ctx *gin.Context, obj interface{}) bool {
	err := ctx.ShouldBindJSON(obj)
	if err == nil {
		return true
	}
	_, ok := err.(validator.ValidationErrors)
	if ok {
		err = apierrors.New(apierrors.InvalidRequest, err)
	} else {
		err = apierrors.New(apierrors.BadParams, err)
	}
	ctx.Error(err).SetType(gin.ErrorTypeBind)

	return false
}

func bindQueryParams(ctx *gin.Context, obj interface{}) bool {
	err := ctx.ShouldBindQuery(obj)

	if err == nil {
		return true
	}
	_, ok := err.(validator.ValidationErrors)
	if ok {
		err = apierrors.New(apierrors.InvalidRequest, err)
	} else {
		err = apierrors.New(apierrors.BadParams, err)
	}
	ctx.Error(err).SetType(gin.ErrorTypeBind)

	return false
}

func bind(ctx *gin.Context, obj interface{}) bool {
	err := ctx.ShouldBind(obj)
	if err == nil {
		return true
	}
	_, ok := err.(validator.ValidationErrors)
	if ok {
		err = apierrors.New(apierrors.InvalidRequest, err)
	} else {
		err = apierrors.New(apierrors.BadParams, err)
	}
	ctx.Error(err).SetType(gin.ErrorTypeBind)

	return false
}

func checkError(ctx *gin.Context, err error) bool {
	if err == nil {
		return false
	}
	_ = ctx.Error(err).SetType(gin.ErrorTypePublic)
	return true
}

func saveUploadedFile(ctx *gin.Context, file *multipart.FileHeader) (path string, err error) {
	// Create folder if not exist
	err = os.MkdirAll("./public/uploads", os.ModePerm)
	if err != nil {
		err = apierrors.NewErrorf(apierrors.InternalError, err.Error())
		return
	}

	fileId := uuid.New()
	fileName := fileId.String() + filepath.Ext(file.Filename)

	path = "./public/uploads/" + fileName
	err = ctx.SaveUploadedFile(file, path)

	return
}
