package auth

import (
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	auth := route.Group("/auth")
	{
		auth.GET("/login", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "ping",
			})
		})
	}
}
