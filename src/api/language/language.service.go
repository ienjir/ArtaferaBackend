package language

import (
	"errors"
	"fmt"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/utils"
	"gorm.io/gorm"
)

func getLanguageByIDService(data models.GetLanguageByIDRequest) (*models.Language, *models.ServiceError) {
	var language models.Language

	if err := database.DB.First(&language, data.TargetID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewLanguageNotFoundError()
		} else {
			return nil, utils.NewDatabaseRetrievalError()
		}
	}

	return &language, nil
}

func listLanguageService(data models.ListLanguageRequest) (*[]models.Language, *int64, *models.ServiceError) {
	var languages []models.Language
	var count int64

	if err := database.DB.Limit(10).Offset(int(data.Offset * 10)).Find(&languages).Error; err != nil {
		return nil, nil, utils.NewDatabaseRetrievalError()
	}

	if err := database.DB.Model(&models.Language{}).Count(&count).Error; err != nil {
		return nil, nil, utils.NewDatabaseCountError()
	}

	return &languages, &count, nil
}

func createLanguageService(data models.CreateLanguageRequest) (*models.Language, *models.ServiceError) {
	var language models.Language
	var newLanguage models.Language

	if err := database.DB.Where("language_name = ?", data.Language).First(&language).Error; err == nil {
		return nil, utils.NewLanguageAlreadyExistsError()
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.NewDatabaseRetrievalError()
	}

	newLanguage = models.Language{
		LanguageName: data.Language,
		LanguageCode: data.LanguageCode,
	}

	if err := database.DB.Create(&newLanguage).Error; err != nil {
		return nil, utils.NewDatabaseCreateError()
	}

	return &newLanguage, nil
}

func updateLanguageService(data models.UpdateLanguageRequest) (*models.Language, *models.ServiceError) {
	var language models.Language

	if err := database.DB.First(&language, "id = ?", data.TargetID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewLanguageNotFoundError()
		}
		return nil, utils.NewDatabaseRetrievalError()
	}

	language.LanguageName = data.Language
	language.LanguageCode = data.LanguageCode

	if err := database.DB.Save(&language).Error; err != nil {
		return nil, utils.NewDatabaseUpdateError()
	}

	return &language, nil
}

func deleteLanguageService(data models.DeleteLanguageRequest) *models.ServiceError {
	var language models.Language

	if err := database.DB.First(&language, "id = ?", data.TargetID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NewLanguageNotFoundError()
		}
		return utils.NewDatabaseRetrievalError()
	}

	if result := database.DB.Delete(&models.Language{}, data.TargetID); result.Error != nil {
		return utils.NewDatabaseDeleteError()
	}

	return nil
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
