package picture

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/validation"
	"net/http"
)

func verifyCreatePicture(data models.CreatePictureRequest, c *gin.Context) *models.ServiceError {
	if data.Priority != nil {
		if *data.Priority < 1 {
			return &models.ServiceError{
				StatusCode: http.StatusBadRequest,
				Message:    "Priority must be greater than 0",
			}
		}
	}

	/*
		if data.UserRole != "admin" {
			return &models.ServiceError{
				StatusCode: http.StatusForbidden,
				Message:    "Only admins can upload pictures",
			}
		}
	*/

	file, err := c.FormFile("image")
	if err != nil {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "No image file uploaded",
		}
	}

	if !validation.IsValidImage(file) {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid image format. Allowed formats: jpg, jpeg, png, gif",
		}
	}

	return nil
}
