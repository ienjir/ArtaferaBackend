package models

type CreateOrderRequest struct {
	UserID   *int64 `json:"userID"`
	ArtID    int64  `json:"artID"`
	UserRole string `json:"-"`
}
