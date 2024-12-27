package requests

import "mime/multipart"

type BookGetRequest struct {
	Page    string `form:"page" binding:"required"`
	Size    string `form:"size" binding:"required"`
	Keyword string `form:"keyword"`
}

type BookPostRequest struct {
	Name        string                `from:"name" binding:"required"`
	Description string                `from:"description" binding:"required"`
	Image       *multipart.FileHeader `from:"image" binding:"required"`
	Price       string                `from:"price" binding:"required"`
	AuthorID    uint                  `from:"author_id" binding:"required"`
	GenreIDs    []uint                `from:"genre_ids" binding:"required"`
}

type BookPutRequest struct {
	ID          uint                  `from:"id" binding:"required"`
	Name        string                `from:"name" binding:"required"`
	Description string                `from:"description" binding:"required"`
	Image       *multipart.FileHeader `from:"image" binding:"required"`
	ImagePath   string                `from:"image_path" binding:"required"`
	Price       string                `from:"price" binding:"required"`
	AuthorID    uint                  `from:"author_id" binding:"required"`
	GenreIDs    []uint                `from:"genre_ids" binding:"required"`
}
