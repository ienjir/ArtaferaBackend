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

type Model struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	Model
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
	RoleID      int64      `gorm:"default:1;not null" json:"role_id"`
	Role        *Role      `gorm:"foreignKey:RoleID;references:ID;constraint:OnDelete:SET DEFAULT" json:"role"`
}

type Role struct {
	Model
	Name string `gorm:"column:name;not null" json:"name"`
}

type Language struct {
	gorm.Model
	LanguageName string `gorm:"size:50;not null;unique" json:"language_name"`
	LanguageCode string `gorm:"size:2;not null;unique;index" json:"language_code"`
}

type Saved struct {
	Model
	UserID int64 `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;not null;index" json:"userID"`
	ArtID  int64 `gorm:"foreignKey:ArtID;references:ID;constraint:OnDelete:SET NULL" json:"artID"`
}

type Order struct {
	Model
	UserID    int64       `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;not null" json:"user_id"`
	User      *User       `json:"user,omitempty"`
	ArtID     int64       `gorm:"foreignKey:ArtID;references:ID;constraint:OnDelete:CASCADE;not null" json:"art_id"`
	Art       *Art        `json:"art,omitempty"`
	OrderDate time.Time   `gorm:"not null;default:CURRENT_TIMESTAMP" json:"order_date"`
	Status    OrderStatus `gorm:"size:50;not null;default:'pending'" json:"status"`
}

type Art struct {
	Model
	Price        int64            `gorm:"not null;check:price >= 0" json:"price"`
	CurrencyID   int64            `gorm:"foreignKey:CurrencyID;references:ID;constraint:OnDelete:SET NULL;index" json:"currency_id"`
	Currency     *Currency        `json:"currency,omitempty"`
	CreationYear int              `gorm:"not null" json:"creation_year" binding:"required,min=1000,max=9999"`
	Width        *float32         `gorm:"type:decimal(8,2)" json:"width,omitempty"`
	Height       *float32         `gorm:"type:decimal(8,2)" json:"height,omitempty"`
	Depth        *float32         `gorm:"type:decimal(8,2)" json:"depth,omitempty"`
	Pictures     []ArtPicture     `gorm:"foreignKey:ArtID" json:"pictures,omitempty"`
	Translations []ArtTranslation `gorm:"foreignKey:ArtID" json:"translations,omitempty"`
	Available    bool             `gorm:"default:true" json:"available"`
}

type ArtTranslation struct {
	Model
	ArtID       int64  `gorm:"foreignKey:ArtID;reference:ID;constraint:OnDelete:CASCADE;not null" json:"artID"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Text        string `gorm:"type:text" json:"text"`
}

type ArtPicture struct {
	Model
	ArtID     int64  `gorm:"foreignKey:ArtID;references:ID;constraint:OnDelete:CASCADE;not null;index" json:"art_id"`
	PictureID int64  `gorm:"not null;index" json:"picture_id"`
	Name      string `gorm:"not null" json:"name"`
	Priority  *int   `json:"priority,omitempty"`
}

type Picture struct {
	Model
	Name        string `gorm:"not null" json:"name"`
	Priority    *int   `json:"priority,omitempty"`
	PictureLink string `gorm:"size:255;not null" json:"picture_link"`
}

type Currency struct {
	Model
	CurrencyCode string `gorm:"size:3;not null;unique" json:"currency_code"`
	CurrencyName string `gorm:"size:50;not null" json:"currency_name"`
}

type Translation struct {
	Model
	EntityID   int64  `gorm:"not null;index" json:"entity_id"`
	LanguageID int64  `gorm:"foreignKey:LanguageID;reference:ID;not null;index" json:"language_id"`
	Context    string `gorm:"size:50;not null" json:"context"`
	Text       string `gorm:"type:text;not null" json:"text"`
}
