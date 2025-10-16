package user

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/validation"
)

func verifyGetUserById(data models.GetUserByIDRequest) *models.ServiceError {
	return validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateID(data.TargetID, "TargetID").
		ValidateUserAccess(data.UserID, data.TargetID, data.UserRole).
		GetFirstError()
}

func verifyGetUserByEmail(data models.GetUserByEmailRequest) *models.ServiceError {
	v := validation.NewValidator().ValidateID(data.UserID, "UserID")
	if err := validation.ValidateEmail(data.Email); err != nil {
		return err
	}
	return v.GetFirstError()
}

func verifyListUserData(data models.ListUserRequest) *models.ServiceError {
	return validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateOffset(int64(data.Offset)).
		ValidateAdminRole(data.UserRole).
		GetFirstError()
}

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

func verifyUpdateUserRequest(data models.UpdateUserRequest) *models.ServiceError {
	validator := validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateID(data.TargetID, "TargetID").
		ValidateUserAccess(data.UserID, data.TargetID, data.UserRole)

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

	if data.RoleID != nil {
		validator = validator.ValidateID(*data.RoleID, "RoleID")
	}

	return validator.GetFirstError()
}

func verifyDeleteUserRequest(data models.DeleteUserRequest) *models.ServiceError {
	return validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateID(data.TargetID, "TargetID").
		ValidateUserAccess(data.UserID, data.TargetID, data.UserRole).
		GetFirstError()
}
