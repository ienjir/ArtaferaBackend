package models

type ListRoleRequest struct {
	Offset int `json:"offset"`
}

type CreateRoleRequest struct {
	Role string `json:"role" binding:"required"`
}

type UpdateRoleRequest struct {
	RoleID string `json:"-"`
	Role   string `json:"role" binding:"required"`
}
