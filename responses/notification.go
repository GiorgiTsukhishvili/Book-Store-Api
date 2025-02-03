package responses

import "github.com/GiorgiTsukhishvili/BookShelf-Api/models"

type NotificationsGetResponse struct {
	Data       []models.Notification `json:"data"`
	Pagination PaginationResponse    `json:"pagination"`
}
