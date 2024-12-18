package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Price       string `gorm:"not null"`
	Image       string
	Author      string
	UserID      uint
}
