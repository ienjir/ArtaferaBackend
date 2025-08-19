package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/models"
)

// Common error messages
const (
	// Generic errors
	ErrInternalServer    = "Internal server error"
	ErrInvalidJSON       = "Invalid JSON format"
	ErrInvalidID         = "Invalid ID format"
	ErrResourceNotFound  = "Resource not found"
	ErrUnauthorized      = "Unauthorized access"
	ErrForbidden         = "Access forbidden"
	ErrBadRequest        = "Bad request"

	// Field validation errors
	ErrFieldRequired     = " is required"
	ErrFieldEmpty        = " cannot be empty"
	ErrFieldInvalid      = " is invalid"
	ErrFieldTooLong      = " is too long"
	ErrFieldTooShort     = " is too short"
	ErrFieldOutOfRange   = " is out of range"

	// Authentication errors
	ErrInvalidCredentials = "Invalid email or password"
	ErrTokenExpired       = "Token has expired"
	ErrTokenInvalid       = "Invalid token"
	ErrRefreshTokenRequired = "Refresh token is required"

	// Authorization errors
	ErrAdminRequired      = "Admin role required"
	ErrOwnerAccess        = "You can only access your own resources"
	ErrInsufficientPerms  = "Insufficient permissions"

	// Database errors
	ErrDatabaseConnection = "Database connection failed"
	ErrDuplicateEntry     = "Resource already exists"
	ErrConstraintViolation = "Database constraint violation"

	// File/Upload errors
	ErrFileUpload         = "File upload failed"
	ErrInvalidFileFormat  = "Invalid file format"
	ErrFileTooLarge       = "File size exceeds limit"

	// Business logic errors
	ErrPasswordInsecure   = "Password does not meet security requirements"
	ErrEmailInvalid       = "Email format is invalid"
	ErrPhoneInvalid       = "Phone number format is invalid"
)

// Error response structure
type ErrorResponse struct {
	Error   string `json:"error"`
	Code    int    `json:"code,omitempty"`
	Details string `json:"details,omitempty"`
}

// Success response structure
type SuccessResponse struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// RespondWithError sends a standardized error response
func RespondWithError(c *gin.Context, statusCode int, message string, details ...string) {
	response := ErrorResponse{
		Error: message,
		Code:  statusCode,
	}
	
	if len(details) > 0 && details[0] != "" {
		response.Details = details[0]
	}
	
	c.JSON(statusCode, response)
}

// RespondWithServiceError sends error response from models.ServiceError
func RespondWithServiceError(c *gin.Context, err *models.ServiceError) {
	if err == nil {
		RespondWithError(c, http.StatusInternalServerError, ErrInternalServer)
		return
	}
	RespondWithError(c, err.StatusCode, err.Message)
}

// RespondWithSuccess sends a standardized success response
func RespondWithSuccess(c *gin.Context, statusCode int, data interface{}, message ...string) {
	response := SuccessResponse{
		Data: data,
	}
	
	if len(message) > 0 && message[0] != "" {
		response.Message = message[0]
	}
	
	c.JSON(statusCode, response)
}

// Common error constructors
func NewBadRequestError(message string) *models.ServiceError {
	return &models.ServiceError{
		StatusCode: http.StatusBadRequest,
		Message:    message,
	}
}

func NewUnauthorizedError(message string) *models.ServiceError {
	return &models.ServiceError{
		StatusCode: http.StatusUnauthorized,
		Message:    message,
	}
}

func NewForbiddenError(message string) *models.ServiceError {
	return &models.ServiceError{
		StatusCode: http.StatusForbidden,
		Message:    message,
	}
}

func NewNotFoundError(message string) *models.ServiceError {
	return &models.ServiceError{
		StatusCode: http.StatusNotFound,
		Message:    message,
	}
}

func NewConflictError(message string) *models.ServiceError {
	return &models.ServiceError{
		StatusCode: http.StatusConflict,
		Message:    message,
	}
}

func NewUnprocessableEntityError(message string) *models.ServiceError {
	return &models.ServiceError{
		StatusCode: http.StatusUnprocessableEntity,
		Message:    message,
	}
}

func NewInternalServerError(message string) *models.ServiceError {
	return &models.ServiceError{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
	}
}

// Field validation error helpers
func NewFieldRequiredError(fieldName string) *models.ServiceError {
	return NewBadRequestError(fieldName + ErrFieldRequired)
}

func NewFieldEmptyError(fieldName string) *models.ServiceError {
	return NewBadRequestError(fieldName + ErrFieldEmpty)
}

func NewFieldInvalidError(fieldName string) *models.ServiceError {
	return NewBadRequestError(fieldName + ErrFieldInvalid)
}

func NewFieldOutOfRangeError(fieldName, rangeDesc string) *models.ServiceError {
	return NewBadRequestError(fieldName + " must be " + rangeDesc)
}

// Common business logic errors
func NewInvalidCredentialsError() *models.ServiceError {
	return NewUnauthorizedError(ErrInvalidCredentials)
}

func NewAdminRequiredError() *models.ServiceError {
	return NewForbiddenError(ErrAdminRequired)
}

func NewOwnerAccessError() *models.ServiceError {
	return NewForbiddenError(ErrOwnerAccess)
}

func NewPasswordInsecureError() *models.ServiceError {
	return NewUnprocessableEntityError(ErrPasswordInsecure)
}

func NewEmailInvalidError() *models.ServiceError {
	return NewUnprocessableEntityError(ErrEmailInvalid)
}

func NewPhoneInvalidError() *models.ServiceError {
	return NewUnprocessableEntityError(ErrPhoneInvalid)
}

// Parsing error helpers
func NewInvalidIDError() *models.ServiceError {
	return NewBadRequestError(ErrInvalidID)
}

func NewInvalidJSONError() *models.ServiceError {
	return NewBadRequestError(ErrInvalidJSON)
}