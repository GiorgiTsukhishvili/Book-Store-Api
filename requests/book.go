package requests

import (
	"mime/multipart"
	"time"
)

type BookGetRequest struct {
	Page    string `form:"page" binding:"required"`
	Size    string `form:"size" binding:"required"`
	Keyword string `form:"keyword"`
}

type BookPostRequest struct {
	Name         string                `form:"name" binding:"required"`
	Description  string                `form:"description" binding:"required"`
	Image        *multipart.FileHeader `form:"image"`
	ImagePath    string                `form:"image_path"`
	CreationDate time.Time             `form:"creation_date" binding:"required"`
	Price        string                `form:"price" binding:"required"`
	AuthorID     uint                  `form:"author_id" binding:"required"`
	GenreIDs     []uint                `form:"genre_ids" binding:"required"`
}

type BookPutRequest struct {
	ID           uint                  `form:"id" binding:"required"`
	Name         string                `form:"name" binding:"required"`
	Description  string                `form:"description" binding:"required"`
	Image        *multipart.FileHeader `form:"image"`
	ImagePath    string                `form:"image_path"`
	CreationDate time.Time             `form:"creation_date" binding:"required"`
	Price        string                `form:"price" binding:"required"`
	AuthorID     uint                  `form:"author_id" binding:"required"`
	GenreIDs     []uint                `form:"genre_ids" binding:"required"`
}
