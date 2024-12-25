package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	Rating  string `gorm:"not null"`
	Comment string `gorm:"not null"`
	Book    Book   `gorm:"foreignKey:BookID"`
	User    User   `gorm:"foreignKey:UserID"`
	UserID  uint
	BookID  uint
}
