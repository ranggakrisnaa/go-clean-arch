package entity

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email,omitempty"`
	FullName  string    `json:"full_name,omitempty"`
	Phone     string    `json:"phone_number,omitempty"`
	Password  string    `json:"password,omitempty"`
	ImageURL  string    `json:"image_url,omitempty"`
	CreatedAt string    `json:"created_at,omitempty"`
	UpdatedAt string    `json:"updated_at,omitempty"`
}
