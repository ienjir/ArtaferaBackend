package role

import (
	"github.com/ienjir/ArtaferaBackend/src/models"
	"net/http"
)

func verifyGetRoleByID(requestUserID int64, requestUserRole, targetUserID string) *models.ServiceError {
	if requestUserRole != "admin" {
		return &models.ServiceError{StatusCode: http.StatusUnauthorized, Message: "You are not allowed to see this role"}
	}

	return nil

}
