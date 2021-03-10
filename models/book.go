package models

type CreateBookModel struct {
	ID     uint   `json:"id" binding:"required"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
