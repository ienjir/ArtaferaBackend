package picture

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/validation"
	"net/http"
)

func verifyGetPictureByIDRequest(data models.GetPictureByIDRequest) *models.ServiceError {
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

	if data.BucketName != "" {
		return &models.ServiceError{
			StatusCode: http.StatusForbidden,
			Message:    "You are not allowed to send with a bucket name",
		}
	}

	if data.UserRole != "admin" {
		return &models.ServiceError{
			StatusCode: http.StatusForbidden,
			Message:    "You are not allowed to see this route",
		}
	}

	return nil
}

func verifyGetPictureByNameRequest(data models.GetPictureByNameRequest) *models.ServiceError {
	if data.BucketName != "" {
		return &models.ServiceError{
			StatusCode: http.StatusForbidden,
			Message:    "You are not allowed to send with a bucket name",
		}
	}

	if data.Name == "" {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "Name is required",
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

	/*
		if data.UserRole != "admin" {
			return &models.ServiceError{
				StatusCode: http.StatusForbidden,
				Message:    "Only admins can upload pictures",
			}
		}
	*/

	if !validation.IsValidImage(&data.Picture) {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid image format. Allowed formats: jpg, jpeg, png, gif",
		}
	}

	if data.BucketName != "" {
		return &models.ServiceError{
			StatusCode: http.StatusForbidden,
			Message:    "You are not allowed to send with a bucket name",
		}
	}

	return nil
}
