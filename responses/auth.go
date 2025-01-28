package responses

import (
	"github.com/GiorgiTsukhishvili/BookShelf-Api/models"
	"github.com/GiorgiTsukhishvili/BookShelf-Api/utils"
)

type LoginResponse struct {
	JWT  utils.JWTInfo `json:"jwt"`
	User models.User   `json:"user"`
}

type RefreshTokenResponse struct {
	JWT utils.JWTInfo `json:"jwt"`
}

type LogoutResponse struct {
	Message string `json:"message"`
}
