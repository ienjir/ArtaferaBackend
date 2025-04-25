package models

type GetRoleByIDRequest struct {
	UserID   int64  `json:"-"`
	UserRole string `json:"-"`
	RoleID   int64  `json:"-"`
}
type ListRoleRequest struct {
	UserID   int64  `json:"-"`
	UserRole string `json:"-"`
	Offset   int64  `json:"offset"`
}

type CreateRoleRequest struct {
	UserID   int64  `json:"-"`
	UserRole string `json:"-"`
	Role     string `json:"role" binding:"required"`
}

type UpdateRoleRequest struct {
	UserID   int64  `json:"-"`
	UserRole string `json:"-"`
	RoleID   int64  `json:"-"`
	Role     string `json:"role" binding:"required"`
}

type DeleteRoleRequest struct {
	UserID   int64  `json:"-"`
	UserRole string `json:"-"`
	RoleID   int64  `json:"-"`
}
