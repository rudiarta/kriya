package user

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	userModel "github.com/rudiarta/kriya/app/model/user"
)

func Routes(route *gin.Engine) {
	item := new(userModel.User)
	item.Data = userModel.UserData{
		"name": "test",
	}
	item.RoleID = 1

	user := route.Group("/user")
	{
		user.GET("/test", func(c *gin.Context) {
			db, _ := gorm.Open("postgres", "host=18.140.67.82 port=8976 user=test dbname=kriya_test password=kriyatest123")
			db.Create(&item)
			defer db.Close()
			c.JSON(200, gin.H{
				"message": item,
			})
		})
	}
}
