package art

import (
	"errors"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"gorm.io/gorm"
	"net/http"
)

func getArtByIDService(data models.GetArtByIDRequest) (*models.Art, *models.ServiceError) {
	var art models.Art

	if err := database.DB.First(&art, data.TargetID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &models.ServiceError{
				StatusCode: http.StatusNotFound,
				Message:    "Art not found",
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
