package models

import "gorm.io/gorm"

type UserType string

const (
	UserTypeUser     UserType = "user"
	UserTypeAdmin    UserType = "admin"
	UserTypeBusiness UserType = "business"
)

type User struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Email       string `gorm:"unique;not null"`
	PhoneNumber string `gorm:"unique;not null"`
	Password    string `gorm:"not null"`
	Image       string
	Type        UserType   `gorm:"type:user_type;not null"`
	Books       []Book     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	Reviews     []Review   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	Favorites   []Favorite `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}
