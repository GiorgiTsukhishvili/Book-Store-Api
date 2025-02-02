package requests

import (
	"mime/multipart"
	"time"
)

type AuthorGetRequest struct {
	Page string `form:"page" binding:"required"`
	Size string `form:"size" binding:"required"`
}

type AuthorPostRequest struct {
	Name        string                `form:"name" binding:"required"`
	BirthDate   time.Time             `form:"birth_date" binding:"required"`
	Description string                `form:"description" binding:"required"`
	Image       *multipart.FileHeader `form:"image"`
	ImagePath   string                `form:"image_path"`
	Nationality string                `form:"nationality" binding:"required"`
}

type AuthorPutRequest struct {
	ID          string                `form:"id" binding:"required"`
	Name        string                `form:"name" binding:"required"`
	BirthDate   time.Time             `form:"birth_date" binding:"required"`
	Description string                `form:"description" binding:"required"`
	Image       *multipart.FileHeader `form:"image"`
	ImagePath   string                `form:"image_path"`
	Nationality string                `form:"nationality" binding:"required"`
}
