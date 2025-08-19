package saved

import (
	"errors"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/utils"
	"gorm.io/gorm"
)

func getSavedByIDService(data models.GetSavedByIDRequest) (*models.Saved, *models.ServiceError) {
	var saved models.Saved

	if err := database.DB.Preload("Art").Preload("User").First(&saved, data.TargetID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewSavedNotFoundError()
		} else {
			return nil, utils.NewDatabaseRetrievalError()
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
			return nil, nil, nil, utils.NewUserNotFoundError()
		} else {
			return nil, nil, nil, utils.NewDatabaseRetrievalError()
		}
	}

	if err := database.DB.Preload("Art").Where("user_id = ?", data.TargetUserID).Find(&saved).Limit(5).Offset(int(data.Offset) * 5).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, nil, utils.NewSavedNotFoundError()
		}
		return nil, nil, nil, utils.NewDatabaseRetrievalError()
	}

	if err := database.DB.Model(&models.Saved{}).Where("user_id = ?", data.TargetUserID).Count(&count).Error; err != nil {
		return nil, nil, nil, utils.NewDatabaseCountError()
	}

	return &saved, &user, &count, nil
}

func listSavedService(data models.ListSavedRequest) (*[]models.Saved, *int64, *models.ServiceError) {
	var saved []models.Saved
	var count int64

	if err := database.DB.Preload("Art").Preload("User").Limit(5).Offset(int(data.Offset * 5)).Find(&saved).Error; err != nil {
		return nil, nil, utils.NewDatabaseRetrievalError()
	}

	if err := database.DB.Model(&models.Saved{}).Count(&count).Error; err != nil {
		return nil, nil, utils.NewDatabaseCountError()
	}

	return &saved, &count, nil
}

func createSavedService(data models.CreateSavedRequest) (*models.Saved, *models.ServiceError) {
	var art models.Art
	var user models.User
	var saved models.Saved

	if err := database.DB.First(&art, data.ArtID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewArtNotFoundError()
		}
		return nil, utils.NewDatabaseRetrievalError()
	}

	if err := database.DB.First(&user, data.TargetUserID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewUserNotFoundError()
		}
		return nil, utils.NewDatabaseRetrievalError()
	}

	if err := database.DB.Where("user_id = ? AND art_id = ?", user.ID, data.ArtID).First(&saved).Error; err == nil {
		// Record already exists
		return nil, utils.NewArtAlreadySavedError()
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.NewDatabaseRetrievalError()
	}

	newSaved := models.Saved{
		UserID: user.ID,
		ArtID:  data.ArtID,
	}

	if err := database.DB.Create(&newSaved).Error; err != nil {
		return nil, utils.NewDatabaseCreateError()
	}

	return &newSaved, nil
}

func updateSavedService(data models.UpdateSavedRequest) (*models.Saved, *models.ServiceError) {
	var saved models.Saved

	if err := database.DB.First(&saved, "id = ?", data.TargetID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewSavedNotFoundError()
		}
		return nil, utils.NewDatabaseRetrievalError()
	}

	if data.TargetUserID != nil {
		saved.UserID = *data.TargetUserID
	}

	if data.ArtID != nil {
		saved.ArtID = *data.ArtID
	}

	if err := database.DB.Save(&saved).Error; err != nil {
		return nil, utils.NewDatabaseUpdateError()
	}

	return &saved, nil
}

func deleteSavedService(data models.DeleteSavedRequest) *models.ServiceError {
	var saved models.Saved

	if err := database.DB.First(&saved, "id = ?", data.TargetID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NewSavedNotFoundError()
		}
		return utils.NewDatabaseRetrievalError()
	}

	if data.UserRole != "admin" {
		if saved.UserID != data.UserID {
			return utils.NewNotAllowedRouteError()
		}
	}

	if result := database.DB.Delete(&models.Saved{}, saved.ID); result.Error != nil {
		return utils.NewDatabaseDeleteError()
	}

	return nil
}
