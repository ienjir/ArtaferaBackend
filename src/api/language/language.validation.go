package language

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"net/http"
)

func verifyGetLanguageByIDRequest(data models.GetLanguageByIDRequest) *models.ServiceError {
	if data.UserID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "UserID has to be at least 1",
		}
	}

	if data.TargetID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "TargetID has to be at least 1",
		}
	}

	if data.UserRole != "admin" {
		return &models.ServiceError{
			StatusCode: http.StatusForbidden,
			Message:    "You are not allowed to see this route",
		}
	}

	return nil
}

func verifyListLanguagesRequest(data models.ListLanguageRequest) *models.ServiceError {
	if data.UserID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "UserID has to be at least 1",
		}
	}

	if data.UserRole != "admin" {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "You are not allowed to see this route",
		}
	}

	if data.Offset < 0 {
		return &models.ServiceError{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Offset can't be less than 0",
		}
	}

	return nil
}

func verifyCreateLanguageRequest(data models.CreateLanguageRequest) *models.ServiceError {
	if data.UserID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "UserID has to be at least 1",
		}
	}

	if data.UserRole != "admin" {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "You are not allowed to see this route",
		}
	}

	if data.Language == "" {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "Language can't be empty",
		}
	}

	if data.LanguageCode == "" {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "LanguageCode can't be empty",
		}
	}

	return nil
}

func verifyUpdateLanguageRequest(data models.UpdateLanguageRequest) *models.ServiceError {
	if data.UserID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "UserID has to be at least 1",
		}
	}

	if data.UserRole != "admin" {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "You are not allowed to see this route",
		}
	}

	if data.Language == "" {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "Language can't be empty",
		}
	}

	if data.LanguageCode == "" {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "LanguageCode can't be empty",
		}
	}

	if data.TargetID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "TargetID has to be at least 1",
		}
	}

	return nil
}

func verifyDeleteLanguageRequest(data models.DeleteLanguageRequest) *models.ServiceError {
	if data.UserID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "UserID has to be at least 1",
		}
	}

	if data.TargetID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "TargetID has to be at least 1",
		}
	}

	if data.UserRole != "admin" {
		return &models.ServiceError{
			StatusCode: http.StatusForbidden,
			Message:    "You are not allowed to see this route",
		}
	}

	return nil
}
