package validation

import (
	"github.com/ienjir/ArtaferaBackend/src/api/auth"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/nyaruka/phonenumbers"
	passwordvalidator "github.com/wagslane/go-password-validator"
	"net/http"
	"strings"
)

func validatePassword(password string) *models.ServiceError {
	if password == "" {
		return &models.ServiceError{StatusCode: http.StatusUnprocessableEntity, Message: "Password can't be empty"}
	}

	if err := passwordvalidator.Validate(password, auth.MinEntropyBits); err != nil {
		return &models.ServiceError{StatusCode: http.StatusForbidden, Message: "Password is insecure"}
	}

	return nil
}

func validateEmail(email string) *models.ServiceError {
	if email == "" {
		return &models.ServiceError{StatusCode: http.StatusUnprocessableEntity, Message: "Email can't be empty"}
	}

	// Add additional email format validation logic if necessary
	return nil
}

func validateName(name, fieldName string) *models.ServiceError {
	if name == "" {
		return &models.ServiceError{StatusCode: http.StatusUnprocessableEntity, Message: fieldName + " can't be empty"}
	}

	return nil
}

func validatePhone(phone, phoneRegion *string) *models.ServiceError {
	if phone == nil && phoneRegion == nil {
		return nil // Phone is optional if both are nil
	}

	if phoneRegion == nil {
		return &models.ServiceError{StatusCode: http.StatusUnprocessableEntity, Message: "Phone region has to be sent"}
	}

	if phone == nil || *phone == "" {
		return &models.ServiceError{StatusCode: http.StatusUnprocessableEntity, Message: "Phone number can't be empty"}
	}

	if *phoneRegion == "" {
		return &models.ServiceError{StatusCode: http.StatusUnprocessableEntity, Message: "Phone region can't be empty"}
	}

	upperRegion := strings.ToUpper(*phoneRegion)
	parsedNumber, err := phonenumbers.Parse(*phone, upperRegion)
	if err != nil {
		return &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: "Error while trying to parse phone number"}
	}

	if !phonenumbers.IsValidNumber(parsedNumber) {
		return &models.ServiceError{StatusCode: http.StatusUnprocessableEntity, Message: "Phone format is not valid"}
	}

	return nil
}

func validateAddress(field *string, fieldName string) *models.ServiceError {
	if field != nil && *field == "" {
		return &models.ServiceError{StatusCode: http.StatusUnprocessableEntity, Message: fieldName + " can't be empty"}
	}

	return nil
}
