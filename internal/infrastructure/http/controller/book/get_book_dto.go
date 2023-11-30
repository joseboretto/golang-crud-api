package book

type GetAllBookResponse struct {
	Title      string `json:"title"`
	TotalPages int    `json:"total_pages"`
	Isbn       string `json:"isbn"`
}
