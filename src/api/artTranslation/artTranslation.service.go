package artTranslation

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

func listArtTranslationService(data models.ListArtTranslationRequest) (*[]models.ArtTranslation, *int64, *models.ServiceError) {
	var artTranslations []models.ArtTranslation
	var count int64

	if err := database.DB.Limit(10).Offset(int(data.Offset * 10)).Find(&artTranslations).Error; err != nil {
		return nil, nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while retrieving artTranslations from database",
		}
	}

	if err := database.DB.Model(&models.ArtTranslation{}).Count(&count).Error; err != nil {
		return nil, nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while counting artTranslations in database",
		}
	}

	return &artTranslations, &count, nil
}
