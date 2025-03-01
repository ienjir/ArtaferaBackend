package models

import "mime/multipart"

type CreatePictureRequest struct {
	Name       string               `json:"name" binding:"required" form:"name"`
	Priority   *int                 `json:"priority"`
	IsPublic   bool                 `json:"isPublic;default:false"`
	Picture    multipart.FileHeader `json:"-"`
	BucketName string               `json:"-"`
}
