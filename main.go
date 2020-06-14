package main

import (
	"github.com/gin-gonic/gin"
	. "github.com/rudiarta/kriya/config"
)

func main() {
	App := gin.Default()
	gin.SetMode(DotEnvVariable("APP_ENV"))
	InitRoutes(App)
	App.Run(":" + DotEnvVariable("PORT"))
}
