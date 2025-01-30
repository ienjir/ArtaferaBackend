package models

type CreateOrderRequest struct {
	UserID   *int64 `json:"userID"`
	ArtID    int64  `json:"artID"`
	UserRole string `json:"-"`
	AuthID   int64  `json:"-"`
}

type GetOrderByIDRequest struct {
	UserID   int64  `json:"-"`
	OrderID  int64  `json:"-"`
	UserRole string `json:"-"`
}

type GetOrdersForUser struct {
	TargetUserID int64  `json:"-"`
	UserID       int64  `json:"-"`
	UserRole     string `json:"-"`
	Offset       int64  `json:"offset"`
}
