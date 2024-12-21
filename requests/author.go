package requests

type AuthorGetRequest struct {
	Page string `json:"page" binding:"required"`
	Size string `json:"size" binding:"required"`
}
