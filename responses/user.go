package responses

import "time"

type UserMe struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Image       string    `json:"image"`
	Type        string    `json:"type"`
	CreatedAt   time.Time `json:"created_at"`
}

type MeResponse struct {
	User UserMe `json:"user"`
}
