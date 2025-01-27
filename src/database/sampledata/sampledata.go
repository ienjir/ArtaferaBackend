package sampledata

import (
	"github.com/ienjir/ArtaferaBackend/src/api/auth"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"time"
)

func GenerateFakeData() error {
	// Sample Users
	users := []models.User{
		{Firstname: "Firstname1", Lastname: "Lastname1", Email: "User1@example.com", Phone: nil, PhoneRegion: nil, Address1: nil, Address2: nil, City: nil, PostalCode: nil, Password: []byte("Password"), Salt: []byte("Salt"), LastAccess: nil},
		{Firstname: "Firstname2", Lastname: "Lastname2", Email: "User2@example.com", Password: []byte("Password"), Salt: []byte("Salt"), Address1: toPointer("123 Maple Street"), PostalCode: toPointer("10001"), RoleID: 2},
		{Firstname: "Firstname3", Lastname: "Lastname3", Email: "User3@example.com", Password: []byte("Password"), Salt: []byte("Salt"), PhoneRegion: toPointer("US"), RoleID: 3},
		{Firstname: "Lastname4", Lastname: "Lastname4", Email: "User4@example.com", Password: []byte("Password"), Salt: []byte("Salt"), Address2: toPointer("Apt 4B"), LastAccess: toPointer(time.Now()), RoleID: 1},
		{Firstname: "Lastname5", Lastname: "Lastname5", Email: "User5@example.com", Password: []byte("Password"), Salt: []byte("Salt"), RoleID: 2},
	}

	for i := 0; i < len(users); i++ {
		user := users[i]
		password, err := auth.HashPassword(string(user.Password))
		if err != nil {
			return err
		}
		users[i].Password = password.Hash
		users[i].Salt = password.Salt
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
	database.DB.Create(&users)
	database.DB.Create(&languages)
	database.DB.Create(&texts)
	database.DB.Create(&translations)
	database.DB.Create(&currencies)
	database.DB.Create(&arts)
	database.DB.Create(&pictures)
	database.DB.Create(&artPictures)
	database.DB.Create(&orders)
	database.DB.Create(&orderDetails)
	database.DB.Create(&payments)

	return nil
}

func toPointer[T any](v T) *T {
	return &v
}
