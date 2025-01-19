package models

import (
	"gorm.io/gorm"
	"time"
)

var AllModels = []interface{}{
	&User{},
	&Role{},
	&UserRole{},
	&Text{},
	&Translation{},
	&Language{},
	&Art{},
	&ArtPicture{},
	&Picture{},
	&Order{},
	&OrderDetail{},
	&Payment{},
	&Currency{},
}

// User model
type User struct {
	gorm.Model
	Firstname   string     `gorm:"size:255;not null" json:"firstname"`
	Lastname    string     `gorm:"size:255;not null" json:"lastname"`
	Email       string     `gorm:"size:255;not null;unique;index" json:"email"`
	Phone       *string    `gorm:"size:20" json:"phone,omitempty"`
	PhoneRegion *string    `gorm:"size:2" json:"phone_region,omitempty"`
	Address1    *string    `gorm:"size:255" json:"address1,omitempty"`
	Address2    *string    `gorm:"size:255" json:"address2,omitempty"`
	City        *string    `gorm:"size:255" json:"city,omitempty"`
	PostalCode  *string    `gorm:"size:32" json:"postal_code,omitempty"`
	Password    []byte     `gorm:"type:bytea;not null" json:"-"`
	Salt        []byte     `gorm:"type:bytea;not null" json:"-"`
	LastAccess  *time.Time `json:"last_access,omitempty"`
	IsDeleted   bool       `gorm:"default:false" json:"-"`
	RoleID      uint       `gorm:"not null" json:"-"`
	Role        Role       `gorm:"foreignKey:RoleID" json:"role"`
}

// Role model
type Role struct {
	gorm.Model
	Role string `gorm:"size:50;not null" json:"role"`
}

// Text model
type Text struct {
	gorm.Model
}

// Translation model
type Translation struct {
	gorm.Model
	EntityID   int    `gorm:"not null;index" json:"entity_id"`
	LanguageID int    `gorm:"not null;index" json:"language_id"`
	Context    string `gorm:"size:50;not null" json:"context"`
	Text       string `gorm:"type:text;not null" json:"text"`
}

// Language model
type Language struct {
	gorm.Model
	LanguageName string `gorm:"size:50;not null;unique" json:"language_name"`
	LanguageCode string `gorm:"size:2;not null;unique" json:"language_code"`
}

// Art model
type Art struct {
	gorm.Model
	Price        int      `gorm:"not null" json:"price"`
	CurrencyID   int      `gorm:"not null;index" json:"currency_id"`
	CreationYear string   `gorm:"size:4;not null" json:"creation_year"`
	Width        *float64 `gorm:"type:decimal(8,2)" json:"width,omitempty"`
	Height       *float64 `gorm:"type:decimal(8,2)" json:"height,omitempty"`
	Depth        *float64 `gorm:"type:decimal(8,2)" json:"depth,omitempty"`
}

// ArtPicture model
type ArtPicture struct {
	gorm.Model
	ArtID     int  `gorm:"not null;index" json:"art_id"`
	PictureID int  `gorm:"not null;index" json:"picture_id"`
	Priority  *int `json:"priority,omitempty"`
}

// Picture model
type Picture struct {
	gorm.Model
	PictureLink string `gorm:"size:255;not null" json:"picture_link"`
}

// Order model
type Order struct {
	gorm.Model
	UserID     int       `gorm:"not null;index" json:"user_id"`
	OrderDate  time.Time `gorm:"not null" json:"order_date"`
	TotalPrice float64   `gorm:"type:decimal(10,2);not null" json:"total_price"`
	Status     string    `gorm:"size:50;not null" json:"status"`
}

// OrderDetail model
type OrderDetail struct {
	gorm.Model
	OrderID  int     `gorm:"not null;index" json:"order_id"`
	ArtID    int     `gorm:"not null;index" json:"art_id"`
	Quantity int     `gorm:"not null" json:"quantity"`
	Price    float64 `gorm:"type:decimal(10,2);not null" json:"price"`
}

// Payment model
type Payment struct {
	gorm.Model
	OrderID       int       `gorm:"not null;index" json:"order_id"`
	PaymentDate   time.Time `gorm:"not null" json:"payment_date"`
	Amount        float64   `gorm:"type:decimal(10,2);not null" json:"amount"`
	PaymentMethod string    `gorm:"size:50;not null" json:"payment_method"`
	Status        string    `gorm:"size:50;not null" json:"status"`
}

// Currency model
type Currency struct {
	gorm.Model
	CurrencyCode string `gorm:"size:3;not null;unique" json:"currency_code"`
	CurrencyName string `gorm:"size:50;not null" json:"currency_name"`
}
