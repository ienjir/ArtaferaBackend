package artTranslation

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/validation"
)

func verifyGetArtTranslationByIDRequest(data models.GetArtTranslationByIDRequest) *models.ServiceError {
	return validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateID(data.TargetID, "TargetID").
		GetFirstError()
}

func verifyListArtTranslation(data models.ListArtTranslationRequest) *models.ServiceError {
	return validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateAdminRole(data.UserRole).
		ValidateOffset(int64(data.Offset)).
		GetFirstError()
}

func verifyCreateArtTranslation(data models.CreateArtTranslationRequest) *models.ServiceError {
	return validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateAdminRole(data.UserRole).
		ValidateID(data.ArtID, "ArtID").
		ValidateNotEmpty(&data.LanguageCode, "LanguageCode").
		ValidateNotEmpty(&data.Title, "Title").
		ValidateNotEmpty(&data.Description, "Description").
		ValidateNotEmpty(&data.Text, "Text").
		GetFirstError()
}

func verifyUpdateArtTranslation(data models.UpdateArtTranslationRequest) *models.ServiceError {
	validator := validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateAdminRole(data.UserRole).
		ValidateID(data.TargetID, "TargetID")

	if data.LanguageCode != nil {
		validator = validator.ValidateNotEmpty(data.LanguageCode, "LanguageCode")
	}

	if data.Title != nil {
		validator = validator.ValidateNotEmpty(data.Title, "Title")
	}

	if data.Description != nil {
		validator = validator.ValidateNotEmpty(data.Description, "Description")
	}

	if data.Text != nil {
		validator = validator.ValidateNotEmpty(data.Text, "Text")
	}

	return validator.GetFirstError()
}

func verifyDeleteArtTranslation(data models.DeleteArtTranslationRequest) *models.ServiceError {
	return validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateAdminRole(data.UserRole).
		ValidateID(data.TargetID, "TargetID").
		GetFirstError()
}
