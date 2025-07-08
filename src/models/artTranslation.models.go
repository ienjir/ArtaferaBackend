package models

type GetArtTranslationByIDRequest struct {
	UserID   int64  `json:"-"`
	UserRole string `json:"-"`
	TargetID int64  `json:"-"`
}

type ListArtTranslationRequest struct {
	UserID   int64  `json:"-"`
	UserRole string `json:"-"`
	Offset   int    `json:"offset"`
}
