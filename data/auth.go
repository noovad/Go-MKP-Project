package data

import "github.com/google/uuid"

type RegisterRequest struct {
	Username string `validate:"required,min=3,max=50" json:"username"`
	Password string `validate:"required,min=6" json:"password"`
}

type LoginRequest struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}

type AuthResponse struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Token    string    `json:"token"`
}

type UserResponse struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}
