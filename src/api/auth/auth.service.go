package auth

import (
	"errors"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"gorm.io/gorm"
	"net/http"
)

var Argon2IDHash *Argon2idHash

func HashPassword(password string) (*HashSalt, error) {
	bytePassword := []byte(password)

	hashSalt, err := Argon2IDHash.GenerateHash(bytePassword, nil)
	if err != nil {
		return nil, err
	}

	return hashSalt, nil
}

// ServiceError defines a custom error with an HTTP status code
type ServiceError struct {
	StatusCode int
	Message    string
}

func CreateUserService(request models.CreateUserRequest) (*models.User, *ServiceError) {
	var existingUser models.User

	// Check if email already exists
	if err := database.DB.Where("email = ?", request.Email).First(&existingUser).Error; err == nil {
		return nil, &ServiceError{StatusCode: http.StatusConflict, Message: "Email already in use"}
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, &ServiceError{StatusCode: http.StatusInternalServerError, Message: "Database error"}
	}

	// Hash the password
	hashedPassword, err := HashPassword(request.Password)
	if err != nil {
		return nil, &ServiceError{StatusCode: http.StatusInternalServerError, Message: "Failed to hash password"}
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

	// Save user to the database
	if err := database.DB.Create(user).Error; err != nil {
		return nil, &ServiceError{StatusCode: http.StatusInternalServerError, Message: "Failed to save user"}
	}

	return user, nil
}
