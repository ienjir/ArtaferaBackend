package models

import "mime/multipart"

type CreateArtPicture struct {
	TargetUserID *int64  `json:"userID"`
	UserRole     string  `json:"-"`
	ArtID        int64   `json:"artID"`
	Name         *string `json:"name"`
	Image        multipart.FileHeader
}
