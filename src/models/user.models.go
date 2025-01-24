package models

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

type ListUserRequest struct {
	Offset int `json:"offset"`
}

type GetUserByEmail struct {
	Email       string  `json:"email"`
	RequestID   float64 `json:"-"`
	RequestRole string  `json:"-"`
}
