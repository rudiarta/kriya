package user

import (
	"math"
	"os"
	"strconv"

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
		user.PUT("/updateUser", middleware.AdminMiddleware(), func(c *gin.Context) {
			userID := c.PostForm("id")
			userName := c.PostForm("username")
			userEmail := c.PostForm("email")
			password, _ := service.HashPassword(c.PostForm("password"))
			userPassword := password
			db, _ := database.InitDatabase()
			defer db.Close()

			var userData userModel.User
			var roleData role.Role
			db.Where("id = ?", userID).First(&userData)
			tmpUser := userData
			tmpUser.Data.Username = userName
			tmpUser.Data.Email = userEmail
			tmpUser.Data.Password = userPassword
			db.Model(&userData).Where("id = ?", userID).Update(tmpUser)
			db.Where("id = ?", tmpUser.RoleID).First(&roleData)

			result := userModel.UserResponse{
				Role: roleData.Data.RoleName,
				Data: tmpUser.Data,
			}

			c.JSON(200, gin.H{
				"message": result,
			})
		})
		user.DELETE("/deleteUser/:id", middleware.AdminMiddleware(), func(c *gin.Context) {
			db, _ := database.InitDatabase()
			defer db.Close()
			id := c.Param("id")

			db.Where("id = ?", id).Delete(userModel.User{})
			c.JSON(200, gin.H{
				"message": "success",
			})
		})
		user.GET("/getUser/:id", func(c *gin.Context) {
			db, _ := database.InitDatabase()
			defer db.Close()
			userID := c.Param("id")
			userData := userModel.User{}
			roleData := role.Role{}
			if resultUser := db.Where("id = ?", userID).First(&userData); resultUser.Error != nil {
				c.JSON(422, gin.H{
					"message": resultUser.Error,
					"id":      userID,
				})

				return
			}

			if resultRole := db.Where("id = ?", userData.RoleID).First(&roleData); resultRole.Error != nil {
				c.JSON(422, gin.H{
					"message": resultRole.Error,
				})
			}

			response := userModel.UserGetResponse{
				RoleName: roleData.Data.RoleName,
				Username: userData.Data.Username,
				Email:    userData.Data.Email,
				UserID:   userData.ID.String(),
			}

			c.JSON(200, response)
			return
		})
		user.GET("/listUser", func(c *gin.Context) {
			db, _ := database.InitDatabase()
			defer db.Close()

			userData := []userModel.User{}
			if result := db.Limit(5).Find(&userData); result.Error != nil {
				c.JSON(422, gin.H{
					"message": "Error",
				})

				return
			}

			response := []userModel.UserGetListResponse{}
			for _, v := range userData {
				response = append(response, userModel.UserGetListResponse{
					Email:    v.Data.Email,
					Status:   v.Data.Status,
					Username: v.Data.Username,
				})
			}

			link := "http://" + os.Getenv("HOST") + ":" + os.Getenv("PORT") + "/user/listUser/2"
			c.JSON(200, gin.H{
				"data": response,
				"next": link,
			})
		})
		user.GET("/listUser/:page", func(c *gin.Context) {
			db, _ := database.InitDatabase()
			defer db.Close()
			pageString := c.Param("page")
			page, _ := strconv.Atoi(pageString)

			var count float64
			if result := db.Model(&userModel.User{}).Count(&count); result.Error != nil {
				c.JSON(422, gin.H{
					"message": "Error",
				})

				return
			}
			result := count / float64(5)
			pages := int(math.Ceil(result))
			userData := []userModel.User{}

			if page <= pages {
				if result := db.Limit(5).Offset((page * 5) - 5).Find(&userData); result.Error != nil {
					c.JSON(422, gin.H{
						"message": "Error",
					})
					return
				}

				response := []userModel.UserGetListResponse{}
				for _, v := range userData {
					response = append(response, userModel.UserGetListResponse{
						Email:    v.Data.Email,
						Status:   v.Data.Status,
						Username: v.Data.Username,
					})
				}

				if (page + 1) <= pages {
					link := "http://" + os.Getenv("HOST") + ":" + os.Getenv("PORT") + "/user/listUser/" + strconv.Itoa(page+1)
					linkPrev := "http://" + os.Getenv("HOST") + ":" + os.Getenv("PORT") + "/user/listUser/" + strconv.Itoa(page-1)

					if page < 2 {
						c.JSON(200, gin.H{
							"data": response,
							"next": link,
						})

						return
					}
					c.JSON(200, gin.H{
						"data": response,
						"next": link,
						"prev": linkPrev,
					})

					return
				}

				c.JSON(200, gin.H{
					"data": response,
				})

				return
			}

			c.JSON(422, gin.H{
				"message": "over lap offset",
			})
			return
		})
	}
}
