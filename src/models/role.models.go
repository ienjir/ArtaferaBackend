package models

type ListRoleRequest struct {
	Offset int `json:"offset"`
}

type CreateRoleRequest struct {
	Role string `json:"role" binding:"required"`
}
