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

func listSavedService(data models.ListSavedRequest) (*[]models.Saved, *int64, *models.ServiceError) {
	var saved []models.Saved
	var count int64

	if err := database.DB.Preload("Art").Preload("User").Limit(5).Offset(int(data.Offset * 5)).Find(&saved).Error; err != nil {
		return nil, nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while retrieving saved from database",
		}
	}

	if err := database.DB.Model(&models.Saved{}).Count(&count).Error; err != nil {
		return nil, nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while counting saved in database",
		}
	}

	return &saved, &count, nil
}

func createSavedService(data models.CreateSavedRequest) (*models.Saved, *models.ServiceError) {
	var art models.Art
	var user models.User
	var saved models.Saved

	if err := database.DB.First(&art, data.ArtID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &models.ServiceError{
				StatusCode: http.StatusNotFound,
				Message:    "Art not found",
			}
		}
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while retrieving art",
		}
	}

	if err := database.DB.First(&user, data.TargetUserID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &models.ServiceError{
				StatusCode: http.StatusNotFound,
				Message:    "User not found",
			}
		}
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while retrieving user",
		}
	}

	if err := database.DB.Where("user_id = ? AND art_id = ?", user.ID, data.ArtID).First(&saved).Error; err == nil {
		// Record already exists
		return nil, &models.ServiceError{
			StatusCode: http.StatusConflict,
			Message:    "Art is already saved for this user",
		}
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while checking existing saved record",
		}
	}

	newSaved := models.Saved{
		UserID: user.ID,
		ArtID:  data.ArtID,
	}

	if err := database.DB.Create(&newSaved).Error; err != nil {
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to save record",
		}
	}

	return &newSaved, nil
}

func updateSavedService(data models.UpdateSavedRequest) (*models.Saved, *models.ServiceError) {
	var saved models.Saved

	if err := database.DB.First(&saved, "id = ?", data.TargetID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &models.ServiceError{StatusCode: http.StatusNotFound, Message: "Saved not found"}
		}
		return nil, &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}

	if data.TargetUserID != nil {
		saved.UserID = *data.TargetUserID
	}

	if data.ArtID != nil {
		saved.ArtID = *data.ArtID
	}

	if err := database.DB.Save(&saved).Error; err != nil {
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to update saved",
		}
	}

	return &saved, nil
}
