package auth

import (
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/utils"
)

func VerifyUser(request models.LoginRequest) (*models.User, *models.ServiceError) {
	// Check if user exists
	user, err := database.Repositories.User.FindByField("email", request.Email, "Role")
	if err != nil {
		if err.StatusCode == 404 {
			return nil, utils.NewUserNotFoundError()
		}
		return nil, err
	}

	// Compare password
	if err := ComparePassword(*user, request.Password); err != nil {
		return nil, &models.ServiceError{StatusCode: err.StatusCode, Message: err.Message}
	}

	return user, nil
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
