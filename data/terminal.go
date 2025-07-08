package data

import "github.com/google/uuid"

type TerminalRequest struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `validate:"required,min=2,max=100" json:"name"`
	Location string    `validate:"required,min=2,max=255" json:"location"`
	Status   string    `validate:"required,min=2,max=50" json:"status"`
}

type TerminalResponse struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Location string    `json:"location"`
	Status   string    `json:"status"`
}
