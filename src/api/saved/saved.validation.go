package saved

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"net/http"
)

func verifyGetSavedById(data models.GetSavedByID) *models.ServiceError {
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
		if data.UserID != data.TargetID {
			return &models.ServiceError{
				StatusCode: http.StatusForbidden,
				Message:    "You are not allowed to see this route",
			}
		}
	}

	return nil
}
