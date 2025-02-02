package models

type GetLanguageByIDRequest struct {
	UserID   int64  `json:"-"`
	UserRole string `json:"-"`
	TargetID int64  `json:"-"`
}

type ListLanguageRequest struct {
	UserID   int64  `json:"-"`
	UserRole string `json:"-"`
	Offset   int    `json:"offset"`
}

type CreateLanguageRequest struct {
	UserID       int64  `json:"-"`
	UserRole     string `json:"-"`
	LanguageCode string `json:"languageCode" binding:"required"`
	Language     string `json:"language" binding:"required"`
}

type UpdateLanguageRequest struct {
	UserID       int64  `json:"-"`
	UserRole     string `json:"-"`
	TargetID     int64  `json:"-"`
	Language     string `json:"language"`
	LanguageCode string `json:"languageCode"`
}

type DeleteLanguageRequest struct {
	UserID   int64  `json:"-"`
	UserRole string `json:"-"`
	TargetID int64  `json:"-"`
}
