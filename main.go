package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rudiarta/kriya/app/controller/auth"
	"github.com/rudiarta/kriya/app/controller/user"
)

func main() {
	App := gin.Default()

	user.Routes(App)
	auth.Routes(App)

	App.Run()
}
