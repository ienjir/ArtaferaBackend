package user

import (
	"errors"
	"github.com/ienjir/ArtaferaBackend/src/api/auth"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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

func GetUserByEmailService(Data models.GetUserByEmail) (*models.User, *models.ServiceError) {
	var user models.User

	if err := database.DB.Preload("Role").Where("email = ?", Data.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &models.ServiceError{StatusCode: http.StatusNotFound, Message: "User not found"}
		}
		return nil, &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}

	if Data.RequestRole != "admin" && int(Data.RequestID) != int(user.ID) {
		return nil, &models.ServiceError{StatusCode: http.StatusUnauthorized, Message: "You can only see your own account"}
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
	parsedUserID, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: "Invalid user ID"}
	}

	if result := database.DB.Delete(&models.User{}, parsedUserID); result.Error != nil {
		return &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: "Error occurred while deleting user"}
	}

	return nil
}

func UpdateUserService(requestUserID int64, requestUserRole string, targetUserID string, req models.UpdateUserRequest) *models.ServiceError {
	targetUserIDInt64, err := strconv.ParseInt(targetUserID, 10, 64)
	if err != nil {
		return &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: "Could not parse numbers"}
	}

	if requestUserRole != "admin" && requestUserID != targetUserIDInt64 {
		return &models.ServiceError{StatusCode: http.StatusForbidden, Message: "You can only update your account"}
	}

	// Find the target user
	var user models.User
	if err := database.DB.First(&user, "id = ?", targetUserID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &models.ServiceError{StatusCode: http.StatusNotFound, Message: "User not found"}
		}
		return &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}

	// Update fields that are provided
	if req.Firstname != nil {
		user.Firstname = *req.Firstname
	}
	if req.Lastname != nil {
		user.Lastname = *req.Lastname
	}
	if req.Email != nil {
		user.Email = *req.Email
	}
	if req.Phone != nil {
		user.Phone = req.Phone
	}
	if req.PhoneRegion != nil {
		user.PhoneRegion = req.PhoneRegion
	}
	if req.Address1 != nil {
		user.Address1 = req.Address1
	}
	if req.Address2 != nil {
		user.Address2 = req.Address2
	}
	if req.City != nil {
		user.City = req.City
	}
	if req.PostalCode != nil {
		user.PostalCode = req.PostalCode
	}

	if req.RoleID != nil {
		if requestUserRole != "admin" {
			return &models.ServiceError{StatusCode: http.StatusForbidden, Message: "You are not authorized to change the role of a user"}
		}
		user.RoleID = *req.RoleID
	}

	// Handle password update with hashing
	if req.Password != nil {
		password, err := auth.HashPassword(*req.Password)
		if err != nil {
			return &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: err.Error()}
		}
		user.Password = password.Hash
		user.Salt = password.Salt
	}

	if err = database.DB.Save(&user).Error; err != nil {
		return &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}

	return nil
}
