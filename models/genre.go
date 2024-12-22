package models

import "gorm.io/gorm"

type Genre struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Books []Book `gorm:"many2many:book_genres;"`
}
