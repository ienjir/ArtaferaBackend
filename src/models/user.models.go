package models

type GetUserByIDRequest struct {
	UserID   int64  `json:"-"`
	UserRole string `json:"-"`
	TargetID int64  `json:"-"`
}

type GetUserByEmailRequest struct {
	UserID   int64  `json:"-"`
	UserRole string `json:"-"`
	Email    string `json:"email"`
}

type ListUserRequest struct {
	UserID   int64  `json:"-"`
	UserRole string `json:"-"`
	Offset   int64  `json:"offset"`
}

type CreateUserRequest struct {
	Firstname   string  `json:"firstname" binding:"required"`
	Lastname    string  `json:"lastname" binding:"required"`
	Email       string  `json:"email" binding:"required,email"`
	Phone       *string `json:"phone"`
	PhoneRegion *string `json:"phone_region"`
	Address1    *string `json:"address1"`
	Address2    *string `json:"address2"`
	City        *string `json:"city"`
	PostalCode  *string `json:"postal_code"`
	Password    string  `json:"password" binding:"required"`
}

type UpdateUserRequest struct {
	UserID      int64   `json:"-"`
	UserRole    string  `json:"-"`
	TargetID    int64   `json:"-"`
	Firstname   *string `json:"firstname"`
	Lastname    *string `json:"lastname"`
	Email       *string `json:"email"`
	Phone       *string `json:"phone"`
	PhoneRegion *string `json:"phone_region"`
	Address1    *string `json:"address1"`
	Address2    *string `json:"address2"`
	City        *string `json:"city"`
	PostalCode  *string `json:"postal_code"`
	Password    *string `json:"password"`
	RoleID      *int64  `json:"roleID"`
}

type DeleteUserRequest struct {
	UserID   int64  `json:"-"`
	UserRole string `json:"-"`
	TargetID int64  `json:"-"`
}
