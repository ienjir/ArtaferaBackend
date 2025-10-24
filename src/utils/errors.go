package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/models"
)

// Common error messages
const (
	// Generic errors
	ErrInternalServer   = "Internal server error"
	ErrInvalidJSON      = "Invalid JSON format"
	ErrInvalidID        = "Invalid ID format"
	ErrResourceNotFound = "Resource not found"
	ErrUnauthorized     = "Unauthorized access"
	ErrForbidden        = "Access forbidden"
	ErrBadRequest       = "Bad request"

	// Field validation errors
	ErrFieldRequired   = " is required"
	ErrFieldEmpty      = " cannot be empty"
	ErrFieldNotEmpty   = " has to be empty"
	ErrFieldInvalid    = " is invalid"
	ErrFieldTooLong    = " is too long"
	ErrFieldTooShort   = " is too short"
	ErrFieldOutOfRange = " is out of range"

	// Authentication errors
	ErrInvalidCredentials   = "Invalid email or password"
	ErrTokenExpired         = "Token has expired"
	ErrTokenInvalid         = "Invalid token"
	ErrRefreshTokenRequired = "Refresh token is required"

	// Authorization errors
	ErrAdminRequired     = "Admin role required"
	ErrOwnerAccess       = "You can only access your own resources"
	ErrInsufficientPerms = "Insufficient permissions"

	// Database errors
	ErrDatabaseConnection  = "Database connection failed"
	ErrDuplicateEntry      = "Resource already exists"
	ErrConstraintViolation = "Database constraint violation"

	// File/Upload errors
	ErrFileUpload        = "File upload failed"
	ErrInvalidFileFormat = "Invalid file format"
	ErrFileTooLarge      = "File size exceeds limit"

	// Business logic errors
	ErrPasswordInsecure = "Password does not meet security requirements"
	ErrEmailInvalid     = "Email format is invalid"
	ErrPhoneInvalid     = "Phone number format is invalid"
	ErrPasswordWrong    = "Password is incorrect"
	ErrEmailLowercase   = "Email must be lowercase"

	// Resource-specific errors
	ErrUserNotFound           = "User not found"
	ErrRoleNotFound           = "Role not found"
	ErrArtNotFound            = "Art not found"
	ErrOrderNotFound          = "Order not found"
	ErrPictureNotFound        = "Picture not found"
	ErrLanguageNotFound       = "Language not found"
	ErrSavedNotFound          = "Saved not found"
	ErrArtTranslationNotFound = "Art translation not found"

	// Conflict/Duplicate errors
	ErrUserAlreadyExists     = "Email already in use"
	ErrRoleAlreadyExists     = "Role already exists"
	ErrLanguageAlreadyExists = "Language already exists"
	ErrArtTranslationExists  = "Art translation already exists for this language"
	ErrArtAlreadySaved       = "Art is already saved for this user"

	// Business rule errors
	ErrArtNotAvailable      = "Art is not available"
	ErrOwnerOnlyOrders      = "You can only see orders for your own user account"
	ErrOwnerOnlyCreate      = "You can only create orders for your own user account"
	ErrOwnerOnlyAccess      = "Access denied: can only access your own saved items"
	ErrOwnerOnlySaved       = "You can only see saved for your own user account"
	ErrOwnerOnlyCreateSaved = "You can only create saved for your own user account"
	ErrOwnerOnlyUpdateSaved = "You can only update saved for your own user account"
	ErrNotAllowedRoute      = "You are not allowed to see this route"
	ErrAdminOnlyPictures    = "You can only update pictures as an admin"

	// Form/Input validation errors
	ErrPictureRequired       = "Picture is required"
	ErrInvalidImageFormat    = "Invalid image format"
	ErrInvalidPriorityFormat = "Invalid priority format"
	ErrInvalidPublicFormat   = "Invalid isPublic format"
	ErrNoContentFound        = "No content found"

	// Auth/Token errors
	ErrAccessTokenRequired = "Access token is required"
	ErrInvalidAuthHeader   = "Invalid authorization header format"
	ErrInvalidTokenClaims  = "Invalid token claims"
	ErrRoleNotInToken      = "Role not found in token"

	// Database operation errors
	ErrDatabaseRetrieval = "Error while retrieving data from database"
	ErrDatabaseUpdate    = "Error while updating data"
	ErrDatabaseDelete    = "Error while deleting data"
	ErrDatabaseCreate    = "Error while creating data"
	ErrDatabaseCount     = "Error while counting records"
	ErrTransactionStart  = "Failed to start transaction"
	ErrTransactionCommit = "Failed to commit transaction"
	ErrHashPassword      = "Failed to hash password"
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

func NewFieldNotEmptyError(fieldName string) *models.ServiceError {
	return NewBadRequestError(fieldName + ErrFieldNotEmpty)
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

// Resource not found errors
func NewUserNotFoundError() *models.ServiceError {
	return NewNotFoundError(ErrUserNotFound)
}

func NewRoleNotFoundError() *models.ServiceError {
	return NewNotFoundError(ErrRoleNotFound)
}

func NewArtNotFoundError() *models.ServiceError {
	return NewNotFoundError(ErrArtNotFound)
}

func NewOrderNotFoundError() *models.ServiceError {
	return NewNotFoundError(ErrOrderNotFound)
}

func NewPictureNotFoundError() *models.ServiceError {
	return NewNotFoundError(ErrPictureNotFound)
}

func NewLanguageNotFoundError() *models.ServiceError {
	return NewNotFoundError(ErrLanguageNotFound)
}

func NewSavedNotFoundError() *models.ServiceError {
	return NewNotFoundError(ErrSavedNotFound)
}

func NewArtTranslationNotFoundError() *models.ServiceError {
	return NewNotFoundError(ErrArtTranslationNotFound)
}

// Conflict/Duplicate errors
func NewUserAlreadyExistsError() *models.ServiceError {
	return NewConflictError(ErrUserAlreadyExists)
}

func NewRoleAlreadyExistsError() *models.ServiceError {
	return NewConflictError(ErrRoleAlreadyExists)
}

func NewLanguageAlreadyExistsError() *models.ServiceError {
	return NewConflictError(ErrLanguageAlreadyExists)
}

func NewArtTranslationExistsError() *models.ServiceError {
	return NewConflictError(ErrArtTranslationExists)
}

func NewArtAlreadySavedError() *models.ServiceError {
	return NewConflictError(ErrArtAlreadySaved)
}

// Business rule errors
func NewArtNotAvailableError() *models.ServiceError {
	return NewBadRequestError(ErrArtNotAvailable)
}

func NewOwnerOnlyOrdersError() *models.ServiceError {
	return NewForbiddenError(ErrOwnerOnlyOrders)
}

func NewOwnerOnlyCreateError() *models.ServiceError {
	return NewForbiddenError(ErrOwnerOnlyCreate)
}

func NewOwnerOnlyAccessError() *models.ServiceError {
	return NewForbiddenError(ErrOwnerOnlyAccess)
}

func NewOwnerOnlySavedError() *models.ServiceError {
	return NewForbiddenError(ErrOwnerOnlySaved)
}

func NewOwnerOnlyCreateSavedError() *models.ServiceError {
	return NewForbiddenError(ErrOwnerOnlyCreateSaved)
}

func NewOwnerOnlyUpdateSavedError() *models.ServiceError {
	return NewForbiddenError(ErrOwnerOnlyUpdateSaved)
}

func NewNotAllowedRouteError() *models.ServiceError {
	return NewForbiddenError(ErrNotAllowedRoute)
}

func NewAdminOnlyPicturesError() *models.ServiceError {
	return NewForbiddenError(ErrAdminOnlyPictures)
}

// Form/Input validation errors
func NewPictureRequiredError() *models.ServiceError {
	return NewBadRequestError(ErrPictureRequired)
}

func NewInvalidImageFormatError() *models.ServiceError {
	return NewBadRequestError(ErrInvalidImageFormat)
}

func NewInvalidPriorityFormatError() *models.ServiceError {
	return NewBadRequestError(ErrInvalidPriorityFormat)
}

func NewInvalidPublicFormatError() *models.ServiceError {
	return NewBadRequestError(ErrInvalidPublicFormat)
}

func NewNoContentFoundError() *models.ServiceError {
	return NewBadRequestError(ErrNoContentFound)
}

// Auth/Token errors
func NewAccessTokenRequiredError() *models.ServiceError {
	return NewUnauthorizedError(ErrAccessTokenRequired)
}

func NewInvalidAuthHeaderError() *models.ServiceError {
	return NewUnauthorizedError(ErrInvalidAuthHeader)
}

func NewInvalidTokenClaimsError() *models.ServiceError {
	return NewUnauthorizedError(ErrInvalidTokenClaims)
}

func NewRoleNotInTokenError() *models.ServiceError {
	return NewUnauthorizedError(ErrRoleNotInToken)
}

func NewPasswordWrongError() *models.ServiceError {
	return NewUnauthorizedError(ErrPasswordWrong)
}

func NewEmailLowercaseError() *models.ServiceError {
	return NewBadRequestError(ErrEmailLowercase)
}

// Database operation errors
func NewDatabaseRetrievalError() *models.ServiceError {
	return NewInternalServerError(ErrDatabaseRetrieval)
}

func NewDatabaseUpdateError() *models.ServiceError {
	return NewInternalServerError(ErrDatabaseUpdate)
}

func NewDatabaseDeleteError() *models.ServiceError {
	return NewInternalServerError(ErrDatabaseDelete)
}

func NewDatabaseCreateError() *models.ServiceError {
	return NewInternalServerError(ErrDatabaseCreate)
}

func NewDatabaseCountError() *models.ServiceError {
	return NewInternalServerError(ErrDatabaseCount)
}

func NewTransactionStartError() *models.ServiceError {
	return NewInternalServerError(ErrTransactionStart)
}

func NewTransactionCommitError() *models.ServiceError {
	return NewInternalServerError(ErrTransactionCommit)
}

func NewHashPasswordError() *models.ServiceError {
	return NewInternalServerError(ErrHashPassword)
}

// Generic record not found error
func NewRecordNotFoundError() *models.ServiceError {
	return NewNotFoundError(ErrResourceNotFound)
}
