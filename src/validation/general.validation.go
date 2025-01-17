package validation

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/nyaruka/phonenumbers"
	passwordvalidator "github.com/wagslane/go-password-validator"
	"net/http"
	"net/mail"
	"os"
	"strconv"
	"strings"
)

var MinEntropyBits float64
var JWTSecret string

func LoadsAuthEnvs() error {
	minEntropyBits, err := strconv.ParseFloat(os.Getenv("ENTROPY_MIN_BITS"), 64)
	if err != nil {
		return err
	}

	MinEntropyBits = minEntropyBits

	JWTSecret = os.Getenv("JWT_SECRET")

	return nil
}

func validatePassword(password string) *models.ServiceError {
	if err := validatePasswordWithoutEntropy(password); err != nil {
		return &models.ServiceError{StatusCode: err.StatusCode, Message: err.Message}
	}

	if err := passwordvalidator.Validate(password, MinEntropyBits); err != nil {
		return &models.ServiceError{StatusCode: http.StatusForbidden, Message: "Password is insecure"}
	}

	return nil
}

func validatePasswordWithoutEntropy(password string) *models.ServiceError {
	if password == "" {
		return &models.ServiceError{StatusCode: http.StatusUnprocessableEntity, Message: "Password can't be empty"}
	}

	return nil
}

func validateEmail(email string) *models.ServiceError {
	if email == "" {
		return &models.ServiceError{StatusCode: http.StatusUnprocessableEntity, Message: "Email can't be empty"}
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return &models.ServiceError{StatusCode: http.StatusUnprocessableEntity, Message: "Error parsing email"}
	}

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
