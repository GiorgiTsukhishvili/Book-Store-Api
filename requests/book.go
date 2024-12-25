package requests

type BookGetRequest struct {
	Page    string `form:"page" binding:"required"`
	Size    string `form:"size" binding:"required"`
	Keyword string `form:"keyword"`
}

type BookPostRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Image       string `json:"image" binding:"required"`
	Price       string `json:"price" binding:"required"`
	AuthorID    uint   `json:"author_id" binding:"required"`
	GenreIDs    []uint `json:"genre_ids" binding:"required"`
}

type BookPutRequest struct {
	ID          uint   `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Image       string `json:"image" binding:"required"`
	Price       string `json:"price" binding:"required"`
	AuthorID    uint   `json:"author_id" binding:"required"`
	GenreIDs    []uint `json:"genre_ids" binding:"required"`
}
