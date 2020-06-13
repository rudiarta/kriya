package config

import (
	"github.com/gin-gonic/gin"
	"github.com/rudiarta/kriya/app/controller/auth"
	"github.com/rudiarta/kriya/app/controller/user"
)

func InitRoutes(app *gin.Engine) {

	user.Routes(app)
	auth.Routes(app)

}
