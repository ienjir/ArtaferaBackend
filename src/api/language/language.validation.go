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
