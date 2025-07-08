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
