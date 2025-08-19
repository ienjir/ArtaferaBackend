package picture

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/validation"
	"net/http"
)

func verifyGetPictureByIDRequest(data models.GetPictureByIDRequest) *models.ServiceError {
	return validation.NewValidator().
		ValidateID(data.TargetID, "TargetID").
		ValidateBucketRestriction(data.PublicBucket, data.PrivateBucket).
		GetFirstError()
}

func verifyGetPictureByNameRequest(data models.GetPictureByNameRequest) *models.ServiceError {
	return validation.NewValidator().
		ValidateNotEmpty(&data.Name, "Name").
		ValidateBucketRestriction(data.PublicBucket, data.PrivateBucket).
		GetFirstError()
}

func verifyListPicture(data models.ListPictureRequest) *models.ServiceError {
	return validation.NewValidator().
		ValidateBucketRestriction(data.PublicBucket, data.PrivateBucket).
		ValidateID(data.UserID, "UserID").
		ValidateOffset(int64(data.Offset)).
		ValidateAdminRole(data.UserRole).
		GetFirstError()
}

func verifyCreatePicture(data models.CreatePictureRequest) *models.ServiceError {
	validator := validation.NewValidator().
		ValidateAdminRole(data.UserRole).
		ValidateID(data.UserID, "UserID").
		ValidateBucketRestriction(data.PublicBucket, data.PrivateBucket)

	if !validation.IsValidImage(&data.Picture) {
		return &models.ServiceError{StatusCode: http.StatusBadRequest, Message: "Invalid image format"}
	}

	if data.Priority != nil {
		validator = validator.ValidatePositiveNumber(int64(*data.Priority), "Priority")
	}

	return validator.GetFirstError()
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

func verifyDeletePicture(data models.DeletePictureRequest) *models.ServiceError {
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
			Message:    "You are not allowed to see this route",
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
