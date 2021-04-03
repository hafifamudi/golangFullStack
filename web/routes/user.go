package routes

import (
	"bwastartup/db"
	"bwastartup/middleware"
	"bwastartup/user"
	webHandler "bwastartup/web/handler"

	"github.com/gin-gonic/gin"
)

func UserWebRoutes(route *gin.Engine) {
	//load user service depedency
	userRepository := user.NewRepository(db.DbConfig())
	userService := user.NewService(userRepository)

	userWebHandler := webHandler.NewUserHandler(userService)
	route.GET("/users", middleware.AuthAdminMiddleware(), userWebHandler.Index)
	route.GET("/users/new", middleware.AuthAdminMiddleware(), userWebHandler.New)
	route.POST("/users", middleware.AuthAdminMiddleware(), userWebHandler.Create)
	route.GET("/users/edit/:id", middleware.AuthAdminMiddleware(), userWebHandler.Edit)
	route.POST("/users/update/:id", middleware.AuthAdminMiddleware(), userWebHandler.Update)
	route.GET("/users/avatar/:id", middleware.AuthAdminMiddleware(), userWebHandler.NewAvatar)
	route.POST("/users/avatar/:id", middleware.AuthAdminMiddleware(), userWebHandler.UploadAvatar)
}
