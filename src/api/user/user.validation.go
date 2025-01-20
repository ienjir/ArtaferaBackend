package user

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/validation"
)

func VerifyCreateUserData(Data models.CreateUserRequest) *models.ServiceError {
	if err := validation.ValidatePassword(Data.Password); err != nil {
		return err
	}

	if err := validation.ValidateEmail(Data.Email); err != nil {
		return err
	}

	if err := validation.ValidateName(Data.Firstname, "Firstname"); err != nil {
		return err
	}

	if err := validation.ValidateName(Data.Lastname, "Lastname"); err != nil {
		return err
	}

	if err := validation.ValidatePhone(Data.Phone, Data.PhoneRegion); err != nil {
		return err
	}

	if err := validation.ValidateAddress(Data.Address1, "Address1"); err != nil {
		return err
	}

	if err := validation.ValidateAddress(Data.Address2, "Address2"); err != nil {
		return err
	}

	if err := validation.ValidateAddress(Data.City, "City"); err != nil {
		return err
	}

	if err := validation.ValidateAddress(Data.PostalCode, "Postal code"); err != nil {
		return err
	}

	return nil
}
