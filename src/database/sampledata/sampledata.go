package sampledata

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/ienjir/ArtaferaBackend/src/api/auth"
	"github.com/ienjir/ArtaferaBackend/src/database"
	miniobucket "github.com/ienjir/ArtaferaBackend/src/minio"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/minio/minio-go/v7"
)

func SeedDatabase() error {

	// Languages
	languages := []models.Language{
		{LanguageName: "English", LanguageCode: "en"},
		{LanguageName: "German", LanguageCode: "de"},
		{LanguageName: "French", LanguageCode: "fr"},
	}
	for i := range languages {
		if err := database.Repositories.Language.Create(&languages[i]); err != nil {
			return fmt.Errorf("failed to create language: %v", err.Message)
		}
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
	for i := range users {
		if err := database.Repositories.User.Create(&users[i]); err != nil {
			return fmt.Errorf("failed to create user: %v", err.Message)
		}
	}

	// Currencies
	currencies := []models.Currency{
		{CurrencyCode: "USD", CurrencyName: "US Dollar"},
		{CurrencyCode: "EUR", CurrencyName: "Euro"},
		{CurrencyCode: "GBP", CurrencyName: "British Pound"},
		{CurrencyCode: "JPY", CurrencyName: "Japanese Yen"},
		{CurrencyCode: "CHF", CurrencyName: "Swiss Franc"},
	}
	for i := range currencies {
		if err := database.Repositories.Currency.Create(&currencies[i]); err != nil {
			return fmt.Errorf("failed to create currency: %v", err.Message)
		}
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
			Featured:     true,
			Visible:      true,
		},
		{
			Price:        1500,
			CurrencyID:   2,
			CreationYear: 2021,
			Width:        float32Ptr(80.0),
			Height:       float32Ptr(60.0),
			Available:    true,
			Featured:     true,
			Visible:      true,
		},
		{
			Price:        2000,
			CurrencyID:   3,
			CreationYear: 2019,
			Width:        float32Ptr(100.0),
			Height:       float32Ptr(70.0),
			Available:    true,
			Featured:     true,
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
	for i := range arts {
		if err := database.Repositories.Art.Create(&arts[i]); err != nil {
			return fmt.Errorf("failed to create art: %v", err.Message)
		}
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
	labels := map[int]map[string]string{
		1: {"en": "Wind", "de": "Wind", "fr": "Windee"},
		2: {"en": "Clouds", "de": "Wolken", "fr": "Wolkeen"},
		3: {"en": "Ice", "de": "Eis", "fr": "Iseeee"},
		4: {"en": "Rectangles", "Vierecke": "Wind", "fr": "Quatreequee"},
		5: {"en": "Idk", "de": "iwn", "fr": "jnsp"},
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
				Text:        "## Test Heading\n\n" +
				"Lorem ipsum dolor sit amet, **consetetur sadipscing elitr**, sed diam nonumy " +
				"eirmod tempor invidunt ut *labore et dolore magna aliquyam erat*. Quisque vel " +
				"augue id metus fermentum aliquet, und irgendwie wirkt das alles weniger stabil als es sollte.\n\n" +
				"> A small blockquote to check indentation and styling.\n\n" +
				"Curabitur 'inline code' wird auch gern uebersehen, waehrend ein laengerer Abschnitt wie " +
				"dieser prueft, ob dein Renderer sauber umbrechen kann.\n\n" +
				"- List item eins\n" +
				"- List item zwei\n" +
				"- List item drei\n\n" +
				"Am Ende noch ein [Testlink](https://example.com), gefolgt von einem abschliessenden Satz, " +
				"der genug Text liefert, um zu sehen, ob dein Markdown System irgendwas verschluckt oder falsch interpretiert.",
				Label:       labels[int(artID)][lang.LanguageCode],
			})
		}
	}
	for i := range artTranslations {
		if err := database.Repositories.ArtTranslation.Create(&artTranslations[i]); err != nil {
			return fmt.Errorf("failed to create art translation: %v", err.Message)
		}
	}

	// Pictures
	pictures := []models.Picture{
		{Name: "bc846bc97a96ddaee17a3d4196bed886", IsPublic: true, Type: ".jpg"},
		{Name: "6b7c732715b67af383765fddd317639c", IsPublic: true, Type: ".jpg"},
		{Name: "b35b304bafaba9e2e8e41f7e9b6006b9", IsPublic: true, Type: ".jpg"},
		{Name: "000bc457d300a6eea6f6e15ebfa8d4a1", IsPublic: true, Type: ".jpg"},
		{Name: "b562e93414be79f9fa91cab1d688264f", IsPublic: true, Type: ".jpg"},
		{Name: "7d0b14209673572cbf361872c327c805", IsPublic: true, Type: ".jpg"},
		{Name: "384c22ed9ad97008f059c79f046d592d", IsPublic: true, Type: ".jpg"},
		{Name: "26dcb69c91de090eeb9914d1190b476c", IsPublic: true, Type: ".jpg"},
		{Name: "bfda02666976025ac5bb0227e40701f2", IsPublic: true, Type: ".jpg"},
		{Name: "51d849089bedcd11e04c5586013854c9", IsPublic: true, Type: ".jpg"},
		{Name: "7d075abf1fe444c33632e29d29fa080f", IsPublic: true, Type: ".jpg"},
		{Name: "a938357207e3caa046dd0186be1c6afa", IsPublic: true, Type: ".jpg"},
		{Name: "fbfe3dedeea9d5156cc327f96282bca6", IsPublic: true, Type: ".jpg"},
		{Name: "f3edef79bb9f99de3b509da3b1cc2810", IsPublic: true, Type: ".jpg"},
		{Name: "9f68aa9d9a6b0b90f4140ec61d4bda80", IsPublic: true, Type: ".jpg"},
	}
	for i := range pictures {
		if err := database.Repositories.Picture.Create(&pictures[i]); err != nil {
			return fmt.Errorf("failed to create picture: %v", err.Message)
		}
	}

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("failed to get current file path")
	}
	currentDir := filepath.Dir(filename)
	imagesDir := filepath.Join(currentDir, "images")

	for i, picture := range pictures {
		srcPath := filepath.Join(imagesDir, fmt.Sprintf("%d.jpg", i+1))

		fmt.Printf("PictureName: %s \n", picture.Name)

		// Determine destination bucket based on IsPublic flag
		bucketName := "pictures-p"
		if picture.IsPublic {
			bucketName = "pictures"
		}

		// Open the image file
		file, err := os.Open(srcPath)
		if err != nil {
			return fmt.Errorf("failed to open image file %s: %w \n", srcPath, err)
		}
		defer file.Close()

		// Get file stats to determine size
		fileInfo, err := file.Stat()
		if err != nil {
			return fmt.Errorf("failed to get file stats for %s: %w \n", srcPath, err)
		}

		// Upload file to MinIO
		objectName := fmt.Sprintf("%s.jpg", picture.Name)
		_, err = miniobucket.MinioClient.PutObject(context.Background(), bucketName, objectName, file, fileInfo.Size(), minio.PutObjectOptions{
			ContentType: "image/jpeg",
		})
		if err != nil {
			return fmt.Errorf("failed to upload image to MinIO: %w \n", err)
		}
	}

	// Art Pictures
	artPictures := []models.ArtPicture{
		{ArtID: 1, PictureID: 1, Name: "Main view", Priority: 1},
		{ArtID: 1, PictureID: 2, Name: "Detail view", Priority: 2},
		{ArtID: 1, PictureID: 3, Name: "Detail view", Priority: 3},
		{ArtID: 2, PictureID: 4, Name: "Main view", Priority: 1},
		{ArtID: 2, PictureID: 5, Name: "Main view", Priority: 2},
		{ArtID: 2, PictureID: 6, Name: "Main view", Priority: 3},
		{ArtID: 3, PictureID: 7, Name: "Main view", Priority: 1},
		{ArtID: 3, PictureID: 8, Name: "Main view", Priority: 2},
		{ArtID: 3, PictureID: 9, Name: "Main view", Priority: 3},
		{ArtID: 4, PictureID: 10, Name: "Main view", Priority: 1},
		{ArtID: 4, PictureID: 11, Name: "Main view", Priority: 2},
		{ArtID: 4, PictureID: 12, Name: "Main view", Priority: 3},
		{ArtID: 5, PictureID: 13, Name: "Main view", Priority: 1},
		{ArtID: 5, PictureID: 14, Name: "Main view", Priority: 2},
		{ArtID: 5, PictureID: 15, Name: "Main view", Priority: 3},
	}
	for i := range artPictures {
		if err := database.Repositories.ArtPicture.Create(&artPictures[i]); err != nil {
			return fmt.Errorf("failed to create art picture: %v", err.Message)
		}
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
	for i := range orders {
		if err := database.Repositories.Order.Create(&orders[i]); err != nil {
			return fmt.Errorf("failed to create order: %v", err.Message)
		}
	}

	// Saved items
	saved := []models.Saved{
		{UserID: 1, ArtID: 2},
		{UserID: 1, ArtID: 3},
		{UserID: 2, ArtID: 1},
		{UserID: 3, ArtID: 4},
		{UserID: 4, ArtID: 5},
	}
	for i := range saved {
		if err := database.Repositories.Saved.Create(&saved[i]); err != nil {
			return fmt.Errorf("failed to create saved: %v", err.Message)
		}
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
	for i := range welcomeMessage {
		if err := database.Repositories.Translation.Create(&welcomeMessage[i]); err != nil {
			return fmt.Errorf("failed to create translation: %v", err.Message)
		}
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

func int64Ptr(f int64) *int64 {
	return &f
}

func seedMinioPictures(pictures []models.Picture) {
}
