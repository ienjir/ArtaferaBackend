package order

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"net/http"
	"strconv"
)

func CreateOrder(c *gin.Context) {
	var json models.CreateOrderRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	json.AuthID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	// Verify the order creation request
	if err := verifyCreateOrder(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	// Create the order
	order, err := createOrderService(json)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

func GetOrderByID(c *gin.Context) {
	var json models.GetOrderByIDRequest
	var order *models.Order
	var err *models.ServiceError

	orderID, err2 := strconv.ParseInt(c.Param("id"), 10, 64)
	if err2 != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "OrderID is wrong"})
	}

	userID := c.GetInt64("userID")
	userRole := c.GetString("userRole")

	json.OrderID = orderID
	json.UserID = userID
	json.UserRole = userRole

	if err = verifyGetOrderByID(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	if order, err = getOrderByIDService(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
	return
}

func GetOrdersForUser(c *gin.Context) {
	var json models.GetOrdersForUserRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	targetID, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusInternalServerError, "Could not convert ID")
		return
	}

	json.TargetUserID = targetID
	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := verifyGetOrdersForUser(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	orders, user, count, err := getOrdersForUserService(json)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
	}

	c.JSON(http.StatusOK, gin.H{"count": count, "user": user, "orders": orders})
	return
}

func ListOrder(c *gin.Context) {
	var json models.ListOrdersRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := verifyListOrders(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	orders, count, err := listOrderService(json)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"count": count, "orders": orders})
	return
}
