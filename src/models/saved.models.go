package models

type GetSavedByIDRequest struct {
	UserID   int64  `json:"-"`
	UserRole string `json:"-"`
	TargetID int64  `json:"-"`
}

type GetSavedForUserRequest struct {
	UserID       int64  `json:"-"`
	UserRole     string `json:"-"`
	TargetUserID int64  `json:"-"`
	Offset       int64  `json:"offset"`
}
