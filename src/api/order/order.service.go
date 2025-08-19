package order

import (
	"errors"
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/utils"
	"github.com/ienjir/ArtaferaBackend/src/validation"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func getOrderByIDService(data models.GetOrderByIDRequest) (*models.Order, *models.ServiceError) {
	var order models.Order

	if err := database.DB.Preload("User").Preload("Art").First(&order, data.OrderID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewOrderNotFoundError()
		}
		return nil, utils.NewDatabaseRetrievalError()
	}

	if data.UserRole != "admin" {
		if data.UserID != order.UserID {
			return nil, utils.NewOwnerOnlyOrdersError()
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
			return nil, nil, nil, utils.NewUserNotFoundError()
		} else {
			return nil, nil, nil, utils.NewDatabaseRetrievalError()
		}
	}

	if err := database.DB.Preload("Art").Where("user_id = ?", data.TargetUserID).Find(&orders).Limit(5).Offset(int(data.Offset) * 5).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, nil, utils.NewOrderNotFoundError()
		}
		return nil, nil, nil, utils.NewDatabaseRetrievalError()
	}

	if err := database.DB.Model(&models.Order{}).Where("user_id = ?", data.TargetUserID).Count(&count).Error; err != nil {
		return nil, nil, nil, utils.NewDatabaseCountError()
	}

	return &orders, &user, &count, nil
}

func listOrderService(data models.ListOrdersRequest) (*[]models.Order, *int64, *models.ServiceError) {
	var orders []models.Order
	var count int64

	if err := database.DB.Preload("Art").Preload("User").Limit(5).Offset(int(data.Offset * 5)).Find(&orders).Error; err != nil {
		return nil, nil, utils.NewDatabaseRetrievalError()
	}

	if err := database.DB.Model(&models.Order{}).Count(&count).Error; err != nil {
		return nil, nil, utils.NewDatabaseCountError()
	}

	return &orders, &count, nil
}

func createOrderService(data models.CreateOrderRequest) (*models.Order, *models.ServiceError) {
	var art models.Art
	var order models.Order

	if err := database.DB.First(&art, data.ArtID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewArtNotFoundError()
		}
		return nil, utils.NewDatabaseRetrievalError()
	}

	if art.Available == false {
		return nil, utils.NewArtNotAvailableError()
	}

	order = models.Order{
		OrderDate: time.Now(),
		UserID:    int64(int(*data.TargetUserID)),
		ArtID:     int64(int(art.ID)),
		Status:    models.OrderStatusPending,
	}

	// Start transaction
	tx := database.DB.Begin()
	if tx.Error != nil {
		return nil, utils.NewTransactionStartError()
	}

	art.Available = false
	if err := tx.Save(&art).Error; err != nil {
		tx.Rollback()
		return nil, utils.NewDatabaseUpdateError()
	}

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return nil, utils.NewDatabaseCreateError()
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return nil, utils.NewTransactionCommitError()
	}

	return &order, nil
}

func updateOrderService(data models.UpdateOrderRequest) (*models.Order, *models.ServiceError) {
	var order models.Order

	if err := database.DB.First(&order, "id = ?", data.TargetID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.NewOrderNotFoundError()
		}
		return nil, utils.NewDatabaseRetrievalError()
	}

	if data.TargetUserID != nil {
		order.UserID = *data.TargetUserID
	}

	if data.ArtID != nil {
		order.ArtID = *data.ArtID
	}

	if data.Status != nil {
		// Already gets validated in the validation file
		status, _ := validation.ValidateStatusString(*data.Status)
		order.Status = status
	}

	if err := database.DB.Save(&order).Error; err != nil {
		return nil, utils.NewDatabaseUpdateError()
	}

	return &order, nil
}

func deleteOrderService(data models.DeleteOrderRequest) *models.ServiceError {
	var order models.Order

	if err := database.DB.First(&order, "id = ?", data.TargetID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NewOrderNotFoundError()
		}
		return utils.NewDatabaseRetrievalError()
	}

	if result := database.DB.Delete(&models.Order{}, data.TargetID); result.Error != nil {
		return utils.NewDatabaseDeleteError()
	}

	return nil
}
