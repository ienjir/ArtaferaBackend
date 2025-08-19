package order

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/utils"
	"github.com/ienjir/ArtaferaBackend/src/validation"
)

func verifyGetOrderByIDRequest(data models.GetOrderByIDRequest) *models.ServiceError {
	return validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateID(data.OrderID, "OrderID").
		GetFirstError()
}

func verifyGetOrdersForUserRequest(data models.GetOrdersForUserRequest) *models.ServiceError {
	validator := validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateID(data.TargetUserID, "TargetUserID").
		ValidateOffset(int64(data.Offset))

	if err := validator.GetFirstError(); err != nil {
		return err
	}

	if data.UserRole != "admin" && data.UserID != data.TargetUserID {
		return utils.NewOwnerOnlyOrdersError()
	}

	return nil
}

func verifyListOrdersRequest(data models.ListOrdersRequest) *models.ServiceError {
	return validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateOffset(int64(data.Offset)).
		ValidateAdminRole(data.UserRole).
		GetFirstError()
}

func verifyCreateOrder(data models.CreateOrderRequest) *models.ServiceError {
	validator := validation.NewValidator().
		ValidateID(data.ArtID, "ArtID")

	if data.TargetUserID == nil {
		return utils.NewFieldRequiredError("UserID")
	}

	validator = validator.ValidateID(*data.TargetUserID, "TargetUserID")

	if err := validator.GetFirstError(); err != nil {
		return err
	}

	if data.UserRole != "admin" && *data.TargetUserID != data.UserID {
		return utils.NewOwnerOnlyCreateError()
	}

	return nil
}

func verifyUpdateOrderRequest(data models.UpdateOrderRequest) *models.ServiceError {
	validator := validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateID(data.TargetID, "TargetID")

	if data.TargetUserID != nil {
		validator = validator.ValidateID(*data.TargetUserID, "TargetUserID")
	}

	if data.ArtID != nil {
		validator = validator.ValidateID(*data.ArtID, "ArtID")
	}

	if data.Status != nil {
		_, err := validation.ValidateStatusString(*data.Status)
		if err != nil {
			return err
		}
	}

	if err := validator.GetFirstError(); err != nil {
		return err
	}

	if data.UserRole != "admin" && data.TargetUserID != nil && data.UserID != *data.TargetUserID {
		return utils.NewOwnerOnlySavedError()
	}

	return nil
}

func verifyDeleteOrderRequest(data models.DeleteOrderRequest) *models.ServiceError {
	validator := validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateID(data.TargetID, "TargetID")

	if err := validator.GetFirstError(); err != nil {
		return err
	}

	if data.UserRole != "admin" && data.UserID != data.TargetID {
		return utils.NewNotAllowedRouteError()
	}

	return nil
}
