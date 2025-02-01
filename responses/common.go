package responses

type PaginationResponse struct {
	CurrentPage string `json:"current_page"`
	FistPage    int    `json:"fist_page"`
	LastPage    int    `json:"last_page"`
	Total       int64  `json:"total"`
}
