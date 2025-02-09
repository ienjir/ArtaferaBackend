package models

type CreatePictureRequest struct {
	UserID    int64  `json:"-"`
	UserRole  string `json:"-"`
	ImageName string `json:"imageName"`
	Priority  *int   `json:"priority,omitempty"`
}
