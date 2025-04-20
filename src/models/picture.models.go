package models

import "mime/multipart"

type GetPictureByIDRequest struct {
	UserID        int64  `json:"-"`
	UserRole      string `json:"-"`
	TargetID      int64  `json:"-"`
	PublicBucket  string `json:"-"`
	PrivateBucket string `json:"-"`
}

type GetPictureByNameRequest struct {
	UserID        int64  `json:"-"`
	UserRole      string `json:"-"`
	Name          string `json:"name"`
	PublicBucket  string `json:"-"`
	PrivateBucket string `json:"-"`
}

type ListPictureRequest struct {
	UserID        int64  `json:"-"`
	UserRole      string `json:"-"`
	Offset        int64  `jsoǹ:"offset"`
	PublicBucket  string `json:"-"`
	PrivateBucket string `json:"-"`
}

type CreatePictureRequest struct {
	UserID        int64                `json:"-"`
	UserRole      string               `json:"-"`
	Name          *string              `json:"name" form:"name"`
	Priority      *int64               `json:"priority" from:"priority"`
	IsPublic      *bool                `json:"isPublic;default:false" form:"isPublic"`
	Picture       multipart.FileHeader `json:"-"`
	PublicBucket  string               `json:"-"`
	PrivateBucket string               `json:"-"`
}
