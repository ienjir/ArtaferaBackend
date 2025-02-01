package order

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/validation"
	"net/http"
)

func verifyCreateOrder(data models.CreateOrderRequest) *models.ServiceError {
	if data.TargetUserID == nil {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "UserID is required",
		}
	}

	if *data.TargetUserID < 1 {
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
		if *data.TargetUserID != data.UserID {
			return &models.ServiceError{
				StatusCode: http.StatusForbidden,
				Message:    "You can only create orders for your own user account",
			}
		}
	}

	return nil
}

func verifyGetOrderByIDRequest(data models.GetOrderByIDRequest) *models.ServiceError {
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

func verifyGetOrdersForUserRequest(data models.GetOrdersForUserRequest) *models.ServiceError {
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

func verifyListOrdersRequest(data models.ListOrdersRequest) *models.ServiceError {
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

func verifyUpdateOrderRequest(data models.UpdateOrderRequest) *models.ServiceError {
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

	if data.TargetUserID != nil {
		if *data.TargetUserID < 1 {
			return &models.ServiceError{
				StatusCode: http.StatusBadRequest,
				Message:    "TargetUserID has to be at least 1",
			}
		}
	}

	if data.ArtID != nil {
		if *data.ArtID < 1 {
			return &models.ServiceError{
				StatusCode: http.StatusBadRequest,
				Message:    "ArtID has to be over 1",
			}
		}
	}

	if data.Status != nil {
		_, err := validation.ValidateStatusString(*data.Status)
		if err != nil {
			return err
		}
	}

	return nil
}

func verifyDeleteOrderRequest(data models.DeleteOrderRequest) *models.ServiceError {
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
