package saved

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/utils"
	"github.com/ienjir/ArtaferaBackend/src/validation"
)

func verifyGetSavedById(data models.GetSavedByIDRequest) *models.ServiceError {
	validator := validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateID(data.TargetID, "TargetID")

	if data.UserRole != "admin" {
		if data.UserID != data.TargetID {
			return utils.NewOwnerOnlyAccessError()
		}
	}

	return validator.GetFirstError()
}

func verifyGetSavedForUserRequest(data models.GetSavedForUserRequest) *models.ServiceError {
	validator := validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateID(data.TargetUserID, "TargetUserID").
		ValidateOffset(int64(data.Offset))

	if data.UserRole != "admin" {
		if data.UserID != data.TargetUserID {
			return utils.NewOwnerOnlySavedError()
		}
	}

	return validator.GetFirstError()
}

func verifyListSavedRequest(data models.ListSavedRequest) *models.ServiceError {
	return validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateOffset(int64(data.Offset)).
		ValidateAdminRole(data.UserRole).
		GetFirstError()
}

func verifyCreateSaved(data models.CreateSavedRequest) *models.ServiceError {
	validator := validation.NewValidator().
		ValidateID(*data.TargetUserID, "TargetUserID").
		ValidateID(data.ArtID, "ArtID")

	if data.UserRole != "admin" && data.TargetUserID != nil {
		if data.UserID != *data.TargetUserID {
			return utils.NewOwnerOnlyCreateSavedError()
		}
	}

	return validator.GetFirstError()
}

func verifyUpdateSavedRequest(data models.UpdateSavedRequest) *models.ServiceError {
	validator := validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateID(data.TargetID, "TargetID")

	if data.TargetUserID != nil {
		validator = validator.ValidateID(*data.TargetUserID, "TargetUserID")
	}

	if data.ArtID != nil {
		validator = validator.ValidateID(*data.ArtID, "ArtID")
	}

	if data.UserRole != "admin" && data.TargetUserID != nil {
		if data.UserID != *data.TargetUserID {
			return utils.NewOwnerOnlyUpdateSavedError()
		}
	}

	return validator.GetFirstError()
}

func verifyDeleteSavedRequest(data models.DeleteSavedRequest) *models.ServiceError {
	return validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateID(data.TargetID, "TargetID").
		GetFirstError()
}
