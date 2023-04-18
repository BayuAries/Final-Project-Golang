package router

import (
	"MyGram/controllers"
	"MyGram/middleware"

	"github.com/gin-gonic/gin"
)

func StarApp() *gin.Engine {
	r := gin.Default()

	// User Route
	userRoute := r.Group("/users")
	{
		// Register [POST] user
		userRoute.POST("/register", controllers.UserRegister)
		// Login [POST] user
		userRoute.POST("/login", controllers.UserLogin)
	}

	// Photo Route
	photoRoute := r.Group("/photo")
	{
		// Authentication for Photo Route
		photoRoute.Use(middleware.Authentication())
		// Create Photo [POST]
		photoRoute.POST("/", controllers.CreatePhoto)
		// Update Photo [PUT]
		photoRoute.PUT("/:photoId", middleware.PhotoAuthorization(), controllers.UpdatePhoto)
		// GetOne Photo [GET]
		photoRoute.GET("/:photoId", controllers.GetOnePhoto)
		// GetAll Photo [GET]
		photoRoute.GET("/all", controllers.GetAllPhoto)
		// Delete Photo [DELETE]
		photoRoute.DELETE("/:photoId", middleware.PhotoAuthorization(), controllers.DeletePhoto)
	}

	// Comment Route
	commentRoute := r.Group("/comment")
	{
		// Authentication for comment Route
		commentRoute.Use(middleware.Authentication())
		// Create Comment [POST]
		commentRoute.POST("/", controllers.CreateComment)
		// Update Comment [PUT]
		commentRoute.PUT("/:commentId", middleware.CommentAuthorization(), controllers.UpdateComment)
		// Get One Comment [GET]
		commentRoute.GET("/:commentId", controllers.GetOneComment)
		// Get All Comment [GET]
		commentRoute.GET("/all/:photoId", controllers.GetAllComment)
		// Delete Comment [DELETE]
		commentRoute.DELETE("/:commentId", middleware.CommentAuthorization(), controllers.DeleteComment)
	}

	// Social Media Route
	socialMediaRoute := r.Group("/socialmedia")
	{
		// authentication for social media
		socialMediaRoute.Use(middleware.Authentication())
		// Create Social Media [POST]
		socialMediaRoute.POST("/", controllers.CreateSocialMedia)
		// Update Social Media [PUT]
		socialMediaRoute.PUT("/:sosmedId", middleware.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		// Get One Social Media [GET]
		socialMediaRoute.GET("/:sosmedId", controllers.GetOneSocialMedia)
		// Get All Social Media [GET]
		socialMediaRoute.GET("/all", controllers.GetAllSocialMedia)
		// Delete Social Media [DELETE]
		socialMediaRoute.DELETE("/:sosmedId", middleware.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
	}

	return r
}
