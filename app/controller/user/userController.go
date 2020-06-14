package user

import (
	"github.com/gin-gonic/gin"
	"github.com/rudiarta/kriya/app/middleware"
	"github.com/rudiarta/kriya/app/model/role"
	userModel "github.com/rudiarta/kriya/app/model/user"
	"github.com/rudiarta/kriya/app/service"
	"github.com/rudiarta/kriya/config/database"
)

func Routes(route *gin.Engine) {
	user := route.Group("/user")
	{
		user.POST("/addUser", middleware.AdminMiddleware(), func(c *gin.Context) {
			db, _ := database.InitDatabase()
			defer db.Close()

			password, _ := service.HashPassword(c.PostForm("password"))
			item := userModel.User{
				Data: userModel.UserData{
					Email:    c.PostForm("email"),
					Username: c.PostForm("username"),
					Password: password,
					Status: userModel.StatusData{
						IsActive: true,
					},
				},
				RoleID: "d57bfbfe-4979-4809-a151-f6cd30de657b",
			}

			var roleData role.Role
			db.Where("id = ?", item.RoleID).First(&roleData)

			response := userModel.UserResponse{
				Role: roleData.Data.RoleName,
				Data: item.Data,
			}

			if err := db.Create(&item).Error; err != nil {
				c.JSON(422, gin.H{
					"message": "fail",
					"data":    response,
				})

				return
			}

			c.JSON(200, gin.H{
				"message": "success",
				"data":    response,
			})

			return
		})
	}
}
