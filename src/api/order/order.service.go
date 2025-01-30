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

	if art.Available == false {
		return nil, &models.ServiceError{
			StatusCode: http.StatusConflict,
			Message:    "Art is not available",
		}
	}

	order = models.Order{
		OrderDate: time.Now(),
		UserID:    int64(int(*data.UserID)),
		ArtID:     int64(int(art.ID)),
		Status:    models.OrderStatusPending,
	}

	// Start transaction
	tx := database.DB.Begin()
	if tx.Error != nil {
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to start transaction",
		}
	}

	art.Available = false
	if err := tx.Save(&art).Error; err != nil {
		tx.Rollback()
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to update corresponding art",
		}
	}

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

func getOrderByIDService(data models.GetOrderByIDRequest) (*models.Order, *models.ServiceError) {
	var order models.Order

	if err := database.DB.Preload("User").Preload("Art").First(&order, data.OrderID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &models.ServiceError{
				StatusCode: http.StatusNotFound,
				Message:    "Order not found",
			}
		}
		return nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while retrieving order",
		}
	}

	if data.UserRole != "admin" {
		if data.UserID != order.UserID {
			return nil, &models.ServiceError{
				StatusCode: http.StatusForbidden,
				Message:    "You can only see your own orders",
			}
		}
	}

	return &order, nil
}

func getOrdersForUserService(data models.GetOrdersForUserRequest) (*[]models.Order, *models.User, *int64, *models.ServiceError) {
	var orders []models.Order
	var user models.User
	var count int64

	if err := database.DB.First(&user, data.TargetUserID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, nil, &models.ServiceError{
				StatusCode: http.StatusNotFound,
				Message:    "User not found",
			}
		} else {
			return nil, nil, nil, &models.ServiceError{StatusCode: http.StatusInternalServerError, Message: "Error while retrieving user"}
		}
	}

	if err := database.DB.Preload("Art").Where("user_id = ?", data.TargetUserID).Find(&orders).Limit(5).Offset(int(data.Offset) * 5).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, nil, &models.ServiceError{
				StatusCode: http.StatusNotFound,
				Message:    "Orders not found",
			}
		}
		return nil, nil, nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while retrieving orders",
		}
	}

	if err := database.DB.Model(&models.Order{}).Where("user_id = ?", data.TargetUserID).Count(&count).Error; err != nil {
		return nil, nil, nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while counting orders in database",
		}
	}

	return &orders, &user, &count, nil
}

func listOrderService(data models.ListOrdersRequest) (*[]models.Order, *int64, *models.ServiceError) {
	var orders []models.Order
	var count int64

	if err := database.DB.Preload("Art").Preload("User").Limit(5).Offset(int(data.Offset * 5)).Find(&orders).Error; err != nil {

		return nil, nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while retrieving orders from database",
		}
	}

	if err := database.DB.Model(&models.Order{}).Count(&count).Error; err != nil {
		return nil, nil, &models.ServiceError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while counting orders in database",
		}
	}

	return &orders, &count, nil
}
