package saved

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"net/http"
)

func verifyGetSavedById(data models.GetSavedByIDRequest) *models.ServiceError {
	if data.UserID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "UserID has to be at least 1",
		}
	}

	if data.TargetID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "TargetID has to be at least 1",
		}
	}

	if data.UserRole != "admin" {
		if data.UserID != data.TargetID {
			return &models.ServiceError{
				StatusCode: http.StatusForbidden,
				Message:    "You are not allowed to see this route",
			}
		}
	}

	return nil
}

func verifyGetSavedForUserRequest(data models.GetSavedForUserRequest) *models.ServiceError {
	if data.UserID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "UserID has to be at least 1",
		}
	}

	if data.TargetUserID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "OrderID has to be at least 1",
		}
	}

	if data.Offset < 0 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "Offset has to  be 0 or more",
		}
	}

	if data.UserRole != "admin" {
		if data.UserID != data.TargetUserID {
			return &models.ServiceError{
				StatusCode: http.StatusForbidden,
				Message:    "You can only see saved for your own user account",
			}
		}
	}

	return nil
}

func verifyListSavedRequest(data models.ListSavedRequest) *models.ServiceError {
	if data.UserID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "UserID has to be at least 1",
		}
	}

	if data.Offset < 0 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "Offset has to be 0 or more",
		}
	}

	if data.UserRole != "admin" {
		return &models.ServiceError{
			StatusCode: http.StatusForbidden,
			Message:    "You are not allowed for this route",
		}
	}

	return nil
}

func verifyCreateSaved(data models.CreateSavedRequest) *models.ServiceError {
	if data.TargetUserID == nil {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "UserID is required",
		}
	}

	if *data.TargetUserID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "UserID has to be over 1",
		}
	}

	if data.ArtID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "ArtID has to be over 1",
		}
	}

	if data.UserRole != "admin" {
		if *data.TargetUserID != data.UserID {
			return &models.ServiceError{
				StatusCode: http.StatusForbidden,
				Message:    "You can only create saved for your own user account",
			}
		}
	}

	return nil
}

func verifyUpdateSavedRequest(data models.UpdateSavedRequest) *models.ServiceError {
	if data.UserID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "UserID has to be at least 1",
		}
	}

	if data.TargetID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "TargetID has to be at least 1",
		}
	}

	if data.TargetUserID != nil {
		if *data.TargetUserID < 1 {
			return &models.ServiceError{
				StatusCode: http.StatusBadRequest,
				Message:    "TargetUserID has to be at least 1",
			}
		}
	}

	if data.ArtID != nil {
		if *data.ArtID < 1 {
			return &models.ServiceError{
				StatusCode: http.StatusBadRequest,
				Message:    "ArtID has to be over 1",
			}
		}
	}

	if data.UserRole != "admin" {
		if data.UserID != *data.TargetUserID {
			return &models.ServiceError{
				StatusCode: http.StatusForbidden,
				Message:    "You can only see saved for your own user account",
			}
		}
	}

	return nil
}

func verifyDeleteSavedRequest(data models.DeleteSavedRequest) *models.ServiceError {
	if data.UserID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "UserID has to be at least 1",
		}
	}

	if data.TargetID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "TargetID has to be at least 1",
		}
	}

	return nil
}
