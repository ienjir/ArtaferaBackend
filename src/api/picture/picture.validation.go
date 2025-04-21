package picture

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/validation"
	"net/http"
)

func verifyGetPictureByIDRequest(data models.GetPictureByIDRequest) *models.ServiceError {
	if data.TargetID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "TargetID has to be over 1",
		}
	}

	if data.PublicBucket != "" {
		return &models.ServiceError{
			StatusCode: http.StatusForbidden,
			Message:    "You are not allowed to send with a bucket name",
		}
	}

	if data.PrivateBucket != "" {
		return &models.ServiceError{
			StatusCode: http.StatusForbidden,
			Message:    "You are not allowed to send with a bucket name",
		}
	}

	return nil
}

func verifyGetPictureByNameRequest(data models.GetPictureByNameRequest) *models.ServiceError {
	if data.Name == "" {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "Name is required",
		}
	}

	if data.PublicBucket != "" {
		return &models.ServiceError{
			StatusCode: http.StatusForbidden,
			Message:    "You are not allowed to send with a bucket name",
		}
	}

	if data.PrivateBucket != "" {
		return &models.ServiceError{
			StatusCode: http.StatusForbidden,
			Message:    "You are not allowed to send with a bucket name",
		}
	}

	return nil
}

func verifyListPicture(data models.ListPictureRequest) *models.ServiceError {
	if data.PublicBucket != "" {
		return &models.ServiceError{
			StatusCode: http.StatusForbidden,
			Message:    "You are not allowed to send with a bucket name",
		}
	}

	if data.PrivateBucket != "" {
		return &models.ServiceError{
			StatusCode: http.StatusForbidden,
			Message:    "You are not allowed to send with a bucket name",
		}
	}

	if data.UserID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "UserID has to be over 1",
		}
	}

	if data.Offset < 0 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "Offset has to  be 0 or more",
		}
	}

	if data.UserRole != "admin" {
		return &models.ServiceError{
			StatusCode: http.StatusForbidden,
			Message:    "You are not authorized to see all role",
		}
	}

	return nil
}

func verifyCreatePicture(data models.CreatePictureRequest) *models.ServiceError {
	if data.Priority != nil {
		if *data.Priority < 1 {
			return &models.ServiceError{
				StatusCode: http.StatusBadRequest,
				Message:    "Priority must be greater than 0",
			}
		}
	}

	if data.UserRole != "admin" {
		return &models.ServiceError{
			StatusCode: http.StatusForbidden,
			Message:    "Only admins can upload pictures",
		}
	}

	if !validation.IsValidImage(&data.Picture) {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid image format. Allowed formats: jpg, jpeg, png, gif",
		}
	}

	if data.PublicBucket != "" {
		return &models.ServiceError{
			StatusCode: http.StatusForbidden,
			Message:    "You are not allowed to send with a bucket name",
		}
	}

	if data.PrivateBucket != "" {
		return &models.ServiceError{
			StatusCode: http.StatusForbidden,
			Message:    "You are not allowed to send with a bucket name",
		}
	}

	return nil
}

func verifyUpdatePicture(data models.UpdatePictureRequest) *models.ServiceError {

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
		return &models.ServiceError{
			StatusCode: http.StatusForbidden,
			Message:    "You can only update pictures as an admin",
		}
	}

	if data.Name != nil {
		if *data.Name == "" {
			return &models.ServiceError{
				StatusCode: http.StatusBadRequest,
				Message:    "Name can not be empty",
			}
		}
	}

	if data.Priority != nil {
		if *data.Priority <= 0 {
			return &models.ServiceError{
				StatusCode: http.StatusBadRequest,
				Message:    "Priority has to be 0 or more",
			}
		}
	}
	
	if data.PublicBucket != "" {
		return &models.ServiceError{
			StatusCode: http.StatusForbidden,
			Message:    "You are not allowed to send with a bucket name",
		}
	}

	if data.PrivateBucket != "" {
		return &models.ServiceError{
			StatusCode: http.StatusForbidden,
			Message:    "You are not allowed to send with a bucket name",
		}
	}

	return nil
}
