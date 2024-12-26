package requests

type ReviewGetRequest struct {
	Page string `form:"page" binding:"required"`
	Size string `form:"size" binding:"required"`
}

type ReviewPostRequest struct {
	Rating  string `json:"rating" binding:"required"`
	Comment string `json:"comment" binding:"required"`
	BookID  string `json:"book_id" binding:"required"`
}

type ReviewPutRequest struct {
	ID      string `json:"id" binding:"required"`
	Rating  string `json:"rating" binding:"required"`
	Comment string `json:"comment" binding:"required"`
	BookID  string `json:"book_id" binding:"required"`
}
