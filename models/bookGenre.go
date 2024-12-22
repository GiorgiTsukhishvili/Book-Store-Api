package models

type BookGenre struct {
	BookID  int `gorm:"primaryKey"`
	GenreID int `gorm:"primaryKey"`
}
