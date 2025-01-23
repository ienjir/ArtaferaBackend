package user

import (
	"errors"
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

func GetUserByEmailService(email string) (*models.User, *models.ServiceError) {
	var user models.User

	if err := database.DB.Where("email = ?", email).First(&user); err != nil {
		if errors.Is(err.Error, gorm.ErrRecordNotFound) {
			return nil, &models.ServiceError{StatusCode: http.StatusNotFound, Message: "User not found"}
		} else {
			return nil, &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: "Error while retrieving user"}
		}

	}

	return &user, nil
}
func GetUserByIDService(userID string) (*models.User, *models.ServiceError) {
	var user models.User

	if err := database.DB.Preload("Role").First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &models.ServiceError{StatusCode: http.StatusNotFound, Message: "User not found"}
		} else {
			return nil, &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: "Error while retrieving user"}
		}
	}

	return &user, nil
}

func ListUsersService(offset int) (*[]models.User, *int64, *models.ServiceError) {
	var users []models.User
	var count int64

	if err := database.DB.Preload("Role").Limit(5).Offset(offset * 10).Find(&users).Error; err != nil {
		return nil, nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while retrieving users from database",
		}
	}

	if err := database.DB.Model(&models.User{}).Count(&count).Error; err != nil {
		return nil, nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while counting users in database",
		}
	}

	return &users, &count, nil
}

func DeleteUserService(userID string) *models.ServiceError {
	if err := database.DB.Where("id = ?", userID).Delete(&models.User{}, userID).Error; err != nil {
		return &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: "Error occured while deleting user"}
	}

	return nil
}
