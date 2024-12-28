package models

import "gorm.io/gorm"

type NotificationType string

const (
	NotificationTypeReview   NotificationType = "review"
	NotificationTypeFavorite NotificationType = "favorite"
)

type Notification struct {
	gorm.Model
	IsNew    bool             `gorm:"not null"`
	Type     NotificationType `gorm:"type:notification_type;not null"`
	User     User             `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	Person   User             `gorm:"foreignKey:PersonID;constraint:OnDelete:CASCADE;"`
	Book     Book             `gorm:"foreignKey:BookID;constraint:OnDelete:CASCADE;"`
	UserID   uint
	BookID   uint
	PersonID uint
}
