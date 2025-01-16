package validation

import "github.com/ienjir/ArtaferaBackend/src/models"

func VerifyCreateUserData(UserData models.CreateUserRequest) *models.ServiceError {
	if err := validatePassword(UserData.Password); err != nil {
		return err
	}

	if err := validateEmail(UserData.Email); err != nil {
		return err
	}

	if err := validateName(UserData.Firstname, "Firstname"); err != nil {
		return err
	}

	if err := validateName(UserData.Lastname, "Lastname"); err != nil {
		return err
	}

	if err := validatePhone(UserData.Phone, UserData.PhoneRegion); err != nil {
		return err
	}

	if err := validateAddress(UserData.Address1, "Address1"); err != nil {
		return err
	}

	if err := validateAddress(UserData.Address2, "Address2"); err != nil {
		return err
	}

	if err := validateAddress(UserData.City, "City"); err != nil {
		return err
	}

	if err := validateAddress(UserData.PostalCode, "Postal code"); err != nil {
		return err
	}

	return nil
}
