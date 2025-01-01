package files

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func UploadFile(c *gin.Context) {
	// Single file
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("Error retrieving the file:", err)
		c.String(http.StatusBadRequest, "No file is uploaded")
		return
	}

	log.Println("File Name:", file.Filename)

	// Ensure the upload directory exists
	uploadDir := "../uploads"
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		err = os.MkdirAll(uploadDir, os.ModePerm)
		if err != nil {
			log.Println("Error creating directory:", err)
			c.String(http.StatusInternalServerError, "Failed to create upload directory")
			return
		}
	}

	// Save the uploaded file
	filePath := fmt.Sprintf("%s/%s", uploadDir, file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		log.Println("Error saving the file:", err)
		c.String(http.StatusInternalServerError, "Failed to save the uploaded file")
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded successfully!", file.Filename))
}

func UploadSingleFile(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("image")
	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}

	fileExt := filepath.Ext(header.Filename)
	originalFileName := strings.TrimSuffix(filepath.Base(header.Filename), filepath.Ext(header.Filename))
	now := time.Now()
	filename := strings.ReplaceAll(strings.ToLower(originalFileName), " ", "-") + "-" + fmt.Sprintf("%v", now.Unix()) + fileExt
	filePath := "http://localhost:8000/images/single/" + filename

	out, err := os.Create("public/single/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"filename": filename, "filepath": filePath})
}
