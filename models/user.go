package models

import "gorm.io/gorm"

type UserType string

const (
	UserTypeUser     UserType = "user"
	UserTypeBusiness UserType = "business"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Image    string
	Type     UserType `gorm:"type:user_type;not null"`
}
