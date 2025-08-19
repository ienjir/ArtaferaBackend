package order

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/utils"
	"net/http"
	"strconv"
)

func GetOrderByID(c *gin.Context) {
	var json models.GetOrderByIDRequest
	var order *models.Order
	var err *models.ServiceError

	orderID, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidID)
		return
	}

	userID := c.GetInt64("userID")
	userRole := c.GetString("userRole")

	json.OrderID = orderID
	json.UserID = userID
	json.UserRole = userRole

	if err = verifyGetOrderByIDRequest(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	if order, err = getOrderByIDService(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"order": order})
	return
}

func GetOrdersForUser(c *gin.Context) {
	var json models.GetOrdersForUserRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidJSON)
		return
	}

	targetID, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidID)
		return
	}

	json.TargetUserID = targetID
	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := verifyGetOrdersForUserRequest(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	orders, user, count, err := getOrdersForUserService(json)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"count": count, "user": user, "orders": orders})
	return
}

func ListOrder(c *gin.Context) {
	var json models.ListOrdersRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidJSON)
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := verifyListOrdersRequest(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	orders, count, err := listOrderService(json)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"count": count, "orders": orders})
	return
}

func CreateOrder(c *gin.Context) {
	var json models.CreateOrderRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidJSON)
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := verifyCreateOrder(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	order, err := createOrderService(json)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"order": order})
}

func UpdateOrder(c *gin.Context) {
	var json models.UpdateOrderRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidJSON)
		return
	}

	targetID, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidID)
		return
	}

	json.TargetID = targetID
	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := verifyUpdateOrderRequest(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	order, err := updateOrderService(json)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"order": order})
	return
}

func DeleteOrder(c *gin.Context) {
	var json models.DeleteOrderRequest

	targetID, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidID)
		return
	}

	json.TargetID = targetID
	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := verifyDeleteOrderRequest(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	if err := deleteOrderService(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"message": "Order successfully deleted"})
}
