package artTranslation

import (
	"errors"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/utils"
	"gorm.io/gorm"
)

func getArtTranslationByIDService(data models.GetArtTranslationByIDRequest) (*models.ArtTranslation, *models.ServiceError) {
	var ArtTranslation models.ArtTranslation

	if err := database.DB.First(&ArtTranslation, data.TargetID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewArtTranslationNotFoundError()
		} else {
			return nil, utils.NewDatabaseRetrievalError()
		}
	}

	return &ArtTranslation, nil
}

func listArtTranslationService(data models.ListArtTranslationRequest) (*[]models.ArtTranslation, *int64, *models.ServiceError) {
	var artTranslations []models.ArtTranslation
	var count int64

	if err := database.DB.Limit(10).Offset(int(data.Offset * 10)).Find(&artTranslations).Error; err != nil {
		return nil, nil, utils.NewDatabaseRetrievalError()
	}

	if err := database.DB.Model(&models.ArtTranslation{}).Count(&count).Error; err != nil {
		return nil, nil, utils.NewDatabaseCountError()
	}

	return &artTranslations, &count, nil
}

func createArtTranslationService(data models.CreateArtTranslationRequest, languageID int64) (*models.ArtTranslation, *models.ServiceError) {
	var artTranslation models.ArtTranslation
	var newArtTranslation models.ArtTranslation

	if err := database.DB.Where("art_id = ? AND language_id = ?", data.ArtID, languageID).First(&artTranslation).Error; err == nil {
		return nil, utils.NewArtTranslationExistsError()
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.NewDatabaseRetrievalError()
	}

	newArtTranslation = models.ArtTranslation{
		ArtID:       data.ArtID,
		LanguageID:  languageID,
		Title:       data.Title,
		Description: data.Description,
		Text:        data.Text,
	}

	if err := database.DB.Create(&newArtTranslation).Error; err != nil {
		return nil, utils.NewDatabaseCreateError()
	}

	return &newArtTranslation, nil
}

func updateArtTranslation(data models.UpdateArtTranslationRequest) (*models.ArtTranslation, *models.ServiceError) {
	var artTranslation models.ArtTranslation

	if err := database.DB.First(&artTranslation, "id = ?", data.TargetID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewArtTranslationNotFoundError()
		}
		return nil, utils.NewDatabaseRetrievalError()
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
		return nil, utils.NewDatabaseUpdateError()
	}

	return &artTranslation, nil
}

func deleteArtTranslationService(data models.DeleteArtTranslationRequest) *models.ServiceError {
	var artTranslation models.ArtTranslation

	if err := database.DB.First(&artTranslation, "id = ?", data.TargetID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NewArtTranslationNotFoundError()
		}
		return utils.NewDatabaseRetrievalError()
	}

	if result := database.DB.Delete(&models.ArtTranslation{}, data.TargetID); result.Error != nil {
		return utils.NewDatabaseDeleteError()
	}

	return nil
}
