package user

import (
	"errors"
	"github.com/ienjir/ArtaferaBackend/src/api/auth"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"gorm.io/gorm"
	"net/http"
)

func GetUserByIDService(data models.GetUserByIDRequest) (*models.User, *models.ServiceError) {
	var user models.User

	if err := database.DB.Preload("Role").First(&user, data.TargetID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &models.ServiceError{
				StatusCode: http.StatusNotFound,
				Message:    "User not found",
			}
		} else {
			return nil, &models.ServiceError{
				StatusCode: http.StatusInternalServerError,
				Message:    "Error while retrieving user",
			}
		}
	}

	return &user, nil
}

func GetUserByEmailService(Data models.GetUserByEmailRequest) (*models.User, *models.ServiceError) {
	var user models.User

	if err := database.DB.Preload("Role").Where("email = ?", Data.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &models.ServiceError{StatusCode: http.StatusNotFound, Message: "User not found"}
		}
		return nil, &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}

	if Data.UserRole != "admin" {
		if int(Data.UserID) != int(user.ID) {
			return nil, &models.ServiceError{
				StatusCode: http.StatusUnauthorized,
				Message:    "You can only see your own account",
			}
		}
	}

	return &user, nil
}

func ListUsersService(data models.ListUserRequest) (*[]models.User, *int64, *models.ServiceError) {
	var users []models.User
	var count int64

	if err := database.DB.Preload("Role").Limit(10).Offset(int(data.Offset * 10)).Find(&users).Error; err != nil {
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

func CreateUserService(data models.CreateUserRequest) (*models.User, *models.ServiceError) {
	var user models.User
	var newUser models.User

	if err := database.DB.Where("email = ?", data.Email).First(&user).Error; err == nil {
		return nil, &models.ServiceError{
			StatusCode: http.StatusConflict,
			Message:    "Email already in use",
		}
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Database error",
		}
	}

	hashedPassword, err := auth.HashPassword(data.Password)
	if err != nil {
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to hash password",
		}
	}

	newUser = models.User{
		Firstname:   data.Firstname,
		Lastname:    data.Lastname,
		Email:       data.Email,
		Phone:       data.Phone,
		PhoneRegion: data.PhoneRegion,
		Address1:    data.Address1,
		Address2:    data.Address2,
		City:        data.City,
		PostalCode:  data.PostalCode,
		Password:    hashedPassword.Hash,
		Salt:        hashedPassword.Salt,
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to save user",
		}
	}

	return &newUser, nil
}

func UpdateUserService(data models.UpdateUserRequest) *models.ServiceError {
	var user models.User

	if err := database.DB.First(&user, "id = ?", data.TargetID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &models.ServiceError{StatusCode: http.StatusNotFound, Message: "User not found"}
		}
		return &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}

	if data.Firstname != nil {
		user.Firstname = *data.Firstname
	}
	if data.Lastname != nil {
		user.Lastname = *data.Lastname
	}
	if data.Email != nil {
		user.Email = *data.Email
	}
	if data.Phone != nil {
		user.Phone = data.Phone
	}
	if data.PhoneRegion != nil {
		user.PhoneRegion = data.PhoneRegion
	}
	if data.Address1 != nil {
		user.Address1 = data.Address1
	}
	if data.Address2 != nil {
		user.Address2 = data.Address2
	}
	if data.City != nil {
		user.City = data.City
	}
	if data.PostalCode != nil {
		user.PostalCode = data.PostalCode
	}

	if data.Password != nil {
		password, err := auth.HashPassword(*data.Password)
		if err != nil {
			return &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: err.Error()}
		}
		user.Password = password.Hash
		user.Salt = password.Salt
	}

	if err := database.DB.Save(&user).Error; err != nil {
		return &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: "Failed to update user"}
	}

	return nil
}

func DeleteUserService(data models.DeleteUserRequest) *models.ServiceError {
	var user models.User

	if err := database.DB.First(&user, "id = ?", data.TargetID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &models.ServiceError{
				StatusCode: http.StatusNotFound,
				Message:    "User not found",
			}
		}
		return &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	if result := database.DB.Delete(&models.Role{}, data.TargetID); result.Error != nil {
		return &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error occurred while deleting user",
		}
	}

	return nil
}
