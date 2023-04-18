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

// CreateComment function
func CreateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}
	photoId := &Comment.PhotoID

	// Check Photo with id exist
	err := db.Select("title").First(&Photo, "id = ?", photoId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Data Not Found",
			"message": err.Error(),
		})
		return
	}

	Comment.UserID = userID

	err = db.Debug().Create(&Comment).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Comment create successfully!",
		"data":    Comment,
	})
}

// UpdateComment function
func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	commentId, _ := strconv.Atoi(c.Param("commentId"))

	Comment := models.Comment{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userID
	Comment.ID = uint(commentId)

	// Check if Data Exist
	err := db.Select("id").First(&Comment, "id = ? ", commentId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Data Not Found",
			"message": err.Error(),
		})
		return
	}

	// err := db.Model(&Comment).Where("id = ? AND photo_id = ?", commentId, photoId).Updates(models.Comment{Message: Comment.Message}).Error
	err = db.Model(&Comment).Where("id = ? ", commentId).Updates(models.Comment{Message: Comment.Message}).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Update Comment successfully!",
		"data":    Comment,
	})
}

// GetOneComment function
func GetOneComment(c *gin.Context) {
	db := database.GetDB()
	commentId, _ := strconv.Atoi(c.Param("commentId"))

	Comment := models.Comment{}

	err := db.Model(&Comment).Preload("Photo").Preload("User").First(&Comment, "id = ? ", commentId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Get Comment successfully!",
		"data":    Comment,
	})
}

// GetAllComment function
func GetAllComment(c *gin.Context) {
	db := database.GetDB()
	photoId, _ := strconv.Atoi(c.Param("photoId"))

	Comment := []models.Comment{}

	// Check if Data Exist
	err := db.Select("message").First(&Comment, "photo_id = ?", photoId).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Data Not Found",
			"message": err.Error(),
		})
		return
	}

	err = db.Preload("Photo").Preload("User").Find(&Comment, "photo_id = ?", photoId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Get Comments successfully!",
		"data":    Comment,
	})
}

// DeleteComment function
func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	commentId, _ := strconv.Atoi(c.Param("commentId"))

	Comment := models.Comment{}

	err := db.Model(&Comment).Delete(&Comment, "id = ?", commentId).Error
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
