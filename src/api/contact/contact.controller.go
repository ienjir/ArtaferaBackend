package contact

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/utils"
)

func CreateContactMessage(c *gin.Context) {
	var json models.CreateContactMessageRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidJSON)
		return
	}

	if err := verifyCreateContactMessage(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	message, err := createContactMessageService(json)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusCreated, gin.H{"contactMessage": message})
}
