package utils

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"gorm.io/gorm"
	"net/http"
	"os"
)

var GinMode int

func SetGinMode() {
	switch os.Getenv("MODE") {
	case "debug":
		gin.SetMode("debug")
		GinMode = 0
		fmt.Println("Running in debug mode")

	case "test":
		gin.SetMode("test")
		GinMode = 1
		fmt.Println("Running in test mode")

	case "release":
		gin.SetMode("release")
		GinMode = 2
		fmt.Println("Running in release mode")

	default:
		fmt.Println("No application mode was provided")
		fmt.Println("Please set 'MODE' in the .env file to one of the following 3 options:")
		fmt.Println("debug")
		fmt.Println("test")
		fmt.Println("release")
		os.Exit(1)
	}

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func LanguageCodeToID(languageCode string) (*models.Language, error) {
	var language models.Language

	if err := database.DB.Where("language_code = ?", languageCode).First(&language).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("language with code '%s' not found", languageCode)
		}
		return nil, fmt.Errorf("error retrieving language: %w", err)
	}

	return &language, nil
}
