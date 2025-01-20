package database

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"gorm.io/gorm"
	"time"
)

func GenerateFakeData(database *gorm.DB) {
	// Sample Users
	users := []models.User{
		{Firstname: "Luis", Lastname: "Hänny", Email: "luis.haenny@swisscom.com", Phone: nil, PhoneRegion: nil, Address1: nil, Address2: nil, City: nil, PostalCode: nil, Password: []byte("Password"), Salt: []byte("Kek"), LastAccess: nil},
		{Firstname: "Jane", Lastname: "Smith", Email: "jane.smith@example.com", Password: []byte("password456"), Salt: []byte("Salt"), Address1: toPointer("123 Maple Street"), PostalCode: toPointer("10001"), RoleID: 2},
		{Firstname: "Alice", Lastname: "Johnson", Email: "alice.johnson@example.com", Password: []byte("password789"), Salt: []byte("Salt"), PhoneRegion: toPointer("US"), RoleID: 3},
		{Firstname: "Bob", Lastname: "Brown", Email: "bob.brown@example.com", Password: []byte("password234"), Salt: []byte("Salt"), Address2: toPointer("Apt 4B"), LastAccess: toPointer(time.Now()), RoleID: 1},
		{Firstname: "Charlie", Lastname: "Williams", Email: "charlie.williams@example.com", Password: []byte("password345"), Salt: []byte("Salt"), RoleID: 2},
	}

	// Sample Roles
	roles := []models.Role{
		{Role: "Admin"},
		{Role: "Customer"},
		{Role: "Artist"},
	}

	// Sample Languages
	languages := []models.Language{
		{LanguageName: "English", LanguageCode: "EN"},
		{LanguageName: "German", LanguageCode: "DE"},
		{LanguageName: "Spanish", LanguageCode: "ES"},
	}

	// Sample Texts and Translations
	texts := []models.Text{
		{}, {}, {}, // Adding dummy entries since Text has no fields
	}

	translations := []models.Translation{
		{EntityID: 1, LanguageID: 1, Context: "WelcomeMessage", Text: "Welcome!"},
		{EntityID: 1, LanguageID: 2, Context: "WelcomeMessage", Text: "Willkommen!"},
		{EntityID: 1, LanguageID: 3, Context: "WelcomeMessage", Text: "¡Bienvenidos!"},
	}

	// Sample Currency
	currencies := []models.Currency{
		{CurrencyCode: "USD", CurrencyName: "US Dollar"},
		{CurrencyCode: "EUR", CurrencyName: "Euro"},
		{CurrencyCode: "CHF", CurrencyName: "Swiss Franc"},
	}

	// Sample Arts
	arts := []models.Art{
		{Price: 100, CurrencyID: 1, CreationYear: "2020", Width: toPointer(50.0), Height: toPointer(70.0)},
		{Price: 200, CurrencyID: 2, CreationYear: "2021", Width: toPointer(60.0), Height: toPointer(80.0), Depth: toPointer(30.0)},
		{Price: 300, CurrencyID: 3, CreationYear: "2022"},
	}

	// Sample Pictures and ArtPictures
	pictures := []models.Picture{
		{PictureLink: "https://example.com/art1.jpg"},
		{PictureLink: "https://example.com/art2.jpg"},
		{PictureLink: "https://example.com/art3.jpg"},
	}

	artPictures := []models.ArtPicture{
		{ArtID: 1, PictureID: 1, Priority: toPointer(1)},
		{ArtID: 2, PictureID: 2, Priority: toPointer(2)},
		{ArtID: 3, PictureID: 3},
	}

	// Sample Orders, OrderDetails, and Payments
	orders := []models.Order{
		{UserID: 1, OrderDate: time.Now(), TotalPrice: 300.0, Status: "Completed"},
		{UserID: 2, OrderDate: time.Now().Add(-24 * time.Hour), TotalPrice: 200.0, Status: "Pending"},
	}

	orderDetails := []models.OrderDetail{
		{OrderID: 1, ArtID: 1, Quantity: 1, Price: 100.0},
		{OrderID: 1, ArtID: 2, Quantity: 1, Price: 200.0},
		{OrderID: 2, ArtID: 3, Quantity: 1, Price: 300.0},
	}

	payments := []models.Payment{
		{OrderID: 1, PaymentDate: time.Now(), Amount: 300.0, PaymentMethod: "Credit Card", Status: "Paid"},
		{OrderID: 2, PaymentDate: time.Now().Add(-24 * time.Hour), Amount: 200.0, PaymentMethod: "Bank Transfer", Status: "Pending"},
	}

	// Bulk insert data
	database.Create(&roles)
	database.Create(&users)
	database.Create(&languages)
	database.Create(&texts)
	database.Create(&translations)
	database.Create(&currencies)
	database.Create(&arts)
	database.Create(&pictures)
	database.Create(&artPictures)
	database.Create(&orders)
	database.Create(&orderDetails)
	database.Create(&payments)
}

func toPointer[T any](v T) *T {
	return &v
}
