package order

import (
	"errors"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func createOrderService(data models.CreateOrderRequest) (*models.Order, *models.ServiceError) {
	var art models.Art
	var order models.Order

	// Check if UserID is nil
	if data.UserID == nil {
		return nil, &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "UserID is required",
		}
	}

	// Find the art piece
	if err := database.DB.First(&art, data.ArtID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &models.ServiceError{
				StatusCode: http.StatusNotFound,
				Message:    "Art not found",
			}
		}
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while retrieving art",
		}
	}

	// Check if art is available
	if art.Available == false {
		return nil, &models.ServiceError{
			StatusCode: http.StatusConflict,
			Message:    "Art is not available",
		}
	}

	// Create the order
	order = models.Order{
		OrderDate: time.Now(),
		UserID:    int64(int(*data.UserID)), // Convert *int64 to int
		ArtID:     int64(int(art.ID)),       // Convert uint to int if needed
		Status:    models.OrderStatusPending,
	}

	// Begin transaction
	tx := database.DB.Begin()
	if tx.Error != nil {
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to start transaction",
		}
	}

	// Update art availability
	art.Available = false
	if err := tx.Save(&art).Error; err != nil {
		tx.Rollback()
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to update corresponding art",
		}
	}

	// Create the order
	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to save order",
		}
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to commit transaction",
		}
	}

	return &order, nil
}
