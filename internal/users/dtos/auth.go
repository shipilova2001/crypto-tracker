package dtos

import "time"

type AuthJSON struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Auth struct {
	Username string
	Password string
}

type UserResponse struct {
	ID 		 int
	Username string
	CreatedAt time.Time
	UpdatedAt time.Time
}