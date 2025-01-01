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
	Firstname  string  `gorm:"size:255;not null"`
	Lastname   string  `gorm:"size:255;not null"`
	Email      string  `gorm:"size:255;not null;unique"`
	Phone      *string `gorm:"size:20"`
	Address1   *string `gorm:"size:255"`
	Address2   *string `gorm:"size:255"`
	City       *string `gorm:"size:255"`
	PostalCode *string `gorm:"size:32"`
	Password   []byte  `gorm:"size:60;not null"`
	LastAccess *time.Time
	IsDeleted  bool `gorm:"default:false"`
}

// Role model
type Role struct {
	gorm.Model
	Role string `gorm:"size:50;not null"`
}

// UserRole model
type UserRole struct {
	gorm.Model
	UserID int `gorm:"not null;index"`
	RoleID int `gorm:"not null;index"`
}

// Text model
type Text struct {
	gorm.Model
}

// Translation model
type Translation struct {
	gorm.Model
	EntityID   int    `gorm:"not null;index"`
	LanguageID int    `gorm:"not null;index"`
	Context    string `gorm:"size:50;not null"`
	Text       string `gorm:"type:text;not null"`
}

// Language model
type Language struct {
	gorm.Model
	LanguageName string `gorm:"size:50;not null;unique"`
	LanguageCode string `gorm:"size:2;not null;unique"`
}

// Art model
type Art struct {
	gorm.Model
	Price        int      `gorm:"not null"`
	CurrencyID   int      `gorm:"not null;index"`
	CreationYear string   `gorm:"size:4;not null"`
	Width        *float64 `gorm:"type:decimal(8,2)"`
	Height       *float64 `gorm:"type:decimal(8,2)"`
	Depth        *float64 `gorm:"type:decimal(8,2)"`
}

// ArtPicture model
type ArtPicture struct {
	gorm.Model
	ArtID     int `gorm:"not null;index"`
	PictureID int `gorm:"not null;index"`
	Priority  *int
}

// Picture model
type Picture struct {
	gorm.Model
	PictureLink string `gorm:"size:255;not null"`
}

// Order model
type Order struct {
	gorm.Model
	UserID     int       `gorm:"not null;index"`
	OrderDate  time.Time `gorm:"not null"`
	TotalPrice float64   `gorm:"type:decimal(10,2);not null"`
	Status     string    `gorm:"size:50;not null"`
}

// OrderDetail model
type OrderDetail struct {
	gorm.Model
	OrderID  int     `gorm:"not null;index"`
	ArtID    int     `gorm:"not null;index"`
	Quantity int     `gorm:"not null"`
	Price    float64 `gorm:"type:decimal(10,2);not null"`
}

// Payment model
type Payment struct {
	gorm.Model
	OrderID       int       `gorm:"not null;index"`
	PaymentDate   time.Time `gorm:"not null"`
	Amount        float64   `gorm:"type:decimal(10,2);not null"`
	PaymentMethod string    `gorm:"size:50;not null"`
	Status        string    `gorm:"size:50;not null"`
}

// Currency model
type Currency struct {
	gorm.Model
	CurrencyCode string `gorm:"size:3;not null;unique"`
	CurrencyName string `gorm:"size:50;not null"`
}
