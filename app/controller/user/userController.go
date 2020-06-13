package user

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	userModel "github.com/rudiarta/kriya/app/model/user"
	"github.com/rudiarta/kriya/config/database"
)

func Routes(route *gin.Engine) {
	item := new(userModel.User)
	item.Data = userModel.UserData{
		"name": "test",
	}
	item.RoleID = "381b7700-fd23-44b7-9d1f-befba9fa7d6a"

	user := route.Group("/user")
	{
		user.GET("/test", func(c *gin.Context) {
			db, _ := database.InitDatabase()
			db.Create(&item)
			defer db.Close()
			c.JSON(200, gin.H{
				"message": item,
			})
		})
	}
}
