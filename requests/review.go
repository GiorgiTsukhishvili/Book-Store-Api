package requests

type ReviewGetRequest struct {
	Page string `form:"page" binding:"required"`
	Size string `form:"size" binding:"required"`
}

type ReviewPostRequest struct {
}

type ReviewPutRequest struct {
}
