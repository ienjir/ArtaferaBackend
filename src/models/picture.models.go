package models

import "mime/multipart"

type GetPictureByIDRequest struct {
	UserID     int64  `json:"-"`
	UserRole   string `json:"-"`
	TargetID   int64  `json:"-"`
	BucketName string `json:"-"`
}
type CreatePictureRequest struct {
	Name       *string              `json:"name" form:"name"`
	Priority   *int64               `json:"priority" from:"priority"`
	IsPublic   *bool                `json:"isPublic;default:false" form:"isPublic"`
	Picture    multipart.FileHeader `json:"-"`
	BucketName string               `json:"-"`
}
