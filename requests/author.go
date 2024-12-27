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
	Image       *multipart.FileHeader `form:"image" binding:"required"`
	Nationality string                `form:"nationality" binding:"required"`
}

type AuthorPutRequest struct {
	ID          string    `json:"id" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	BirthDate   time.Time `json:"birth_date" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Image       string    `json:"image" binding:"required"`
	Nationality string    `json:"nationality" binding:"required"`
}
