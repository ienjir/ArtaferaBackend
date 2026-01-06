package models

type CreateContactMessageRequest struct {
	Email   string `json:"email" binding:"required,email"`
	Message string `json:"message" binding:"required"`
}
