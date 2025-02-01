package sampledata

import (
	"github.com/ienjir/ArtaferaBackend/src/api/auth"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"time"
)

func SeedDatabase() error {

	// Languages
	languages := []models.Language{
		{LanguageName: "English", LanguageCode: "en"},
		{LanguageName: "German", LanguageCode: "de"},
		{LanguageName: "French", LanguageCode: "fr"},
	}
	if err := database.DB.Create(&languages).Error; err != nil {
		return err
	}

	hashPassword, _ := auth.HashPassword("Password")
	// Users
	users := []models.User{
		{
			Firstname:   "User1",
			Lastname:    "User1",
			Email:       "user1@example.com",
			Phone:       stringPtr("+14155552671"),
			PhoneRegion: stringPtr("US"),
			Address1:    stringPtr("123 Main St"),
			City:        stringPtr("San Francisco"),
			PostalCode:  stringPtr("94105"),
			Password:    hashPassword.Hash,
			Salt:        hashPassword.Salt,
			RoleID:      1,
		},
		{
			Firstname:   "User2",
			Lastname:    "User2",
			Email:       "user2@example.com",
			Phone:       stringPtr("+442071234567"),
			PhoneRegion: stringPtr("GB"),
			Address1:    stringPtr("456 High Street"),
			City:        stringPtr("London"),
			PostalCode:  stringPtr("SW1A 1AA"),
			Password:    hashPassword.Hash,
			Salt:        hashPassword.Salt,
			RoleID:      2,
		},
		{
			Firstname:   "User3",
			Lastname:    "User3",
			Email:       "user3@example.com",
			Phone:       stringPtr("+34911234567"),
			PhoneRegion: stringPtr("ES"),
			Address1:    stringPtr("789 Calle Principal"),
			City:        stringPtr("Madrid"),
			PostalCode:  stringPtr("28001"),
			Password:    hashPassword.Hash,
			Salt:        hashPassword.Salt,
			RoleID:      1,
		},
		{
			Firstname:   "User4",
			Lastname:    "User4",
			Email:       "user4@example.com",
			Phone:       stringPtr("+491711234567"),
			PhoneRegion: stringPtr("DE"),
			Address1:    stringPtr("101 Hauptstrasse"),
			City:        stringPtr("Berlin"),
			PostalCode:  stringPtr("10115"),
			Password:    hashPassword.Hash,
			Salt:        hashPassword.Salt,
			RoleID:      1,
		},
		{
			Firstname:   "User5",
			Lastname:    "User5",
			Email:       "user5@example.com",
			Phone:       stringPtr("+33123456789"),
			PhoneRegion: stringPtr("FR"),
			Address1:    stringPtr("202 Rue Principale"),
			City:        stringPtr("Paris"),
			PostalCode:  stringPtr("75001"),
			Password:    hashPassword.Hash,
			Salt:        hashPassword.Salt,
			RoleID:      1,
		},
	}
	if err := database.DB.Create(&users).Error; err != nil {
		return err
	}

	// Currencies
	currencies := []models.Currency{
		{CurrencyCode: "USD", CurrencyName: "US Dollar"},
		{CurrencyCode: "EUR", CurrencyName: "Euro"},
		{CurrencyCode: "GBP", CurrencyName: "British Pound"},
		{CurrencyCode: "JPY", CurrencyName: "Japanese Yen"},
		{CurrencyCode: "CHF", CurrencyName: "Swiss Franc"},
	}
	if err := database.DB.Create(&currencies).Error; err != nil {
		return err
	}

	// Art pieces
	arts := []models.Art{
		{
			Price:        1000,
			CurrencyID:   1,
			CreationYear: 2020,
			Width:        float32Ptr(60.5),
			Height:       float32Ptr(40.2),
			Available:    true,
			Visible:      true,
		},
		{
			Price:        1500,
			CurrencyID:   2,
			CreationYear: 2021,
			Width:        float32Ptr(80.0),
			Height:       float32Ptr(60.0),
			Available:    true,
			Visible:      true,
		},
		{
			Price:        2000,
			CurrencyID:   3,
			CreationYear: 2019,
			Width:        float32Ptr(100.0),
			Height:       float32Ptr(70.0),
			Available:    true,
			Visible:      true,
		},
		{
			Price:        2500,
			CurrencyID:   4,
			CreationYear: 2022,
			Width:        float32Ptr(120.0),
			Height:       float32Ptr(90.0),
			Available:    true,
			Visible:      true,
		},
		{
			Price:        3000,
			CurrencyID:   5,
			CreationYear: 2023,
			Width:        float32Ptr(150.0),
			Height:       float32Ptr(100.0),
			Available:    true,
			Visible:      true,
		},
	}
	if err := database.DB.Create(&arts).Error; err != nil {
		return err
	}

	// Art Translations
	var artTranslations []models.ArtTranslation
	titles := map[int]map[string]string{
		1: {"en": "Sunset", "de": "Sonnenuntergang", "fr": "Coucher de soleil"},
		2: {"en": "Mountain Lake", "de": "Bergsee", "fr": "Lac de montagne"},
		3: {"en": "Forest Path", "de": "Waldweg", "fr": "Chemin forestier"},
		4: {"en": "City Lights", "de": "Stadtlichter", "fr": "Lumières de la ville"},
		5: {"en": "Ocean Waves", "de": "Meereswellen", "fr": "Vagues de l'océan"},
	}
	descriptions := map[int]map[string]string{
		1: {
			"en": "A beautiful sunset over the ocean",
			"de": "Ein wunderschöner Sonnenuntergang über dem Ozean",
			"fr": "Un magnifique coucher de soleil sur l'océan",
		},
		2: {
			"en": "Serene mountain lake in the Alps",
			"de": "Ruhiger Bergsee in den Alpen",
			"fr": "Lac de montagne serein dans les Alpes",
		},
		3: {
			"en": "Mystical path through an ancient forest",
			"de": "Mystischer Weg durch einen alten Wald",
			"fr": "Chemin mystique à travers une forêt ancienne",
		},
		4: {
			"en": "Vibrant city lights at night",
			"de": "Lebhafte Stadtlichter bei Nacht",
			"fr": "Lumières vibrantes de la ville la nuit",
		},
		5: {
			"en": "Powerful ocean waves at sunset",
			"de": "Kraftvolle Meereswellen bei Sonnenuntergang",
			"fr": "Puissantes vagues de l'océan au coucher du soleil",
		},
	}

	for artID := int64(1); artID <= 5; artID++ {
		for _, lang := range languages {
			artTranslations = append(artTranslations, models.ArtTranslation{
				ArtID:       artID,
				LanguageID:  lang.ID,
				Title:       titles[int(artID)][lang.LanguageCode],
				Description: descriptions[int(artID)][lang.LanguageCode],
				Text:        "Detailed artwork description goes here. This would be a longer text about the artwork's history and significance.",
			})
		}
	}
	if err := database.DB.Create(&artTranslations).Error; err != nil {
		return err
	}

	// Pictures
	pictures := []models.Picture{
		{Name: "sunset_main", Priority: intPtr(1), PictureLink: "/images/sunset_1.jpg"},
		{Name: "sunset_detail", Priority: intPtr(2), PictureLink: "/images/sunset_2.jpg"},
		{Name: "mountain_main", Priority: intPtr(1), PictureLink: "/images/mountain_1.jpg"},
		{Name: "forest_main", Priority: intPtr(1), PictureLink: "/images/forest_1.jpg"},
		{Name: "city_main", Priority: intPtr(1), PictureLink: "/images/city_1.jpg"},
	}
	if err := database.DB.Create(&pictures).Error; err != nil {
		return err
	}

	// Art Pictures
	artPictures := []models.ArtPicture{
		{ArtID: 1, PictureID: 1, Name: "Main view", Priority: intPtr(1)},
		{ArtID: 1, PictureID: 2, Name: "Detail view", Priority: intPtr(2)},
		{ArtID: 2, PictureID: 3, Name: "Main view", Priority: intPtr(1)},
		{ArtID: 3, PictureID: 4, Name: "Main view", Priority: intPtr(1)},
		{ArtID: 4, PictureID: 5, Name: "Main view", Priority: intPtr(1)},
	}
	if err := database.DB.Create(&artPictures).Error; err != nil {
		return err
	}

	// Orders
	orders := []models.Order{
		{
			UserID:    1,
			ArtID:     1,
			OrderDate: time.Now().Add(-48 * time.Hour),
			Status:    models.OrderStatusDelivered,
		},
		{
			UserID:    2,
			ArtID:     2,
			OrderDate: time.Now().Add(-24 * time.Hour),
			Status:    models.OrderStatusShipped,
		},
		{
			UserID:    3,
			ArtID:     3,
			OrderDate: time.Now().Add(-12 * time.Hour),
			Status:    models.OrderStatusPaid,
		},
		{
			UserID:    4,
			ArtID:     4,
			OrderDate: time.Now().Add(-6 * time.Hour),
			Status:    models.OrderStatusPending,
		},
		{
			UserID:    5,
			ArtID:     5,
			OrderDate: time.Now(),
			Status:    models.OrderStatusPending,
		},
	}
	if err := database.DB.Create(&orders).Error; err != nil {
		return err
	}

	// Saved items
	saved := []models.Saved{
		{UserID: 1, ArtID: 2},
		{UserID: 1, ArtID: 3},
		{UserID: 2, ArtID: 1},
		{UserID: 3, ArtID: 4},
		{UserID: 4, ArtID: 5},
	}
	if err := database.DB.Create(&saved).Error; err != nil {
		return err
	}

	// Translations
	welcomeMessage := []models.Translation{
		{
			EntityID:   1,
			LanguageID: 1,
			Context:    "welcome_message",
			Text:       "Welcome to our art gallery!",
		},
		{
			EntityID:   1,
			LanguageID: 2,
			Context:    "welcome_message",
			Text:       "Willkommen in unserer Kunstgalerie!",
		},
		{
			EntityID:   1,
			LanguageID: 3,
			Context:    "welcome_message",
			Text:       "Bienvenue dans notre galerie d'art!",
		},
	}
	if err := database.DB.Create(&welcomeMessage).Error; err != nil {
		return err
	}

	return nil
}

// Helper functions for creating pointers
func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}

func float32Ptr(f float32) *float32 {
	return &f
}
