package picture

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/utils"
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
		return utils.NewInvalidImageFormatError()
	}

	if data.Priority != nil {
		validator = validator.ValidatePositiveNumber(int64(*data.Priority), "Priority")
	}

	return validator.GetFirstError()
}

func verifyUpdatePicture(data models.UpdatePictureRequest) *models.ServiceError {
	validator := validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateID(data.TargetID, "TargetID").
		ValidateAdminRole(data.UserRole).
		ValidateBucketRestriction(data.PublicBucket, data.PrivateBucket)

	if data.Name != nil {
		validator = validator.ValidateNotEmpty(data.Name, "Name")
	}

	if data.Priority != nil {
		validator = validator.ValidatePositiveNumber(int64(*data.Priority), "Priority")
	}

	return validator.GetFirstError()
}

func verifyDeletePicture(data models.DeletePictureRequest) *models.ServiceError {
	return validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateID(data.TargetID, "TargetID").
		ValidateAdminRole(data.UserRole).
		ValidateBucketRestriction(data.PublicBucket, data.PrivateBucket).
		GetFirstError()
}
