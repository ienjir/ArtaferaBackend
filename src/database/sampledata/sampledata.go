package sampledata

import (
	"github.com/ienjir/ArtaferaBackend/src/api/auth"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
)

func CreateSampleData() error {
	// Create languages
	languages := []models.Language{
		{LanguageName: "English", LanguageCode: "en"},
		{LanguageName: "Spanish", LanguageCode: "es"},
		{LanguageName: "French", LanguageCode: "fr"},
	}
	if err := database.DB.Create(&languages).Error; err != nil {
		return err
	}

	// Create currencies
	currencies := []models.Currency{
		{Model: models.Model{ID: 1}, CurrencyCode: "USD", CurrencyName: "US Dollar"},
		{Model: models.Model{ID: 2}, CurrencyCode: "EUR", CurrencyName: "Euro"},
		{Model: models.Model{ID: 3}, CurrencyCode: "GBP", CurrencyName: "British Pound"},
	}
	if err := database.DB.Create(&currencies).Error; err != nil {
		return err
	}

	// Create users with hashed passwords
	hashPassword, err := auth.HashPassword("Password")
	if err != nil {
		return err
	}

	address1 := "123 Main St"
	city := "New York"
	postalCode := "10001"
	phone := "+14155552671"
	phoneRegion := "US"

	users := []models.User{
		{
			Model:       models.Model{ID: 1},
			Firstname:   "User1",
			Lastname:    "User1",
			Email:       "User1@example.com",
			Phone:       &phone,
			PhoneRegion: &phoneRegion,
			Address1:    &address1,
			City:        &city,
			PostalCode:  &postalCode,
			Password:    hashPassword.Hash,
			Salt:        hashPassword.Salt,
			RoleID:      1,
		},
		{
			Model:     models.Model{ID: 2},
			Firstname: "User2",
			Lastname:  "User2",
			Email:     "User2@example.com",
			Password:  hashPassword.Hash,
			Salt:      hashPassword.Salt,
			RoleID:    2,
		},
		{
			Model:     models.Model{ID: 3},
			Firstname: "User3",
			Lastname:  "User3",
			Email:     "User3@example.com",
			Password:  hashPassword.Hash,
			Salt:      hashPassword.Salt,
			RoleID:    3,
		},
	}
	if err := database.DB.Create(&users).Error; err != nil {
		return err
	}

	// Create artworks
	width := float32(100.0)
	height := float32(150.0)
	depth := float32(10.0)

	arts := []models.Art{
		{
			Model:        models.Model{ID: 1},
			Price:        1000,
			CurrencyID:   1,
			CreationYear: 2023,
			Width:        &width,
			Height:       &height,
			Depth:        &depth,
			Available:    true,
		},
		{
			Model:        models.Model{ID: 2},
			Price:        2000,
			CurrencyID:   2,
			CreationYear: 2024,
			Width:        &width,
			Height:       &height,
			Available:    false,
		},
	}
	if err := database.DB.Create(&arts).Error; err != nil {
		return err
	}

	// Create art translations
	artTranslations := []models.ArtTranslation{
		{
			Model:       models.Model{ID: 1},
			ArtID:       1,
			Title:       "Sunset",
			Description: "A beautiful sunset painting",
			Text:        "This artwork captures the essence of a summer sunset.",
		},
		{
			Model:       models.Model{ID: 2},
			ArtID:       1,
			Title:       "Puesta del sol",
			Description: "Una hermosa pintura del atardecer",
			Text:        "Esta obra de arte captura la esencia de un atardecer de verano.",
		},
	}
	if err := database.DB.Create(&artTranslations).Error; err != nil {
		return err
	}

	// Create pictures
	pictures := []models.Picture{
		{
			Model:       models.Model{ID: 1},
			Name:        "sunset_main",
			PictureLink: "https://storage.example.com/artworks/sunset_main.jpg",
		},
		{
			Model:       models.Model{ID: 2},
			Name:        "sunset_detail",
			PictureLink: "https://storage.example.com/artworks/sunset_detail.jpg",
		},
	}
	if err := database.DB.Create(&pictures).Error; err != nil {
		return err
	}

	// Create art pictures
	priority := 1
	artPictures := []models.ArtPicture{
		{
			Model:     models.Model{ID: 1},
			ArtID:     1,
			PictureID: 1,
			Name:      "Main View",
			Priority:  &priority,
		},
		{
			Model:     models.Model{ID: 2},
			ArtID:     1,
			PictureID: 2,
			Name:      "Detail View",
			Priority:  &priority,
		},
	}
	if err := database.DB.Create(&artPictures).Error; err != nil {
		return err
	}

	// Create saved items
	saved := []models.Saved{
		{
			Model:  models.Model{ID: 1},
			UserID: 1,
			ArtID:  1,
		},
		{
			Model:  models.Model{ID: 2},
			UserID: 2,
			ArtID:  2,
		},
	}
	if err := database.DB.Create(&saved).Error; err != nil {
		return err
	}

	/***reate orders
	orders := []models.Order{
		{
			Model:     models.Model{ID: 1},
			UserID:    2,
			ArtID:     2,
			OrderDate: time.Now().Add(-24 * time.Hour),
			Status:    models.OrderStatusPending,
		},
	}
	if err := database.DB.Create(&orders).Error; err != nil {
		return err
	}
	*/

	// Create translations
	translations := []models.Translation{
		{
			Model:      models.Model{ID: 1},
			EntityID:   1,
			LanguageID: 1,
			Context:    "art_description",
			Text:       "A beautiful sunset painting in oil",
		},
		{
			Model:      models.Model{ID: 2},
			EntityID:   1,
			LanguageID: 2,
			Context:    "art_description",
			Text:       "Una hermosa pintura al óleo del atardecer",
		},
	}
	if err := database.DB.Create(&translations).Error; err != nil {
		return err
	}

	return nil
}
