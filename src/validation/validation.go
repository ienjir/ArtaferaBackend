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

type Validator struct {
	errors []*models.ServiceError
}

func NewValidator() *Validator {
	return &Validator{errors: make([]*models.ServiceError, 0)}
}

func (v *Validator) ValidateID(id int64, fieldName string) *Validator {
	if id < 1 {
		v.errors = append(v.errors, &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    fieldName + " must be at least 1",
		})
	}
	return v
}

func (v *Validator) ValidateIntID(id int, fieldName string) *Validator {
	if id < 1 {
		v.errors = append(v.errors, &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    fieldName + " must be at least 1",
		})
	}
	return v
}

func (v *Validator) ValidatePositiveFloat(value *float64, fieldName string) *Validator {
	if value != nil && *value < 0 {
		v.errors = append(v.errors, &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    fieldName + " cannot be negative",
		})
	}
	return v
}

func (v *Validator) ValidatePositiveNumber(value int64, fieldName string) *Validator {
	if value < 0 {
		v.errors = append(v.errors, &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    fieldName + " cannot be negative",
		})
	}
	return v
}

func (v *Validator) ValidateRange(value *int, min, max int, fieldName string) *Validator {
	if value != nil && (*value < min || *value > max) {
		v.errors = append(v.errors, &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    fieldName + " must be between " + strconv.Itoa(min) + " and " + strconv.Itoa(max),
		})
	}
	return v
}

func (v *Validator) ValidateIntRange(value int, min, max int, fieldName string) *Validator {
	if value < min || value > max {
		v.errors = append(v.errors, &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    fieldName + " must be between " + strconv.Itoa(min) + " and " + strconv.Itoa(max),
		})
	}
	return v
}

func (v *Validator) ValidateOffset(offset int64) *Validator {
	if offset < 0 {
		v.errors = append(v.errors, &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "Offset must be 0 or greater",
		})
	}
	return v
}

func (v *Validator) ValidatePageSize(pageSize int) *Validator {
	if pageSize < 1 || pageSize > 100 {
		v.errors = append(v.errors, &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "PageSize must be between 1 and 100",
		})
	}
	return v
}

func (v *Validator) ValidateSortOrder(sortOrder *string) *Validator {
	if sortOrder != nil && *sortOrder != "asc" && *sortOrder != "desc" {
		v.errors = append(v.errors, &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "SortOrder must be 'asc' or 'desc'",
		})
	}
	return v
}

func (v *Validator) ValidateAdminRole(userRole string) *Validator {
	if userRole != "admin" {
		v.errors = append(v.errors, &models.ServiceError{
			StatusCode: http.StatusForbidden,
			Message:    "Admin role required",
		})
	}
	return v
}

func (v *Validator) ValidateUserAccess(userID, targetID int64, userRole string) *Validator {
	if userRole != "admin" && userID != targetID {
		v.errors = append(v.errors, &models.ServiceError{
			StatusCode: http.StatusForbidden,
			Message:    "Access denied - you can only access your own resources",
		})
	}
	return v
}

func (v *Validator) ValidateNotEmpty(value *string, fieldName string) *Validator {
	if value != nil && *value == "" {
		v.errors = append(v.errors, &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    fieldName + " cannot be empty",
		})
	}
	return v
}

func (v *Validator) ValidateBucketRestriction(publicBucket, privateBucket string) *Validator {
	if publicBucket != "" || privateBucket != "" {
		v.errors = append(v.errors, &models.ServiceError{
			StatusCode: http.StatusForbidden,
			Message:    "Bucket names are not allowed to be specified",
		})
	}
	return v
}

func (v *Validator) GetFirstError() *models.ServiceError {
	if len(v.errors) > 0 {
		return v.errors[0]
	}
	return nil
}

func (v *Validator) HasErrors() bool {
	return len(v.errors) > 0
}

func ValidateIDField(id int64, fieldName string) *models.ServiceError {
	if id < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    fieldName + " must be at least 1",
		}
	}
	return nil
}

func ValidateAdminRole(userRole string) *models.ServiceError {
	if userRole != "admin" {
		return &models.ServiceError{
			StatusCode: http.StatusForbidden,
			Message:    "Admin role required",
		}
	}
	return nil
}

func ValidateUserAccess(userID, targetID int64, userRole string) *models.ServiceError {
	if userRole != "admin" && userID != targetID {
		return &models.ServiceError{
			StatusCode: http.StatusForbidden,
			Message:    "Access denied - you can only access your own resources",
		}
	}
	return nil
}
