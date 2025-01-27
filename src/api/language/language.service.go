package language

import (
	"errors"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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

func listLanguageService(offset int) (*[]models.Language, *int64, *models.ServiceError) {
	var languages []models.Language
	var count int64

	if err := database.DB.Limit(5).Offset(offset * 10).Find(&languages).Error; err != nil {
		return nil, nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while retrieving languages from database",
		}
	}

	if err := database.DB.Model(&models.Language{}).Count(&count).Error; err != nil {
		return nil, nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while counting languages in database",
		}
	}

	return &languages, &count, nil
}

func createLanguageService(request models.CreateLanguageRequest) (*models.Language, *models.ServiceError) {
	var language models.Language

	// Check if language already exists
	if err := database.DB.Where("language_name = ?", request.Language).First(&language).Error; err == nil {
		return nil, &models.ServiceError{StatusCode: http.StatusConflict, Message: "Language already exists"}
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: "Database error"}
	}

	newLanguage := models.Language{
		LanguageCode: request.LanguageCode,
		LanguageName: request.Language,
	}

	if err := database.DB.Create(&newLanguage).Error; err != nil {
		return nil, &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: "Failed to save language"}
	}

	return &newLanguage, nil
}

func updateRoleService(request models.UpdateLanguageRequest) (*models.Language, *models.ServiceError) {
	var language models.Language

	if err := database.DB.First(&language, "id = ?", request.LanguageID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &models.ServiceError{StatusCode: http.StatusNotFound, Message: "Language not found"}
		}
		return nil, &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}

	language.LanguageName = request.Language
	language.LanguageCode = request.LanguageCode

	if err := database.DB.Save(&language).Error; err != nil {
		return nil, &models.ServiceError{StatusCode: http.StatusUnauthorized, Message: "Error while updating role"}
	}

	return &language, nil
}

func deleteLanguageService(languageID string) *models.ServiceError {
	var language models.Language
	parsedLanguageID, err := strconv.ParseInt(languageID, 10, 64)
	if err != nil {
		return &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: "Invalid languageID"}
	}

	if err := database.DB.First(&language, "id = ?", parsedLanguageID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &models.ServiceError{StatusCode: http.StatusNotFound, Message: "Language not found"}
		}
		return &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}

	if result := database.DB.Delete(&models.Language{}, parsedLanguageID); result.Error != nil {
		return &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: "Error occurred while deleting role"}
	}

	return nil
}
