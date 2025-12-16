package v1

type ErrorResponse struct {
	Title  string `json:"title" example:"Conflict"`
	Status int    `json:"status" example:"409"`
	Detail string `json:"detail" example:"Owner already exists"`
}
