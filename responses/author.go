package responses

import "github.com/GiorgiTsukhishvili/BookShelf-Api/models"

type AuthorGetResponse struct {
	Author models.Author `json:"author"`
}

type AuthorsGetResponse struct {
	Data       []models.Author    `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}
