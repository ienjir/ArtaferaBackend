package models

import (
	"gorm.io/gorm"
	"time"
)

var AllModels = []interface{}{
	&User{},
	&Role{},
	&Language{},
	&Saved{},
	&Order{},
	&Art{},
	&ArtTranslation{},
	&ArtPicture{},
	&Picture{},
	&Currency{},
	&Translation{},
}

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusPaid      OrderStatus = "paid"
	OrderStatusShipped   OrderStatus = "shipped"
	OrderStatusDelivered OrderStatus = "delivered"
	OrderStatusCancelled OrderStatus = "cancelled"
)

// SoftDeleteModel Base model with soft delete
type SoftDeleteModel struct {
	ID        int64          `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// Model Base model without soft delete
type Model struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// User Entities that should be soft deleted
type User struct {
	SoftDeleteModel
	Firstname   string     `gorm:"size:255;not null" json:"firstname" binding:"required"`
	Lastname    string     `gorm:"size:255;not null" json:"lastname" binding:"required"`
	Email       string     `gorm:"size:255;not null;uniqueIndex" json:"email" binding:"required,email"`
	Phone       *string    `gorm:"size:20" json:"phone,omitempty" binding:"omitempty,e164"`
	PhoneRegion *string    `gorm:"size:2" json:"phone_region,omitempty" binding:"omitempty,iso3166_1_alpha2"`
	Address1    *string    `gorm:"size:255" json:"address1,omitempty"`
	Address2    *string    `gorm:"size:255" json:"address2,omitempty"`
	City        *string    `gorm:"size:255" json:"city,omitempty"`
	PostalCode  *string    `gorm:"size:32" json:"postal_code,omitempty"`
	Password    []byte     `gorm:"type:bytea;not null" json:"-"`
	Salt        []byte     `gorm:"type:bytea;not null" json:"-"`
	LastAccess  *time.Time `json:"last_access,omitempty"`
	RoleID      int64      `gorm:"not null" json:"role_id"`
	Role        *Role      `gorm:"foreignKey:RoleID;references:ID" json:"role"`
	Orders      []Order    `gorm:"foreignKey:UserID" json:"orders,omitempty"`
	Saved       []Saved    `gorm:"foreignKey:UserID" json:"saved,omitempty"`
}

type Art struct {
	SoftDeleteModel
	Price        int64            `gorm:"not null;check:price >= 0" json:"price"`
	CurrencyID   int64            `gorm:"not null;index" json:"currency_id"`
	Currency     *Currency        `gorm:"foreignKey:CurrencyID;references:ID" json:"currency,omitempty"`
	CreationYear int              `gorm:"not null" json:"creation_year" binding:"required,min=1000,max=9999"`
	Width        *float32         `gorm:"type:decimal(8,2)" json:"width,omitempty"`
	Height       *float32         `gorm:"type:decimal(8,2)" json:"height,omitempty"`
	Depth        *float32         `gorm:"type:decimal(8,2)" json:"depth,omitempty"`
	Pictures     []ArtPicture     `gorm:"foreignKey:ArtID" json:"pictures,omitempty"`
	Translations []ArtTranslation `gorm:"foreignKey:ArtID" json:"translations,omitempty"`
	Orders       []Order          `gorm:"foreignKey:ArtID" json:"orders,omitempty"`
	Saved        []Saved          `gorm:"foreignKey:ArtID" json:"saved,omitempty"`
	Available    bool             `gorm:"default:true" json:"available"`
	Visible      bool             `gorm:"default:true" json:"-"`
}

// Role Entities that should be hard deleted
type Role struct {
	Model
	Name  string `gorm:"column:name;size:255;not null;uniqueIndex" json:"name"`
	Users []User `gorm:"foreignKey:RoleID" json:"users,omitempty"`
}

type Language struct {
	Model
	LanguageName    string           `gorm:"size:50;not null;unique" json:"language_name"`
	LanguageCode    string           `gorm:"size:2;not null;unique;index" json:"language_code"`
	Translations    []Translation    `gorm:"foreignKey:LanguageID" json:"translations,omitempty"`
	ArtTranslations []ArtTranslation `gorm:"foreignKey:LanguageID" json:"art_translations,omitempty"`
}

type Saved struct {
	Model
	UserID int64 `gorm:"not null;index" json:"user_id"`
	User   *User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
	ArtID  int64 `gorm:"not null;index" json:"art_id"`
	Art    *Art  `gorm:"foreignKey:ArtID;references:ID;constraint:OnDelete:CASCADE" json:"art,omitempty"`
}

type Order struct {
	Model
	UserID    int64       `gorm:"not null;index" json:"user_id"`
	User      *User       `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE" json:"user,omitempty"`
	ArtID     int64       `gorm:"not null;index" json:"art_id"`
	Art       *Art        `gorm:"foreignKey:ArtID;references:ID;constraint:OnDelete:CASCADE" json:"art,omitempty"`
	OrderDate time.Time   `gorm:"not null;default:CURRENT_TIMESTAMP" json:"order_date"`
	Status    OrderStatus `gorm:"size:50;not null;default:'pending'" json:"status"`
}

type ArtTranslation struct {
	Model
	ArtID       int64     `gorm:"not null;index" json:"art_id"`
	Art         *Art      `gorm:"foreignKey:ArtID;references:ID;constraint:OnDelete:CASCADE" json:"art,omitempty"`
	LanguageID  int64     `gorm:"not null;index" json:"language_id"`
	Language    *Language `gorm:"foreignKey:LanguageID;references:ID;constraint:OnDelete:CASCADE" json:"language,omitempty"`
	Title       string    `gorm:"size:255;not null" json:"title"`
	Description string    `gorm:"size:1000;not null" json:"description"`
	Text        string    `gorm:"type:text;not null" json:"text"`
}

type ArtPicture struct {
	Model
	ArtID     int64    `gorm:"not null;index" json:"art_id"`
	Art       *Art     `gorm:"foreignKey:ArtID;references:ID;constraint:OnDelete:CASCADE" json:"art,omitempty"`
	PictureID int64    `gorm:"not null;index" json:"picture_id"`
	Picture   *Picture `gorm:"foreignKey:PictureID;references:ID;constraint:OnDelete:CASCADE" json:"picture,omitempty"`
	Name      string   `gorm:"size:255;not null" json:"name"`
	Priority  *int     `json:"priority,omitempty"`
}

type Picture struct {
	Model
	Name        string       `gorm:"size:255;not null" json:"name"`
	Priority    *int         `json:"priority,omitempty"`
	PictureLink string       `gorm:"size:255;not null" json:"picture_link"`
	ArtPictures []ArtPicture `gorm:"foreignKey:PictureID" json:"art_pictures,omitempty"`
}

type Currency struct {
	Model
	CurrencyCode string `gorm:"size:3;not null;unique" json:"currency_code"`
	CurrencyName string `gorm:"size:50;not null" json:"currency_name"`
	Arts         []Art  `gorm:"foreignKey:CurrencyID" json:"arts,omitempty"`
}

type Translation struct {
	Model
	EntityID   int64     `gorm:"not null;index" json:"entity_id"`
	LanguageID int64     `gorm:"not null;index" json:"language_id"`
	Language   *Language `gorm:"foreignKey:LanguageID;references:ID;constraint:OnDelete:CASCADE" json:"language,omitempty"`
	Context    string    `gorm:"size:50;not null" json:"context"`
	Text       string    `gorm:"type:text;not null" json:"text"`
}
