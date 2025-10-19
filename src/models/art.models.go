package models

type GetArtByIDRequest struct {
	UserID   int64  `json:"-"`
	UserRole string `json:"-"`
	TargetID int64  `json:"-"`
}

type ListArtRequest struct {
	UserID   int64  `json:"-"`
	UserRole string `json:"-"`
	Offset   int64  `json:"offset"`
}

type CreateArtRequest struct {
	UserID       int64    `json:"-"`
	UserRole     string   `json:"-"`
	Price        int64    `json:"price" binding:"required,min=0"`
	CurrencyID   int64    `json:"currency_id" binding:"required"`
	CreationYear int      `json:"creation_year" binding:"required,min=1000,max=9999"`
	Width        *float32 `json:"width,omitempty"`
	Height       *float32 `json:"height,omitempty"`
	Depth        *float32 `json:"depth,omitempty"`
	Available    *bool    `json:"available,omitempty"`
	Visible      *bool    `json:"visible,omitempty"`
}

type UpdateArtRequest struct {
	UserID       int64    `json:"-"`
	UserRole     string   `json:"-"`
	TargetID     int64    `json:"-"`
	Price        *int64   `json:"price,omitempty" binding:"omitempty,min=0"`
	CurrencyID   *int64   `json:"currency_id,omitempty"`
	CreationYear *int     `json:"creation_year,omitempty" binding:"omitempty,min=1000,max=9999"`
	Width        *float32 `json:"width,omitempty"`
	Height       *float32 `json:"height,omitempty"`
	Depth        *float32 `json:"depth,omitempty"`
	Available    *bool    `json:"available,omitempty"`
	Visible      *bool    `json:"visible,omitempty"`
}

type DeleteArtRequest struct {
	UserID   int64  `json:"-"`
	UserRole string `json:"-"`
	TargetID int64  `json:"-"`
}
