package order

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"net/http"
)

func verifyCreateOrder(data models.CreateOrderRequest) *models.ServiceError {
	if data.UserID == nil {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "UserID is required",
		}
	}

	if *data.UserID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "UserID has to be over 1",
		}
	}

	if data.ArtID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "ArtID has to be over 1",
		}
	}

	// Get the userID from the context
	contextUserID := *data.UserID

	// If user is not an admin, they can only create orders for themselves
	if data.UserRole != "admin" {
		requestedUserID := *data.UserID
		if contextUserID != requestedUserID {
			return &models.ServiceError{
				StatusCode: http.StatusForbidden,
				Message:    "You can only create orders for your own user account",
			}
		}
	}

	return nil
}
