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

	if data.UserRole != "admin" {
		if *data.UserID != data.AuthID {
			return &models.ServiceError{
				StatusCode: http.StatusForbidden,
				Message:    "You can only create orders for your own user account",
			}
		}
	}

	return nil
}

func verifyGetOrderByID(data models.GetOrderByIDRequest) *models.ServiceError {
	if data.UserID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "UserID has to be over 1",
		}
	}

	if data.OrderID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "OrderID has to be over 1",
		}
	}

	return nil
}
