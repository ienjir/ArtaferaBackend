package contact

import (
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
)

func createContactMessageService(data models.CreateContactMessageRequest) (*models.ContactMessage, *models.ServiceError) {
	message := models.ContactMessage{
		Email:   data.Email,
		Message: data.Message,
	}

	if err := database.Repositories.ContactMessage.Create(&message); err != nil {
		return nil, err
	}

	notifyContactMessageReceived(&message)

	return &message, nil
}

func notifyContactMessageReceived(message *models.ContactMessage) {
	// TODO: hook up notifications (email, SMS, webhook, etc.).
}
