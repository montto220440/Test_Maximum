package dto

type ListRequest struct {
	Name string `json:"name" binding:"required"`
}

type ListResponse struct {
	ID uint `json:"id"`
	Name string `json:"name"`
}