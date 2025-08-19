package auth

import (
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/utils"
)

func VerifyUser(request models.LoginRequest) (*models.User, *models.ServiceError) {
	var User models.User

	// Check if user exists
	if err := database.DB.Preload("Role").Where("email = ?", request.Email).First(&User).Error; err != nil {
		return nil, utils.NewUserNotFoundError()
	}

	// Compare password
	if err := ComparePassword(User, request.Password); err != nil {
		return nil, &models.ServiceError{StatusCode: err.StatusCode, Message: err.Message}
	}

	return &User, nil
}

func ComparePassword(user models.User, password string) *models.ServiceError {
	var Password, Hash, Salt []byte

	Password = []byte(password)
	Hash = user.Password
	Salt = user.Salt

	err := Argon2IDHash.Compare(Hash, Salt, Password)
	if err != nil {
		return utils.NewPasswordWrongError()
	}

	return nil
}
