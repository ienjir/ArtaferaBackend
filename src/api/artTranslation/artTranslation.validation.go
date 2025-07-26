package artTranslation

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"net/http"
)

func verifyGetArtTranslationByIDRequest(data models.GetArtTranslationByIDRequest) *models.ServiceError {
	if data.UserID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "UserID has to be over 1",
		}
	}

	if data.TargetID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "TargetID has to be over 1",
		}
	}

	return nil
}

func verifyListArtTranslation(data models.ListArtTranslationRequest) *models.ServiceError {
	if data.UserID < 0 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "UserID has to be at least 0",
		}
	}

	if data.UserRole != "admin" {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "You are not allowed to see this route",
		}
	}

	if data.Offset < -1 {
		return &models.ServiceError{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Offset can't be less than -1",
		}
	}

	return nil
}

func verifyCreateArtTranslation(data models.CreateArtTranslationRequest) *models.ServiceError {
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

	if data.ArtID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "ArtID has to be at least 1",
		}
	}

	if len(data.LanguageCode) != 2 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "Language code must be 2 chars",
		}
	}

	if data.Title == "" {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "Title can't be empty",
		}
	}

	if data.Description == "" {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "Title can't be empty",
		}
	}

	if data.Text == "" {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "Title can't be empty",
		}
	}

	return nil
}

func verifyUpdateArtTranslation(data models.UpdateArtTranslationRequest) *models.ServiceError {
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

	if data.TargetID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "TargetID has to be at least 1",
		}
	}

	if data.LanguageCode != nil {
		if l := len(*data.LanguageCode); l != 2 {
			return &models.ServiceError{
				StatusCode: http.StatusBadRequest,
				Message:    "Language code needs to be 2 chars",
			}
		}
	}

	if data.Title != nil {
		if *data.Text == "" {
			return &models.ServiceError{
				StatusCode: http.StatusBadRequest,
				Message:    "Title can't be empty",
			}
		}
	}

	if data.Description != nil {
		if *data.Description == "" {
			return &models.ServiceError{
				StatusCode: http.StatusBadRequest,
				Message:    "Description can't be empty",
			}
		}
	}

	if data.Text != nil {
		if *data.Text == "" {
			return &models.ServiceError{
				StatusCode: http.StatusBadRequest,
				Message:    "Text can't be empty",
			}
		}
	}

	return nil
}

func verifyDeleteArtTranslation(data models.DeleteArtTranslationRequest) *models.ServiceError {
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

	if data.TargetID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "TargetID has to be at least 1",
		}
	}

	return nil
}
