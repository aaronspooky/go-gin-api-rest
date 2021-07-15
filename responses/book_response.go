package responses

type BookResponse struct {
	ID     uint           `json:"id" binding:"required"`
	Name   string         `json:"name" binding:"required"`
	Author AuthorResponse `json:"author"`
}

//
type AuthorResponse struct {
	ID   uint   `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
