package controllers

import (
	"MyGram/database"
	"MyGram/helpers"
	"MyGram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Create Social Media Function
func CreateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	SocialMedia := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	// Check if there is social media from user id exist
	err := db.Select("social_media_url").First(&SocialMedia, "user_id = ?", userID).Error
	if err == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Data already exist can't add another social media",
		})
		return
	}

	SocialMedia.UserID = userID
	err = db.Debug().Create(&SocialMedia).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Social Media Create successfully!",
		"data":    SocialMedia,
	})
}

// Update Social Media Function
func UpdateSocialMedia(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	sosmedId, _ := strconv.Atoi(c.Param("sosmedId"))

	SocialMedia := models.SocialMedia{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&SocialMedia)
	} else {
		c.ShouldBind(&SocialMedia)
	}

	SocialMedia.ID = uint(sosmedId)
	SocialMedia.UserID = userID

	err := db.Model(&SocialMedia).Where("id = ?", sosmedId).Updates(models.SocialMedia{Name: SocialMedia.Name, SocialMediaUrl: SocialMedia.SocialMediaUrl}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Update Social Media successfully!",
		"data":    SocialMedia,
	})
}

// Get One Social Media Function
func GetOneSocialMedia(c *gin.Context) {
	db := database.GetDB()
	sosmedId, _ := strconv.Atoi(c.Param("sosmedId"))

	SocialMedia := models.SocialMedia{}
	SocialMedia.ID = uint(sosmedId)

	// Get sosmed by user id
	// userData := c.MustGet("userData").(jwt.MapClaims)
	// userID := uint(userData["id"].(float64))
	// err := db.Model(&SocialMedia).First(&SocialMedia, "user_id = ?", userID).Error

	// Get sosmed by id
	err := db.Model(&SocialMedia).Preload("User").First(&SocialMedia, "id = ?", sosmedId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Get Social Media successfully!",
		"data":    SocialMedia,
	})
}

// Get All Social Media Function
func GetAllSocialMedia(c *gin.Context) {
	db := database.GetDB()

	SocialMedia := []models.SocialMedia{}

	err := db.Model(&SocialMedia).Preload("User").Find(&SocialMedia).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Get Social Media successfully!",
		"data":    SocialMedia,
	})
}

// Delete Social Media Function
func DeleteSocialMedia(c *gin.Context) {
	db := database.GetDB()
	sosmedId, _ := strconv.Atoi(c.Param("sosmedId"))

	SocialMedia := models.SocialMedia{}
	SocialMedia.ID = uint(sosmedId)

	err := db.Model(&SocialMedia).Delete(&SocialMedia, "id = ?", sosmedId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete Succesfully",
	})
}
