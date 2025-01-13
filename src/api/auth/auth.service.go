package auth

import (
	"errors"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"gorm.io/gorm"
)

func VerifyUserExists(email, encryptedPassword string) (*models.User, error) {
	var err error
	var user models.User

	if err = database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found or incorrect credentials")
		}
		return nil, err
	}

	if user.Password != encryptedPassword {
		return nil, errors.New("incorrect password")
	}

	return &user, nil
}
