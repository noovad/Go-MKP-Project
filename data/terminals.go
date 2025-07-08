package data

type TerminalRequest struct {
	Name     string `validate:"required,min=2,max=100" json:"name"`
	Location string `validate:"required,min=2,max=255" json:"location"`
	Status   string `validate:"required,min=2,max=50" json:"status"`
}

type TerminalResponse struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Status   string `json:"status"`
}
