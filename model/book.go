package model

type Book struct {
	ID     uint    `json:"id" gorm:"primaryKey"`
	Title  string  `json:"title" gorm:"not null"`
	Author string  `json:"author" gorm:"not null"`
}