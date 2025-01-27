package language

import (
	"errors"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"gorm.io/gorm"
	"net/http"
)

func getLanguageByIDService(targetLanguageID string) (*models.Language, *models.ServiceError) {
	var language models.Language

	if err := database.DB.First(&language, targetLanguageID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &models.ServiceError{StatusCode: http.StatusNotFound, Message: "Language not found"}
		} else {
			return nil, &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: "Error while retrieving language"}
		}
	}

	return &language, nil
}
