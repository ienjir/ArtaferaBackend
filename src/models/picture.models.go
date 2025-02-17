package models

type CreatePictureRequest struct {
	Name     string `json:"name" binding:"required"`
	Priority *int   `json:"priority"`
	IsPublic bool   `json:"is_public"`
	Image    []byte `json:"-"`
}
