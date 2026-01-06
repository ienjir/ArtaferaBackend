package art

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/validation"
)

func verifyGetArtByID(data models.GetArtByIDRequest) *models.ServiceError {
	return validation.NewValidator().
		ValidateID(data.TargetID, "TargetID").
		GetFirstError()
}

func verifyListArt(data models.ListArtRequest) *models.ServiceError {
	return validation.NewValidator().
		ValidateOffset(int64(data.Offset)).
		GetFirstError()
}
func verifyListArtForArtPage(data models.ListArtForArtPageRequest) *models.ServiceError {
	return nil
}

func verifyListFeaturedArt(data models.ListFeaturedArtRequest) *models.ServiceError {
	validator := validation.NewValidator()
	if data.Limit > 0 {
		validator = validator.ValidateRange(&data.Limit, 1, 20, "Limit")
	}
	return validator.GetFirstError()
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
	validator := validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateAdminRole(data.UserRole).
		ValidateID(data.TargetID, "TargetID")

	if data.Price != nil {
		validator = validator.ValidatePositiveNumber(*data.Price, "Price")
	}

	if data.CurrencyID != nil {
		validator = validator.ValidateID(*data.CurrencyID, "CurrencyID")
	}

	if data.CreationYear != nil {
		validator = validator.ValidateIntRange(*data.CreationYear, 1000, 9999, "CreationYear")
	}

	return validator.GetFirstError()
}

func verifyDeleteArt(data models.DeleteArtRequest) *models.ServiceError {
	return validation.NewValidator().
		ValidateID(data.UserID, "UserID").
		ValidateAdminRole(data.UserRole).
		ValidateID(data.TargetID, "TargetID").
		GetFirstError()
}
