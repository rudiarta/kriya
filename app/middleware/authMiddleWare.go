package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rudiarta/kriya/app/model/role"
	userModel "github.com/rudiarta/kriya/app/model/user"
	"github.com/rudiarta/kriya/app/service"
	"github.com/rudiarta/kriya/config/database"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Request.Header["Authorization"]) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized.",
			})
			c.Abort()
			return
		}
		token := c.Request.Header["Authorization"][0]
		rune := []rune(token)
		tokenBearer := string(rune[7:])
		data, ok := service.VerifyToken(tokenBearer)
		if ok != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized.",
			})
			c.Abort()
			return
		}
		claims, err := data.Claims.(jwt.MapClaims)
		if !err && !data.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized.",
			})
			c.Abort()
			return
		}

		db, _ := database.InitDatabase()
		var roleData role.Role
		var userData userModel.User
		db.Where("id = ?", claims["user_id"]).First(&userData)
		db.Where("id = ?", userData.RoleID).First(&roleData)
		if roleData.Data.RoleName != "Admin" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
