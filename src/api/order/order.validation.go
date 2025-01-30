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

func verifyGetOrdersForUser(data models.GetOrdersForUserRequest) *models.ServiceError {
	if data.UserID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "UserID has to be at least 1",
		}
	}

	if data.TargetUserID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "OrderID has to be at least 1",
		}
	}

	if data.Offset < 0 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "Offset has to  be 0 or more",
		}
	}

	if data.UserRole != "admin" {
		if data.UserID != data.TargetUserID {
			return &models.ServiceError{
				StatusCode: http.StatusForbidden,
				Message:    "You can only see orders for your own user account",
			}
		}
	}

	return nil
}

func verifyListOrders(data models.ListOrdersRequest) *models.ServiceError {
	if data.UserID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "UserID has to be at least 1",
		}
	}

	if data.Offset < 0 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "Offset has to be 0 or more",
		}
	}

	if data.UserRole != "admin" {
		return &models.ServiceError{
			StatusCode: http.StatusForbidden,
			Message:    "You are not allowed for this route",
		}
	}

	return nil
}
