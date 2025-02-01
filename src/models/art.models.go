package models

type GetArtByIDRequest struct {
	UserID   int64  `json:"-"`
	UserRole string `json:"-"`
	TargetID int64  `json:"-"`
}
