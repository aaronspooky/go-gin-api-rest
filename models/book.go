package models

type CreateBookModel struct {
	ID        uint    `json:"id" binding:"required"`
	Title     string  `json:"title"`
	Author    string  `json:"author"`
	Sex       string  `json:"sex" binding:"required,oneof=male female"`
	Status    uint    `json:"status" binding:"required,numeric,min=1,max=6"`
	Birthdate *string `json:"birthdate"`
	IsActive  *bool   `json:"isActive" binding:"required"`
}
