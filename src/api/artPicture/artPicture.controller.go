package artPicture

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

const UploadDir = "uploads/"

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "File upload failed")
		return
	}

	// Ensure the upload directory exists
	if err := os.MkdirAll(UploadDir, os.ModePerm); err != nil {
		c.String(http.StatusInternalServerError, "Could not create upload directory")
		return
	}

	dst := UploadDir + file.Filename
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.String(http.StatusInternalServerError, "Failed to save file")
		return
	}

	log.Println("File uploaded:", file.Filename)
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
