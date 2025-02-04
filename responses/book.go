package responses

import "github.com/GiorgiTsukhishvili/BookShelf-Api/models"

type BookRetrieveResponse struct {
	Book models.Book `json:"book"`
}

type BooksGetResponse struct {
	Data       []models.Book      `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}
