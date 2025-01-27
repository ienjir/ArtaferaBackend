package language

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"net/http"
)

func verifyGetLanguageByID(requestUserID int64, requestUserRole, targetLanguage string) *models.ServiceError {
	if requestUserRole != "admin" {
		return &models.ServiceError{StatusCode: http.StatusUnauthorized, Message: "You are not allowed to see this language"}
	}

	return nil
}

func verifyUpdateLanguage(request models.UpdateLanguageRequest) *models.ServiceError {
	if request.LanguageCode == "" {
		return &models.ServiceError{StatusCode: http.StatusBadRequest, Message: "Language code can't be null"}
	}

	if request.Language == "" {
		return &models.ServiceError{StatusCode: http.StatusBadRequest, Message: "Language can't be null"}
	}

	return nil
}
