package entities

type News struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Comment []Comment `json:"-" gorm:"constraint:OnDelete:CASCADE;"`
}