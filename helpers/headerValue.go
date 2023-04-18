package helpers

import "github.com/gin-gonic/gin"

// Get content type for api
func GetContentType(c *gin.Context) string {
	return c.Request.Header.Get("Content-Type")
}
