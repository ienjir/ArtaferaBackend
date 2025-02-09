package artPicture

import (
	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"mime/multipart"
	"net/http"
	"strings"
)

func verifyCreateArtPicture(data models.CreateArtPictureRequest, c *gin.Context) *models.ServiceError {
	if data.ArtID < 1 {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "ArtID must be greater than 0",
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

	if !isValidImage(file) {
		return &models.ServiceError{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid image format. Allowed formats: jpg, jpeg, png, gif",
		}
	}

	return nil
}

// isValidImage checks if the file is a valid image format.
func isValidImage(file *multipart.FileHeader) bool {
	allowedExtensions := []string{".jpg", ".jpeg", ".png", ".gif"}
	ext := strings.ToLower(file.Filename[strings.LastIndex(file.Filename, "."):])
	for _, validExt := range allowedExtensions {
		if ext == validExt {
			return true
		}
	}
	return false
}
