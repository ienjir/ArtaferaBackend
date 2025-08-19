package auth

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/validation"
)

func VerifyLoginData(Data models.LoginRequest) *models.ServiceError {
	if err := validation.ValidateEmail(Data.Email); err != nil {
		return err
	}
	return validation.ValidatePasswordWithoutEntropy(Data.Password)
}
