package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/database"
	"github.com/ienjir/ArtaferaBackend/models"
	"net/http"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}
