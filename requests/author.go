package requests

import "time"

type AuthorGetRequest struct {
	Page string `form:"page" binding:"required"`
	Size string `form:"size" binding:"required"`
}

type AuthorPostRequest struct {
	Name        string    `json:"name" binding:"required"`
	BirthDate   time.Time `json:"birth_date" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Image       string    `json:"image" binding:"required"`
	Nationality string    `json:"nationality" binding:"required"`
}

type AuthorPutRequest struct {
	ID          string    `json:"id" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	BirthDate   time.Time `json:"birth_date" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Image       string    `json:"image" binding:"required"`
	Nationality string    `json:"nationality" binding:"required"`
}
