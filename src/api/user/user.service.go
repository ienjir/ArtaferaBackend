package user

import (
	"errors"
	"github.com/ienjir/ArtaferaBackend/src/api/auth"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/validation"
	"gorm.io/gorm"
	"net/http"
)

func CreateUserService(request models.CreateUserRequest) (*models.User, *models.ServiceError) {
	var existingUser models.User

	// Check if email already exists
	if err := database.DB.Where("email = ?", request.Email).First(&existingUser).Error; err == nil {
		return nil, &models.ServiceError{StatusCode: http.StatusConflict, Message: "Email already in use"}
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: "Database error"}
	}

	// Hash the password
	hashedPassword, err := auth.HashPassword(request.Password)
	if err != nil {
		return nil, &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: "Failed to hash password"}
	}

	// Create user model
	user := &models.User{
		Firstname:   request.Firstname,
		Lastname:    request.Lastname,
		Email:       request.Email,
		Phone:       request.Phone,
		PhoneRegion: request.PhoneRegion,
		Address1:    request.Address1,
		Address2:    request.Address2,
		City:        request.City,
		PostalCode:  request.PostalCode,
		Password:    hashedPassword.Hash,
		Salt:        hashedPassword.Salt,
	}

	// Save user to the database
	if err := database.DB.Create(user).Error; err != nil {
		return nil, &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: "Failed to save user"}
	}

	return user, nil
}

func GetUserByEmail(email string) (*models.User, *models.ServiceError) {
	var user models.User

	err := database.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, &models.ServiceError{StatusCode: http.StatusNotFound, Message: "User with email not found"}
	}

	return &user, nil
}

/*
	Verify
*/

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
