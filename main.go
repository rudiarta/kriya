package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rudiarta/kriya/config"
)

func main() {
	App := gin.Default()
	gin.SetMode(config.DotEnvVariable("APP_ENV"))
	config.InitRoutes(App)
	App.Run(":" + config.DotEnvVariable("PORT"))
}
