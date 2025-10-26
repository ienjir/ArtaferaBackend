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
	UserID        int64                `json:"-" form:"-"`
	UserRole      string               `json:"-" form:"-"`
	Name          *string              `json:"-" form:"name"`
	IsPublic      *bool                `form:"isPublic" json:"-"`
	Picture       multipart.FileHeader `json:"-" form:"-"`
	PublicBucket  string               `json:"-" form:"-"`
	PrivateBucket string               `json:"-" form:"-"`
}

type UpdatePictureRequest struct {
	UserID        int64   `json:"-"`
	UserRole      string  `json:"-"`
	TargetID      int64   `json:"-"`
	Name          *string `json:"name"`
	Priority      *int64  `json:"priority"`
	IsPublic      *bool   `json:"isPublic"`
	PublicBucket  string  `json:"-"`
	PrivateBucket string  `json:"-"`
}

type DeletePictureRequest struct {
	UserID        int64  `json:"-"`
	UserRole      string `json:"-"`
	TargetID      int64  `json:"-"`
	PublicBucket  string `json:"-"`
	PrivateBucket string `json:"-"`
}
