package language

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/validation"
)

func verifyGetLanguageByIDRequest(data models.GetLanguageByIDRequest) *models.ServiceError {
	return validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateID(data.TargetID, "TargetID").
		ValidateAdminRole(data.UserRole).
		GetFirstError()
}

func verifyListLanguagesRequest(data models.ListLanguageRequest) *models.ServiceError {
	return validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateAdminRole(data.UserRole).
		ValidateOffset(int64(data.Offset)).
		GetFirstError()
}

func verifyCreateLanguageRequest(data models.CreateLanguageRequest) *models.ServiceError {
	return validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateAdminRole(data.UserRole).
		ValidateNotEmpty(&data.Language, "Language").
		ValidateNotEmpty(&data.LanguageCode, "LanguageCode").
		GetFirstError()
}

func verifyUpdateLanguageRequest(data models.UpdateLanguageRequest) *models.ServiceError {
	return validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateAdminRole(data.UserRole).
		ValidateNotEmpty(&data.Language, "Language").
		ValidateNotEmpty(&data.LanguageCode, "LanguageCode").
		ValidateID(data.TargetID, "TargetID").
		GetFirstError()
}

func verifyDeleteLanguageRequest(data models.DeleteLanguageRequest) *models.ServiceError {
	return validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateID(data.TargetID, "TargetID").
		ValidateAdminRole(data.UserRole).
		GetFirstError()
}
