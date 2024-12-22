package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	Rating  string `gorm:"not null"`
	Comment string `gorm:"not null"`
	UserID  uint
	BookID  uint
}
