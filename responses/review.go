package responses

import "github.com/GiorgiTsukhishvili/BookShelf-Api/models"

type ReviewRetrieveResponse struct {
	Review models.Review `json:"review"`
}

type ReviewsGetResponse struct {
	Data       []models.Review    `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}
