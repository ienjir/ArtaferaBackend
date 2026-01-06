package art

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ienjir/ArtaferaBackend/src/models"
	"github.com/ienjir/ArtaferaBackend/src/utils"
)

func GetArtByID(c *gin.Context) {
	var json models.GetArtByIDRequest
	artID, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidID)
		return
	}
	json.TargetID = artID

	json.LanguageCode = c.Query("lang")

	if err := verifyGetArtByID(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}
	art, err := getArtByIDService(json)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, art)
	return
}

func ListArts(c *gin.Context) {
	var json models.ListArtRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidJSON)
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := verifyListArt(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	arts, count, err := listArtService(json)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"count": count, "arts": arts})
	return
}

func ListArtForArtPage(c *gin.Context) {
	var json models.ListArtForArtPageRequest
	
	json.Offset = 0
	json.Limit = 20 	

	if err := c.ShouldBindQuery(&json); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request parameters")
		return
	}
	
	if err := verifyListArtForArtPage(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}
	
	arts, count, err := listArtForArtPageService(json)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}
	
	utils.RespondWithSuccess(c, http.StatusOK, gin.H{
		"arts":   arts,
		"count":  count,
		"offset": json.Offset,
		"limit":  json.Limit,
	})
}

func ListFeaturedArt(c *gin.Context) {
	var json models.ListFeaturedArtRequest
	json.Limit = 3

	if err := c.ShouldBindQuery(&json); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request parameters")
		return
	}

	if err := verifyListFeaturedArt(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	arts, count, err := listFeaturedArtService(json)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{
		"arts":  arts,
		"count": count,
		"limit": json.Limit,
	})
}

func CreateArt(c *gin.Context) {
	var json models.CreateArtRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidJSON)
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")

	if err := verifyCreateArt(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	art, err := createArtService(json)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusCreated, gin.H{"art": art})
	return
}

func UpdateArt(c *gin.Context) {
	var json models.UpdateArtRequest

	if err := c.ShouldBindJSON(&json); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidJSON)
		return
	}

	artID, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidID)
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")
	json.TargetID = artID

	if err := verifyUpdateArt(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	art, err := updateArtService(json)
	if err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"art": art})
	return
}

func DeleteArt(c *gin.Context) {
	var json models.DeleteArtRequest

	artID, parseErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if parseErr != nil {
		utils.RespondWithError(c, http.StatusBadRequest, utils.ErrInvalidID)
		return
	}

	json.UserID = c.GetInt64("userID")
	json.UserRole = c.GetString("userRole")
	json.TargetID = artID

	if err := verifyDeleteArt(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	if err := deleteArtService(json); err != nil {
		utils.RespondWithServiceError(c, err)
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, gin.H{"message": "Art successfully deleted"})
	return
}
