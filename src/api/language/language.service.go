package language

import (
	"fmt"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/utils"
)

func getLanguageByIDService(data models.GetLanguageByIDRequest) (*models.Language, *models.ServiceError) {
	language, err := database.Repositories.Language.GetByID(data.TargetID)
	if err != nil {
		if err.StatusCode == 404 {
			return nil, utils.NewLanguageNotFoundError()
		}
		return nil, err
	}

	return language, nil
}

func listLanguageService(data models.ListLanguageRequest) (*[]models.Language, *int64, *models.ServiceError) {
	languages, err := database.Repositories.Language.List(int(data.Offset*10), 10)
	if err != nil {
		return nil, nil, err
	}

	count, err := database.Repositories.Language.Count()
	if err != nil {
		return nil, nil, err
	}

	return languages, count, nil
}

func createLanguageService(data models.CreateLanguageRequest) (*models.Language, *models.ServiceError) {
	// Check if language already exists
	if existingLanguage, err := database.Repositories.Language.FindByField("language_name", data.Language); err == nil && existingLanguage != nil {
		return nil, utils.NewLanguageAlreadyExistsError()
	}

	newLanguage := models.Language{
		LanguageName: data.Language,
		LanguageCode: data.LanguageCode,
	}

	if err := database.Repositories.Language.Create(&newLanguage); err != nil {
		return nil, err
	}

	return &newLanguage, nil
}

func updateLanguageService(data models.UpdateLanguageRequest) (*models.Language, *models.ServiceError) {
	language, err := database.Repositories.Language.GetByID(data.TargetID)
	if err != nil {
		if err.StatusCode == 404 {
			return nil, utils.NewLanguageNotFoundError()
		}
		return nil, err
	}

	language.LanguageName = data.Language
	language.LanguageCode = data.LanguageCode

	if err := database.Repositories.Language.Update(language); err != nil {
		return nil, err
	}

	return language, nil
}

func deleteLanguageService(data models.DeleteLanguageRequest) *models.ServiceError {
	if err := database.Repositories.Language.Delete(data.TargetID); err != nil {
		if err.StatusCode == 404 {
			return utils.NewLanguageNotFoundError()
		}
		return err
	}

	return nil
}

func LanguageCodeToID(languageCode string) (*models.Language, error) {
	language, err := database.Repositories.Language.FindByField("language_code", languageCode)
	if err != nil {
		if err.StatusCode == 404 {
			return nil, fmt.Errorf("language with code '%s' not found", languageCode)
		}
		return nil, fmt.Errorf("error retrieving language: %v", err.Message)
	}

	return language, nil
}
