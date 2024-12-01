package requests

import "github.com/GiorgiTsukhishvili/BookShelf-Api/models"

type UserRegisterRequest struct {
	Name     string          `json:"name" binding:"required"`
	Email    string          `json:"email" binding:"required,email"`
	Password string          `json:"password" binding:"required"`
	Type     models.UserType `json:"type" binding:"required"`
}
