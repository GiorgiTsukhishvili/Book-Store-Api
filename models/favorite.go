package models

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	BookID uint
	UserID uint
	Book   Book `gorm:"foreignKey:BookID"`
	User   User `gorm:"foreignKey:UserID"`
}
