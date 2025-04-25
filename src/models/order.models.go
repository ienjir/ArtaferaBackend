package models

type GetOrderByIDRequest struct {
	UserID   int64  `json:"-"`
	OrderID  int64  `json:"-"`
	UserRole string `json:"-"`
}

type GetOrdersForUserRequest struct {
	UserID       int64  `json:"-"`
	UserRole     string `json:"-"`
	TargetUserID int64  `json:"-"`
	Offset       int64  `json:"offset"`
}

type ListOrdersRequest struct {
	UserID   int64  `json:"-"`
	UserRole string `json:"-"`
	Offset   int64  `jsoǹ:"offset"`
}

type CreateOrderRequest struct {
	TargetUserID *int64 `json:"userID"`
	UserRole     string `json:"-"`
	ArtID        int64  `json:"artID"`
	UserID       int64  `json:"-"`
}

type UpdateOrderRequest struct {
	UserID       int64   `json:"-"`
	UserRole     string  `json:"-"`
	TargetID     int64   `json:"-"`
	TargetUserID *int64  `json:"userID"`
	ArtID        *int64  `json:"artID"`
	Status       *string `json:"status"`
}

type DeleteOrderRequest struct {
	UserID   int64  `json:"-"`
	UserRole string `json:"-"`
	TargetID int64  `json:"-"`
}
