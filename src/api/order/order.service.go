package order

import (
	"github.com/ienjir/ArtaferaBackend/src/database"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/utils"
	"github.com/ienjir/ArtaferaBackend/src/validation"
	"time"
)

func getOrderByIDService(data models.GetOrderByIDRequest) (*models.Order, *models.ServiceError) {
	order, err := database.Repositories.Order.GetByID(data.OrderID, "User", "Art")
	if err != nil {
		if err.StatusCode == 404 {
			return nil, utils.NewOrderNotFoundError()
		}
		return nil, err
	}

	if data.UserRole != "admin" {
		if data.UserID != order.UserID {
			return nil, utils.NewOwnerOnlyOrdersError()
		}
	}

	return order, nil
}

func getOrdersForUserService(data models.GetOrdersForUserRequest) (*[]models.Order, *models.User, *int64, *models.ServiceError) {
	user, err := database.Repositories.User.GetByID(data.TargetUserID)
	if err != nil {
		if err.StatusCode == 404 {
			return nil, nil, nil, utils.NewUserNotFoundError()
		}
		return nil, nil, nil, err
	}

	orders, err := database.Repositories.Order.FindAllByField("user_id", data.TargetUserID, int(data.Offset)*5, 5, "Art")
	if err != nil {
		return nil, nil, nil, err
	}

	// Count orders for user
	query := database.Repositories.Order.Query().Where("user_id = ?", data.TargetUserID)
	var count int64
	if countErr := query.Count(&count).Error; countErr != nil {
		return nil, nil, nil, utils.NewDatabaseCountError()
	}

	return orders, user, &count, nil
}

func listOrderService(data models.ListOrdersRequest) (*[]models.Order, *int64, *models.ServiceError) {
	orders, err := database.Repositories.Order.List(int(data.Offset*5), 5, "Art", "User")
	if err != nil {
		return nil, nil, err
	}

	count, err := database.Repositories.Order.Count()
	if err != nil {
		return nil, nil, err
	}

	return orders, count, nil
}

func createOrderService(data models.CreateOrderRequest) (*models.Order, *models.ServiceError) {
	art, err := database.Repositories.Art.GetByID(data.ArtID)
	if err != nil {
		if err.StatusCode == 404 {
			return nil, utils.NewArtNotFoundError()
		}
		return nil, err
	}

	if art.Available == false {
		return nil, utils.NewArtNotAvailableError()
	}

	order := models.Order{
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
	order, err := database.Repositories.Order.GetByID(data.TargetID)
	if err != nil {
		if err.StatusCode == 404 {
			return nil, utils.NewOrderNotFoundError()
		}
		return nil, err
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

	if err := database.Repositories.Order.Update(order); err != nil {
		return nil, err
	}

	return order, nil
}

func deleteOrderService(data models.DeleteOrderRequest) *models.ServiceError {
	if err := database.Repositories.Order.Delete(data.TargetID); err != nil {
		if err.StatusCode == 404 {
			return utils.NewOrderNotFoundError()
		}
		return err
	}

	return nil
}
