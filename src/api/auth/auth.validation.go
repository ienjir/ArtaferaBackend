package auth

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/validation"
)

func VerifyLoginData(Data models.LoginRequest) *models.ServiceError {
	err := validation.ValidatePasswordWithoutEntropy(Data.Password)
	if err != nil {
		return err
	}

	err = validation.ValidateEmail(Data.Email)
	if err != nil {
		return err
	}

	return nil
}
