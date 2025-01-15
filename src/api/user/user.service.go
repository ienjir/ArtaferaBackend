package user

import (
	"errors"
	"fmt"
	"github.com/ienjir/ArtaferaBackend/src/api/auth"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
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
