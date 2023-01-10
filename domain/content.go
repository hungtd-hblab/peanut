package domain

import (
	"mime/multipart"
	"time"
)

type Content struct {
	BaseModel
	Title       string
	Description string
	FilePath    string
	PlayTime    time.Time
	Resolution  int
	AspectRatio string
	Tag         string
	Category    string
}

type CreateContentReq struct {
	Title       string                `form:"title" binding:"required"`
	Description string                `form:"description" binding:"required"`
	PlayTime    string                `form:"play_time" binding:"required,datetime=2006-01-02 15:04:05"`
	Resolution  int                   `form:"resolution"`
	AspectRatio string                `form:"aspect_ratio"`
	Tag         string                `form:"tag" binding:"required"`
	Category    string                `form:"category" binding:"required"`
	File        *multipart.FileHeader `form:"file" binding:"required"`
}

type CreateContentResp struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PlayTime    time.Time `json:"play_time"`
	Resolution  int       `json:"resolution"`
	AspectRatio string    `json:"aspect_ratio"`
	Tag         string    `json:"tag"`
	Category    string    `json:"category"`
	File        string    `json:"file"`
}
