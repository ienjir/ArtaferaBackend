package auth

import (
	"errors"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"gorm.io/gorm"
)

func VerifyUserExists(email string, encryptedPassword []byte) (*models.User, error) {
	var err error
	var user models.User

	if err = database.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("404: userNotFound")
		}
		return nil, err
	}

	if true {
		return nil, errors.New("401: wrongPassword")
	}

	return &user, nil
}
