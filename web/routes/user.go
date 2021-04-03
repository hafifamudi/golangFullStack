package routes

import (
	"bwastartup/db"
	"bwastartup/user"
	webHandler "bwastartup/web/handler"

	"github.com/gin-gonic/gin"
)

func UserWebRoutes(route *gin.Engine) {
	//load user service depedency
	userRepository := user.NewRepository(db.DbConfig())
	userService := user.NewService(userRepository)

	userWebHandler := webHandler.NewUserHandler(userService)
	route.GET("/users", userWebHandler.Index)
	route.GET("/users/new", userWebHandler.New)
	route.POST("/users", userWebHandler.Create)
	route.GET("/users/edit/:id", userWebHandler.Edit)
	route.POST("/users/update/:id", userWebHandler.Update)
	route.GET("/users/avatar/:id", userWebHandler.NewAvatar)
	route.POST("/users/avatar/:id", userWebHandler.UploadAvatar)
}
