package auth

import (
	"github.com/gin-gonic/gin"
	userModel "github.com/rudiarta/kriya/app/model/user"
	"github.com/rudiarta/kriya/app/service"
	"github.com/rudiarta/kriya/config/database"
)

func Routes(route *gin.Engine) {
	auth := route.Group("/auth")
	{
		auth.POST("/login", func(c *gin.Context) {
			db, _ := database.InitDatabase()
			defer db.Close()

			var userData []userModel.User
			email := c.PostForm("email")
			password := c.PostForm("password")

			db.Find(&userData)

			var response userModel.User
			response = userModel.User{}
			verify := false
			for _, v := range userData {
				if v.Data.Email == email && service.CheckPasswordHash(password, v.Data.Password) {
					response = v
					verify = true
				}
			}

			result, _ := service.CreateToken(response.ID)

			if verify {
				c.JSON(200, gin.H{
					"message": "success",
					"token":   result,
				})

				return
			}

			c.JSON(401, gin.H{
				"message": "Unauthorized",
			})

			return
		})

		auth.POST("/check", func(c *gin.Context) {
			token := c.PostForm("token")
			result, _ := service.VerifyToken(token)
			c.JSON(200, gin.H{
				"data": result,
			})
		})
	}
}
