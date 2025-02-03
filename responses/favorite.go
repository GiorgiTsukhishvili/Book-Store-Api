package responses

import "github.com/GiorgiTsukhishvili/BookShelf-Api/models"

type FavoritePostResponse struct {
	Favorite models.Favorite `json:"favorite"`
}

type FavoritesGetResponse struct {
	Data       []models.Favorite  `json:"data"`
	Pagination PaginationResponse `json:"pagination"`
}
