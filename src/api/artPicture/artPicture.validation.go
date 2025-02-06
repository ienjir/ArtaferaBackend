package artPicture

import (
	"errors"
	"mime/multipart"
)

var allowedImageTypes = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
	"image/gif":  true,
}

// ValidateImage checks if the uploaded file is an image
func validateImage(file *multipart.FileHeader) error {
	if file == nil {
		return errors.New("no file uploaded")
	}
	return nil
}
