package user

import (
	"github.com/go-playground/validator/v10"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/validation"
	"net/http"
	"strconv"
)

func VerifyCreateUserData(Data models.CreateUserRequest) *models.ServiceError {
	if err := validation.ValidatePassword(Data.Password); err != nil {
		return err
	}

	if err := validation.ValidateEmail(Data.Email); err != nil {
		return err
	}

	if err := validation.ValidateName(Data.Firstname, "Firstname"); err != nil {
		return err
	}

	if err := validation.ValidateName(Data.Lastname, "Lastname"); err != nil {
		return err
	}

	if err := validation.ValidatePhone(Data.Phone, Data.PhoneRegion); err != nil {
		return err
	}

	if err := validation.ValidateAddress(Data.Address1, "Address1"); err != nil {
		return err
	}

	if err := validation.ValidateAddress(Data.Address2, "Address2"); err != nil {
		return err
	}

	if err := validation.ValidateAddress(Data.City, "City"); err != nil {
		return err
	}

	if err := validation.ValidateAddress(Data.PostalCode, "Postal code"); err != nil {
		return err
	}

	return nil
}

func VerifyListUserData(Data models.ListUserRequest) *models.ServiceError {

	if Data.Offset < 0 {
		return &models.ServiceError{StatusCode: http.StatusUnprocessableEntity, Message: "Offset can't be less than 0"}
	}

	return nil
}

func VerifyDeleteUserRequest(requestUserID float64, requestUserRole, targetUserID string) *models.ServiceError {
	targetUserIDFloat, err := strconv.ParseFloat(targetUserID, 64)
	if err != nil {
		return &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: "Error while parsing userID"}
	}

	if requestUserRole != "admin" && requestUserID != targetUserIDFloat {
		return &models.ServiceError{StatusCode: http.StatusUnauthorized, Message: "You can only delete your own account"}
	}

	return nil
}

func VerifyGetUserById(requestUserID int64, requestUserRole, targetUserID string) *models.ServiceError {
	targetUserIDFloat, err := strconv.ParseInt(targetUserID, 10, 64)
	if err != nil {
		return &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: "Error while parsing userID"}
	}

	if requestUserRole != "admin" && requestUserID != targetUserIDFloat {
		return &models.ServiceError{StatusCode: http.StatusUnauthorized, Message: "You can only see your own account"}
	}

	return nil
}

func VerifyGetUserByEmail(Data models.GetUserByEmail) *models.ServiceError {

	if err := validation.ValidateEmail(Data.Email); err != nil {
		return &models.ServiceError{StatusCode: err.StatusCode, Message: err.Message}
	}

	return nil
}

func ValidateUpdateUserRequest(req models.UpdateUserRequest) error {
	validate := validator.New()
	return validate.Struct(req)
}
