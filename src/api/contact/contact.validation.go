package contact

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/utils"
	"github.com/ienjir/ArtaferaBackend/src/validation"
)

func verifyCreateContactMessage(data models.CreateContactMessageRequest) *models.ServiceError {
	if err := validation.ValidateEmail(data.Email); err != nil {
		return err
	}

	if data.Message == "" {
		return utils.NewFieldEmptyError("Message")
	}

	return nil
}
