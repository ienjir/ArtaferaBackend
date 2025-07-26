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

func createArtTranslationService(data models.CreateArtTranslationRequest, languageID int64) (*models.ArtTranslation, *models.ServiceError) {
	var artTranslation models.ArtTranslation
	var newArtTranslation models.ArtTranslation

	if err := database.DB.Where("art_id = ? AND language_id = ?", data.ArtID, languageID).First(&artTranslation).Error; err == nil {
		return nil, &models.ServiceError{
			StatusCode: http.StatusConflict,
			Message:    "Art translation already exists for this language",
		}
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Database error",
		}
	}

	newArtTranslation = models.ArtTranslation{
		ArtID:       data.ArtID,
		LanguageID:  languageID,
		Title:       data.Title,
		Description: data.Description,
		Text:        data.Text,
	}

	if err := database.DB.Create(&newArtTranslation).Error; err != nil {
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to save art translation",
		}
	}

	return &newArtTranslation, nil
}

func updateArtTranslation(data models.UpdateArtTranslationRequest) (*models.ArtTranslation, *models.ServiceError) {
	var artTranslation models.ArtTranslation

	if err := database.DB.First(&artTranslation, "id = ?", data.TargetID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &models.ServiceError{
				StatusCode: http.StatusNotFound,
				Message:    "ArtTranslation not found",
			}
		}
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Database error",
		}
	}

	if data.LanguageID != nil {
		artTranslation.LanguageID = *data.LanguageID
	}

	if data.Title != nil {
		artTranslation.Title = *data.Title
	}

	if data.Description != nil {
		artTranslation.Description = *data.Description
	}

	if data.Text != nil {
		artTranslation.Text = *data.Text
	}

	if err := database.DB.Save(&artTranslation).Error; err != nil {
		return nil, &models.ServiceError{
			StatusCode: http.StatusUnauthorized,
			Message:    "Error while updating artTranslation",
		}
	}

	return &artTranslation, nil
}

func deleteArtTranslationService(data models.DeleteArtTranslationRequest) *models.ServiceError {
	var artTranslation models.ArtTranslation

	if err := database.DB.First(&artTranslation, "id = ?", data.TargetID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &models.ServiceError{
				StatusCode: http.StatusNotFound,
				Message:    "ArtTranslation not found",
			}
		}
		return &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Database error",
		}
	}

	if result := database.DB.Delete(&models.ArtTranslation{}, data.TargetID); result.Error != nil {
		return &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error occurred while deleting Art Translation",
		}
	}

	return nil
}
