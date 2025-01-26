package role

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRoleByID(c *gin.Context) {
	requestUserID := c.GetInt64("userID")
	requestUserRole := c.GetString("userRole")
	targetRoleID := c.Param("id")

	if err := verifyGetRoleByID(requestUserID, requestUserRole, targetRoleID); err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	role, err := getRoleByIDService(targetRoleID)
	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"role": role})
	return
}
