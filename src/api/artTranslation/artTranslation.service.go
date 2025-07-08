package ArtTranslation

import (
	"errors"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"gorm.io/gorm"
	"net/http"
)

func getArtTranslationByIDService(data models.GetArtTranslationByIDRequest) (*models.ArtTranslation, *models.ServiceError) {
	var ArtTranslation models.ArtTranslation

	if err := database.DB.First(&ArtTranslation, data.TargetID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &models.ServiceError{
				StatusCode: http.StatusNotFound,
				Message:    "ArtTranslation not found",
			}
		} else {
			return nil, &models.ServiceError{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error while retrieving ArtTranslation",
			}
		}
	}

	return &ArtTranslation, nil
}
