package user

import (
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	user := route.Group("/user")
	{
		user.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "ping",
			})
		})
	}
}
