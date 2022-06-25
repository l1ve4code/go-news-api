package entities

type News struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Comment []Comment `json:"comments" gorm:"constraint:OnDelete:CASCADE;"`
}