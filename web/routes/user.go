package routes

import (
	webHandler "bwastartup/web/handler"

	"github.com/gin-gonic/gin"
)

func UserWebRoutes(route *gin.Engine) {
	userWebHandler := webHandler.NewUserHandler()
	route.GET("/users", userWebHandler.Index)
}
