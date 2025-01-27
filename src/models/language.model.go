package models

type ListLanguageRequest struct {
	Offset int `json:"offset"`
}

type CreateLanguageRequest struct {
	LanguageCode string `json:"languageCode" binding:"required"`
	Language     string `json:"language" binding:"required"`
}
