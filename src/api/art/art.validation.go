package art

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/validation"
	"net/http"
)

func verifyGetArtByID(data models.GetArtByIDRequest) *models.ServiceError {
	return validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateID(data.TargetID, "TargetID").
		GetFirstError()
}

func verifyListArt(data models.ListArtRequest) *models.ServiceError {
	return validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateIntID(data.Page, "Page").
		ValidatePageSize(data.PageSize).
		ValidateSortOrder(data.SortOrder).
		GetFirstError()
}

func verifyCreateArt(data models.CreateArtRequest) *models.ServiceError {
	return validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateAdminRole(data.UserRole).
		ValidatePositiveNumber(data.Price, "Price").
		ValidateID(data.CurrencyID, "CurrencyID").
		ValidateIntRange(data.CreationYear, 1000, 9999, "CreationYear").
		GetFirstError()
}

func verifyUpdateArt(data models.UpdateArtRequest) *models.ServiceError {
	if data.UserID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "UserID has to be over 1",
		}
	}

	if data.UserRole != "admin" {
		return &models.ServiceError{
			StatusCode: http.StatusUnauthorized,
			Message:    "Only admins can update art",
		}
	}

	if data.TargetID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "TargetID has to be over 1",
		}
	}

	if data.Price != nil && *data.Price < 0 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "Price cannot be negative",
		}
	}

	if data.CurrencyID != nil && *data.CurrencyID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "CurrencyID has to be over 1",
		}
	}

	if data.CreationYear != nil && (*data.CreationYear < 1000 || *data.CreationYear > 9999) {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "CreationYear must be between 1000 and 9999",
		}
	}

	return nil
}

func verifyDeleteArt(data models.DeleteArtRequest) *models.ServiceError {
	return validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateAdminRole(data.UserRole).
		ValidateID(data.TargetID, "TargetID").
		GetFirstError()
}
