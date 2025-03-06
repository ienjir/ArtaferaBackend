package models

import "mime/multipart"

type CreatePictureRequest struct {
	Name       *string              `json:"name" form:"name"`
	Priority   *int64               `json:"priority" from:"priority"`
	IsPublic   *bool                `json:"isPublic;default:false" form:"isPublic"`
	Picture    multipart.FileHeader `json:"-"`
	BucketName string               `json:"-"`
}
