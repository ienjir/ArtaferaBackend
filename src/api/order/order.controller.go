package order

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"net/http"
)

func CreateOrder(c *gin.Context) {
	var json models.CreateOrderRequest

	// First bind the JSON
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the authenticated user's ID and role from context
	userID := c.GetInt64("userID")
	json.UserID = &userID
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
