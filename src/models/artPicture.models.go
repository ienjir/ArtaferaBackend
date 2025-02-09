package models

type CreateArtPictureRequest struct {
	UserID    int64  `json:"-"`
	UserRole  string `json:"-"`
	ArtID     int64  `json:"artID"`
	ImageName string `json:"imageName"`
}
