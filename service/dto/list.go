package dto

type ListRequest struct {
	Content string `json:"content" binding:"required"`
}

type ListResponse struct {
	ID uint `json:"id"`
	Content string `json:"content"`
}