package user

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/validation"
	"net/http"
)

func verifyCreateUserData(data models.CreateUserRequest) *models.ServiceError {

	if err := validation.ValidatePassword(data.Password); err != nil {
		return err
	}

	if err := validation.ValidateEmail(data.Email); err != nil {
		return err
	}

	if err := validation.ValidateName(data.Firstname, "Firstname"); err != nil {
		return err
	}

	if err := validation.ValidateName(data.Lastname, "Lastname"); err != nil {
		return err
	}

	if err := validation.ValidatePhone(data.Phone, data.PhoneRegion); err != nil {
		return err
	}

	if err := validation.ValidateAddress(data.Address1, "Address1"); err != nil {
		return err
	}

	if err := validation.ValidateAddress(data.Address2, "Address2"); err != nil {
		return err
	}

	if err := validation.ValidateAddress(data.City, "City"); err != nil {
		return err
	}

	if err := validation.ValidateAddress(data.PostalCode, "Postal code"); err != nil {
		return err
	}

	return nil
}

func verifyListUserData(data models.ListUserRequest) *models.ServiceError {
	if data.UserID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "UserID has to be at least 1",
		}
	}

	if data.Offset < 0 {
		return &models.ServiceError{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "Offset can't be less than 0",
		}
	}

	if data.UserID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "UserID has to be at least 1",
		}
	}

	return nil
}

func verifyDeleteUserRequest(data models.DeleteUserRequest) *models.ServiceError {
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
				Message:    "You can only create orders for your own user account",
			}
		}
	}

	return nil
}

func verifyGetUserById(data models.GetUserByIDRequest) *models.ServiceError {
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

func verifyGetUserByEmail(data models.GetUserByEmailRequest) *models.ServiceError {
	if data.UserID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "UserID has to be at least 1",
		}
	}

	if err := validation.ValidateEmail(data.Email); err != nil {
		return err
	}

	return nil
}

func ValidateUpdateUserRequest(data models.UpdateUserRequest) *models.ServiceError {
	if data.UserRole != "admin" {
		if data.UserID != data.TargetID {
			return &models.ServiceError{
				StatusCode: http.StatusForbidden,
				Message:    "You are not allowed to see this route",
			}
		}
	}
	
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

	if data.Firstname != nil {
		if err := validation.ValidateName(*data.Firstname, "Firstname"); err != nil {
			return err
		}
	}

	if data.Lastname != nil {
		if err := validation.ValidateName(*data.Lastname, "Lastname"); err != nil {
			return err
		}
	}

	if data.Password != nil {
		if err := validation.ValidatePassword(*data.Password); err != nil {
			return err
		}
	}

	if data.Email != nil {
		if err := validation.ValidateEmail(*data.Email); err != nil {
			return err
		}
	}

	if data.Phone != nil {
		if err := validation.ValidatePhone(data.Phone, data.PhoneRegion); err != nil {
			return err
		}
	}

	if data.Address1 != nil {
		if err := validation.ValidateAddress(data.Address1, "Address1"); err != nil {
			return err
		}
	}

	if data.Address2 != nil {
		if err := validation.ValidateAddress(data.Address2, "Address2"); err != nil {
			return err
		}
	}

	if data.City != nil {
		if err := validation.ValidateAddress(data.City, "City"); err != nil {
			return err
		}
	}

	if data.PostalCode != nil {
		if err := validation.ValidateAddress(data.PostalCode, "Postal code"); err != nil {
			return err
		}
	}

	if data.Password != nil {
		if err := validation.ValidatePassword(*data.Password); err != nil {
			return err
		}
	}

	if data.RoleID != nil {
		if data.TargetID < 1 {
			return &models.ServiceError{
				StatusCode: http.StatusBadRequest,
				Message:    "RoleID has to be at least 1",
			}
		}
	}

	return nil
}
