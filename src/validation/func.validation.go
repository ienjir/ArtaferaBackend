package validation

import "github.com/ienjir/ArtaferaBackend/src/models"

func VerifyLoginData(Data models.LoginRequest) *models.ServiceError {
	err := validatePassword(Data.Password)
	if err != nil {
		return err
	}

	err = validateEmail(Data.Email)
	if err != nil {
		return err
	}

	return nil
}

func VerifyCreateUserData(Data models.CreateUserRequest) *models.ServiceError {
	if err := validatePassword(Data.Password); err != nil {
		return err
	}

	if err := validateEmail(Data.Email); err != nil {
		return err
	}

	if err := validateName(Data.Firstname, "Firstname"); err != nil {
		return err
	}

	if err := validateName(Data.Lastname, "Lastname"); err != nil {
		return err
	}

	if err := validatePhone(Data.Phone, Data.PhoneRegion); err != nil {
		return err
	}

	if err := validateAddress(Data.Address1, "Address1"); err != nil {
		return err
	}

	if err := validateAddress(Data.Address2, "Address2"); err != nil {
		return err
	}

	if err := validateAddress(Data.City, "City"); err != nil {
		return err
	}

	if err := validateAddress(Data.PostalCode, "Postal code"); err != nil {
		return err
	}

	return nil
}
