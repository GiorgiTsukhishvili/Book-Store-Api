package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name         string    `gorm:"not null"`
	Description  string    `gorm:"not null"`
	Price        string    `gorm:"not null"`
	Image        string    `gorm:"not null"`
	CreationDate time.Time `gorm:"type:date"`
	AuthorID     uint
	UserID       uint
	Author       Author     `gorm:"foreignKey:AuthorID"`
	User         User       `gorm:"foreignKey:UserID"`
	Genres       []Genre    `gorm:"many2many:book_genres;"`
	Reviews      []Review   `gorm:"foreignKey:BookID;constraint:OnDelete:CASCADE;"`
	Favorites    []Favorite `gorm:"foreignKey:BookID;constraint:OnDelete:CASCADE;"`
}
