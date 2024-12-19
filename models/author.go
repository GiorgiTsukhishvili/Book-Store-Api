package models

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name        string    `gorm:"not null"`
	BirthDate   time.Time `gorm:"type:date"`
	Description string    `gorm:"not null"`
	Image       string    `gorm:"not null"`
	Nationality string    `gorm:"not null"`
	Book        []Book    `gorm:"constraint:OnDelete:SET NULL;"`
}
