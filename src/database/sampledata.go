package database

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"gorm.io/gorm"
	"time"
)

func GenerateFakeData(database *gorm.DB) {
	// Sample Users
	users := []models.User{
		{Firstname: "John", Lastname: "Doe", Email: "john.doe@example.com", Password: []byte("password123")},
		{Firstname: "Jane", Lastname: "Smith", Email: "jane.smith@example.com", Password: []byte("password456")},
	}

	// Sample Roles
	roles := []models.Role{
		{Role: "Admin"},
		{Role: "Customer"},
	}

	// Sample UserRoles
	userRoles := []models.UserRole{
		{UserID: 1, RoleID: 1},
		{UserID: 2, RoleID: 2},
	}

	// Sample Languages
	languages := []models.Language{
		{LanguageName: "English", LanguageCode: "EN"},
		{LanguageName: "German", LanguageCode: "DE"},
	}

	// Sample Texts and Translations
	texts := []models.Text{
		{}, {}, // Adding dummy entries since Text has no fields
	}

	translations := []models.Translation{
		{EntityID: 1, LanguageID: 1, Context: "WelcomeMessage", Text: "Welcome!"},
		{EntityID: 1, LanguageID: 2, Context: "WelcomeMessage", Text: "Willkommen!"},
	}

	// Sample Currency
	currencies := []models.Currency{
		{CurrencyCode: "USD", CurrencyName: "US Dollar"},
		{CurrencyCode: "EUR", CurrencyName: "Euro"},
	}

	// Sample Arts
	arts := []models.Art{
		{Price: 100, CurrencyID: 1, CreationYear: "2020", Width: toPointer(50.0), Height: toPointer(70.0)},
		{Price: 200, CurrencyID: 2, CreationYear: "2021", Width: toPointer(60.0), Height: toPointer(80.0)},
	}

	// Sample Pictures and ArtPictures
	pictures := []models.Picture{
		{PictureLink: "https://example.com/art1.jpg"},
		{PictureLink: "https://example.com/art2.jpg"},
	}

	artPictures := []models.ArtPicture{
		{ArtID: 1, PictureID: 1, Priority: toPointer(1)},
		{ArtID: 2, PictureID: 2, Priority: toPointer(2)},
	}

	// Sample Orders, OrderDetails, and Payments
	orders := []models.Order{
		{UserID: 1, OrderDate: time.Now(), TotalPrice: 150.0, Status: "Completed"},
	}

	orderDetails := []models.OrderDetail{
		{OrderID: 1, ArtID: 1, Quantity: 1, Price: 100.0},
		{OrderID: 1, ArtID: 2, Quantity: 1, Price: 200.0},
	}

	payments := []models.Payment{
		{OrderID: 1, PaymentDate: time.Now(), Amount: 150.0, PaymentMethod: "Credit Card", Status: "Paid"},
	}

	// Bulk insert data
	DB.Create(&users)
	DB.Create(&roles)
	DB.Create(&userRoles)
	DB.Create(&languages)
	DB.Create(&texts)
	DB.Create(&translations)
	DB.Create(&currencies)
	DB.Create(&arts)
	DB.Create(&pictures)
	DB.Create(&artPictures)
	DB.Create(&orders)
	DB.Create(&orderDetails)
	DB.Create(&payments)
}

func toPointer[T any](v T) *T {
	return &v
}
