package models

import (
	"gorm.io/gorm"
	"time"
)

var AllModels = []interface{}{
	&User{},
	&Role{},
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

type Model struct {
	ID        int            `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAT gorm.DeletedAt `gorm:"index" json:"-"`
}

type User struct {
	Model
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
	LastAccess  *time.Time `json:"-,omitempty"`
	RoleID      uint       `gorm:"default:1;not null"` // Ensure default is set
	Role        *Role      `gorm:"foreignKey:RoleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET DEFAULT" json:"role"`
}

type Role struct {
	Model
	Name  string `gorm:"column:name;not null"`
	Users []User `gorm:"foreignKey:RoleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET DEFAULT"`
}

type Translation struct {
	Id         int    `gorm:"column:id"`
	EntityID   int    `gorm:"not null;index" json:"entity_id"`
	LanguageID int    `gorm:"not null;index" json:"language_id"`
	Context    string `gorm:"size:50;not null" json:"context"`
	Text       string `gorm:"type:text;not null" json:"text"`
}

type Language struct {
	gorm.Model
	LanguageName string `gorm:"size:50;not null;unique" json:"language_name"`
	LanguageCode string `gorm:"size:2;not null;unique" json:"language_code"`
}

type Art struct {
	gorm.Model
	Price        int      `gorm:"not null" json:"price"`
	CurrencyID   int      `gorm:"not null;index" json:"currency_id"`
	CreationYear string   `gorm:"size:4;not null" json:"creation_year"`
	Width        *float64 `gorm:"type:decimal(8,2)" json:"width,omitempty"`
	Height       *float64 `gorm:"type:decimal(8,2)" json:"height,omitempty"`
	Depth        *float64 `gorm:"type:decimal(8,2)" json:"depth,omitempty"`
}

type ArtPicture struct {
	gorm.Model
	ArtID     int  `gorm:"not null;index" json:"art_id"`
	PictureID int  `gorm:"not null;index" json:"picture_id"`
	Priority  *int `json:"priority,omitempty"`
}

type Picture struct {
	gorm.Model
	PictureLink string `gorm:"size:255;not null" json:"picture_link"`
}

type Order struct {
	gorm.Model
	UserID     int       `gorm:"not null;index" json:"user_id"`
	OrderDate  time.Time `gorm:"not null" json:"order_date"`
	TotalPrice float64   `gorm:"type:decimal(10,2);not null" json:"total_price"`
	Status     string    `gorm:"size:50;not null" json:"status"`
}

type OrderDetail struct {
	gorm.Model
	OrderID  int     `gorm:"not null;index" json:"order_id"`
	ArtID    int     `gorm:"not null;index" json:"art_id"`
	Quantity int     `gorm:"not null" json:"quantity"`
	Price    float64 `gorm:"type:decimal(10,2);not null" json:"price"`
}

type Payment struct {
	gorm.Model
	OrderID       int       `gorm:"not null;index" json:"order_id"`
	PaymentDate   time.Time `gorm:"not null" json:"payment_date"`
	Amount        float64   `gorm:"type:decimal(10,2);not null" json:"amount"`
	PaymentMethod string    `gorm:"size:50;not null" json:"payment_method"`
	Status        string    `gorm:"size:50;not null" json:"status"`
}

type Currency struct {
	gorm.Model
	CurrencyCode string `gorm:"size:3;not null;unique" json:"currency_code"`
	CurrencyName string `gorm:"size:50;not null" json:"currency_name"`
}
