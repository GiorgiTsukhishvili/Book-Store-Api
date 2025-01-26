package responses

import "github.com/GiorgiTsukhishvili/BookShelf-Api/utils"

type User struct {
	Email       string `json:"email"`
	ID          int    `json:"id"`
	Image       string `json:"image"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Type        string `json:"type"`
}

type LoginResponse struct {
	JWT  utils.JWTInfo `json:"jwt"`
	User User          `json:"user"`
}

type RefreshTokenResponse struct {
	JWT utils.JWTInfo `json:"jwt"`
}
