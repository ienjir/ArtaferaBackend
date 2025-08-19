package art

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"net/http"
)

func verifyGetArtByID(data models.GetArtByIDRequest) *models.ServiceError {
	if data.UserID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "UserID has to be over 1",
		}
	}

	if data.TargetID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "TargetID has to be over 1",
		}
	}

	return nil
}

func verifyListArt(data models.ListArtRequest) *models.ServiceError {
	if data.UserID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "UserID has to be over 1",
		}
	}

	if data.Page < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "Page has to be over 0",
		}
	}

	if data.PageSize < 1 || data.PageSize > 100 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "PageSize has to be between 1 and 100",
		}
	}

	if data.SortOrder != nil && *data.SortOrder != "asc" && *data.SortOrder != "desc" {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "SortOrder must be 'asc' or 'desc'",
		}
	}

	return nil
}

func verifyCreateArt(data models.CreateArtRequest) *models.ServiceError {
	if data.UserID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "UserID has to be over 1",
		}
	}

	if data.UserRole != "admin" {
		return &models.ServiceError{
			StatusCode: http.StatusUnauthorized,
			Message:    "Only admins can create art",
		}
	}

	if data.Price < 0 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "Price cannot be negative",
		}
	}

	if data.CurrencyID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "CurrencyID has to be over 1",
		}
	}

	if data.CreationYear < 1000 || data.CreationYear > 9999 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "CreationYear must be between 1000 and 9999",
		}
	}

	return nil
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
	if data.UserID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "UserID has to be over 1",
		}
	}

	if data.UserRole != "admin" {
		return &models.ServiceError{
			StatusCode: http.StatusUnauthorized,
			Message:    "Only admins can delete art",
		}
	}

	if data.TargetID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "TargetID has to be over 1",
		}
	}

	return nil
}
