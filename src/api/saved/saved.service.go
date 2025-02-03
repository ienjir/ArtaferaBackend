package saved

import (
	"errors"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"gorm.io/gorm"
	"net/http"
)

func getSavedByIDService(data models.GetSavedByIDRequest) (*models.Saved, *models.ServiceError) {
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

func getSavedForUserService(data models.GetSavedForUserRequest) (*[]models.Saved, *models.User, *int64, *models.ServiceError) {
	var saved []models.Saved
	var user models.User
	var count int64

	if err := database.DB.First(&user, data.TargetUserID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, nil, &models.ServiceError{
				StatusCode: http.StatusNotFound,
				Message:    "User not found",
			}
		} else {
			return nil, nil, nil, &models.ServiceError{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error while retrieving user",
			}
		}
	}

	if err := database.DB.Preload("Art").Where("user_id = ?", data.TargetUserID).Find(&saved).Limit(5).Offset(int(data.Offset) * 5).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, nil, &models.ServiceError{
				StatusCode: http.StatusNotFound,
				Message:    "Saved not found",
			}
		}
		return nil, nil, nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while retrieving saved",
		}
	}

	if err := database.DB.Model(&models.Saved{}).Where("user_id = ?", data.TargetUserID).Count(&count).Error; err != nil {
		return nil, nil, nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while counting saved in database",
		}
	}

	return &saved, &user, &count, nil
}
