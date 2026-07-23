package contact

import (
	"log"

	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/utils"
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
	if err := utils.SendContactNotification(message.Email, message.Message); err != nil {
		log.Printf("Failed to send contact notification email: %v", err)
	}
}
