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

type CreateArtTranslationRequest struct {
	UserID       int64  `json:"-"`
	UserRole     string `json:"-"`
	ArtID        int64  `json:"artID"`
	LanguageCode string `json:"languageCode"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Text         string `json:"text"`
}
