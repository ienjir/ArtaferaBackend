package user

import (
	"errors"
	"fmt"
	"github.com/ienjir/ArtaferaBackend/src/api/auth"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/nyaruka/phonenumbers"
	passwordvalidator "github.com/wagslane/go-password-validator"
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
		Firstname:  request.Firstname,
		Lastname:   request.Lastname,
		Email:      request.Email,
		Phone:      request.Phone,
		Address1:   request.Address1,
		Address2:   request.Address2,
		City:       request.City,
		PostalCode: request.PostalCode,
		Password:   hashedPassword.Hash,
		Salt:       hashedPassword.Salt,
	}

	fmt.Println()
	// Save user to the database
	if err := database.DB.Create(user).Error; err != nil {
		return nil, &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: "Failed to save user"}
	}

	return user, nil
}

func VerifyData(UserData models.CreateUserRequest) *models.ServiceError {
	if UserData.Password == "" {
		return &models.ServiceError{StatusCode: http.StatusUnprocessableEntity, Message: "Password can't be empty"}
	}

	err := passwordvalidator.Validate(UserData.Password, auth.MinEntropyBits)
	if err != nil {
		return &models.ServiceError{StatusCode: http.StatusForbidden, Message: "Password is insecure"}
	}

	if UserData.Email == "" {
		return &models.ServiceError{StatusCode: http.StatusUnprocessableEntity, Message: "Email can't be empty"}
	}

	if UserData.Email == "" {
		return &models.ServiceError{StatusCode: http.StatusUnprocessableEntity, Message: "Email format is wrong"}
	}

	if UserData.Firstname == "" {
		return &models.ServiceError{StatusCode: http.StatusUnprocessableEntity, Message: "Firstname can't be empty"}
	}

	if UserData.Lastname == "" {
		return &models.ServiceError{StatusCode: http.StatusUnprocessableEntity, Message: "Lastname can't be empty"}
	}

	if UserData.Phone != nil {

		if UserData.PhoneRegion == nil {
			return &models.ServiceError{StatusCode: http.StatusUnprocessableEntity, Message: "Phone region has to be sent"}
		}

		if *UserData.Phone == "" {
			return &models.ServiceError{StatusCode: http.StatusUnprocessableEntity, Message: "Phone number can't be empty"}
		}

		if *UserData.PhoneRegion == "" {
			return &models.ServiceError{StatusCode: http.StatusUnprocessableEntity, Message: "Phone region can't be empty"}
		}

		ParsedNumber, err := phonenumbers.Parse(*UserData.Phone, *UserData.PhoneRegion)
		if err != nil {
			return &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: "Error while trying to parse phone number"}
		}

		if phonenumbers.IsValidNumber(ParsedNumber) == false {
			return &models.ServiceError{StatusCode: http.StatusUnprocessableEntity, Message: "Phone format is not valid"}
		}
	}

	if UserData.Address1 != nil {
		if *UserData.Address1 == "" {
			return &models.ServiceError{StatusCode: http.StatusUnprocessableEntity, Message: "Address1 can't be empty"}
		}
	}

	if UserData.Address2 != nil {
		if *UserData.Address2 == "" {
			return &models.ServiceError{StatusCode: http.StatusUnprocessableEntity, Message: "Address2 can't be empty"}
		}
	}

	if UserData.City != nil {
		if *UserData.City == "" {
			return &models.ServiceError{StatusCode: http.StatusUnprocessableEntity, Message: "City can't be empty"}
		}
	}

	if UserData.PostalCode != nil {
		if *UserData.PostalCode == "" {
			return &models.ServiceError{StatusCode: http.StatusUnprocessableEntity, Message: "Postal code can't be empty"}
		}
	}

	return nil
}
