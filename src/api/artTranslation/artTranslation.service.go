package artTranslation

import (
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/utils"
)

func getArtTranslationByIDService(data models.GetArtTranslationByIDRequest) (*models.ArtTranslation, *models.ServiceError) {
	artTranslation, err := database.Repositories.ArtTranslation.GetByID(data.TargetID)
	if err != nil {
		if err.StatusCode == 404 {
			return nil, utils.NewArtTranslationNotFoundError()
		}
		return nil, err
	}

	return artTranslation, nil
}

func listArtTranslationService(data models.ListArtTranslationRequest) (*[]models.ArtTranslation, *int64, *models.ServiceError) {
	artTranslations, err := database.Repositories.ArtTranslation.List(int(data.Offset*10), 10)
	if err != nil {
		return nil, nil, err
	}

	count, err := database.Repositories.ArtTranslation.Count()
	if err != nil {
		return nil, nil, err
	}

	return artTranslations, count, nil
}

func createArtTranslationService(data models.CreateArtTranslationRequest, languageID int64) (*models.ArtTranslation, *models.ServiceError) {
	// Check if art translation already exists for this art and language
	query := database.Repositories.ArtTranslation.Query().Where("art_id = ? AND language_id = ?", data.ArtID, languageID)
	var existingTranslation models.ArtTranslation
	if queryErr := query.First(&existingTranslation).Error; queryErr == nil {
		return nil, utils.NewArtTranslationExistsError()
	}

	newArtTranslation := models.ArtTranslation{
		ArtID:       data.ArtID,
		LanguageID:  languageID,
		Title:       data.Title,
		Description: data.Description,
		Text:        data.Text,
	}

	if err := database.Repositories.ArtTranslation.Create(&newArtTranslation); err != nil {
		return nil, err
	}

	return &newArtTranslation, nil
}

func updateArtTranslation(data models.UpdateArtTranslationRequest) (*models.ArtTranslation, *models.ServiceError) {
	artTranslation, err := database.Repositories.ArtTranslation.GetByID(data.TargetID)
	if err != nil {
		if err.StatusCode == 404 {
			return nil, utils.NewArtTranslationNotFoundError()
		}
		return nil, err
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

	if err := database.Repositories.ArtTranslation.Update(artTranslation); err != nil {
		return nil, err
	}

	return artTranslation, nil
}

func deleteArtTranslationService(data models.DeleteArtTranslationRequest) *models.ServiceError {
	if err := database.Repositories.ArtTranslation.Delete(data.TargetID); err != nil {
		if err.StatusCode == 404 {
			return utils.NewArtTranslationNotFoundError()
		}
		return err
	}

	return nil
}
