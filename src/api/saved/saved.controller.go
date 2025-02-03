package saved

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"net/http"
	"strconv"
)

func GetSavedByID(c *gin.Context) {
	var json models.GetSavedByIDRequest

	targetID, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not convert ID"})
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")
	json.TargetID = targetID

	fmt.Printf("userID: %d", c.GetInt64("userID"))

	if err := verifyGetSavedById(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	user, err := getSavedByIDService(json)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
	return
}

func GetSavedForUser(c *gin.Context) {
	var json models.GetSavedForUserRequest

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

	if err := verifyGetSavedForUserRequest(json); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	orders, user, count, err := getSavedForUserService(json)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
	}

	c.JSON(http.StatusOK, gin.H{"count": count, "user": user, "orders": orders})
	return
}
