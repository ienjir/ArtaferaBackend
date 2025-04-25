package validation

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/nyaruka/phonenumbers"
	passwordvalidator "github.com/wagslane/go-password-validator"
	"log"
	"mime/multipart"
	"net/http"
	"net/mail"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var MinEntropyBits float64

func LoadsValidationEnvs() {
	minEntropyBits, err := strconv.ParseFloat(os.Getenv("ENTROPY_MIN_BITS"), 64)
	if err != nil {
		log.Fatal("Failed to load minimal entropy bits: " + err.Error())
		return
	}

	MinEntropyBits = minEntropyBits

	return
}

func ValidatePassword(password string) *models.ServiceError {
	if err := ValidatePasswordWithoutEntropy(password); err != nil {
		return &models.ServiceError{StatusCode: err.StatusCode, Message: err.Message}
	}

	if err := passwordvalidator.Validate(password, MinEntropyBits); err != nil {
		return &models.ServiceError{StatusCode: http.StatusForbidden, Message: "Password is insecure"}
	}

	return nil
}

func ValidatePasswordWithoutEntropy(password string) *models.ServiceError {
	if password == "" {
		return &models.ServiceError{StatusCode: http.StatusUnprocessableEntity, Message: "Password can't be empty"}
	}

	return nil
}

func ValidateEmail(email string) *models.ServiceError {
	if email == "" {
		return &models.ServiceError{StatusCode: http.StatusUnprocessableEntity, Message: "Email can't be empty"}
	}

	if IsLower(email) == false {
		return &models.ServiceError{StatusCode: http.StatusUnauthorized, Message: "Email has to be lowercase"}
	}

	_, err := mail.ParseAddress(email)
	if err != nil {
		return &models.ServiceError{StatusCode: http.StatusUnprocessableEntity, Message: "Error parsing email"}
	}

	return nil
}

func ValidateName(name, fieldName string) *models.ServiceError {
	if name == "" {
		return &models.ServiceError{StatusCode: http.StatusUnprocessableEntity, Message: fieldName + " can't be empty"}
	}

	return nil
}

func ValidatePhone(phone, phoneRegion *string) *models.ServiceError {
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

func ValidateAddress(field *string, fieldName string) *models.ServiceError {
	if field != nil && *field == "" {
		return &models.ServiceError{StatusCode: http.StatusUnprocessableEntity, Message: fieldName + " can't be empty"}
	}

	return nil
}

func ValidateStatusString(status string) (models.OrderStatus, *models.ServiceError) {
	switch status {
	case string(models.OrderStatusPending):
		return models.OrderStatusPending, nil
	case string(models.OrderStatusPaid):
		return models.OrderStatusPaid, nil
	case string(models.OrderStatusShipped):
		return models.OrderStatusShipped, nil
	case string(models.OrderStatusDelivered):
		return models.OrderStatusDelivered, nil
	case string(models.OrderStatusCancelled):
		return models.OrderStatusCancelled, nil
	default:
		return "", &models.ServiceError{StatusCode: http.StatusBadRequest, Message: "Order status is invalid"}
	}
}

func IsValidImage(file *multipart.FileHeader) bool {
	allowedExtensions := []string{".jpg", ".jpeg", ".png", ".gif"}
	ext := strings.ToLower(file.Filename[strings.LastIndex(file.Filename, "."):])
	for _, validExt := range allowedExtensions {
		if ext == validExt {
			return true
		}
	}
	return false
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
