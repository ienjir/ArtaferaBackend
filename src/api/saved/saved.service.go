package saved

import (
	"errors"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"gorm.io/gorm"
	"net/http"
)

func getSavedByIDService(data models.GetSavedByID) (*models.Saved, *models.ServiceError) {
	var saved models.Saved

	if err := database.DB.Preload("Art").Preload("User").First(&saved, data.TargetID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &models.ServiceError{
				StatusCode: http.StatusNotFound,
				Message:    "Saved not found",
			}
		} else {
			return nil, &models.ServiceError{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error while retrieving saved",
			}
		}
	}

	return &saved, nil
}
