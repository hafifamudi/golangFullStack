package routes

import (
	"bwastartup/auth"
	"bwastartup/db"
	"bwastartup/handler"
	"bwastartup/user"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	//setup the handler,service and repo
	userRepository := user.NewRepository(db.DbConfig())
	userService := user.NewService(userRepository)
	useAuth := auth.NewService()

	userHandler := handler.NewUserHandler(userService, useAuth)

	//setup the router
	user := route.Group("/api/v1")
	user.POST("/users", userHandler.RegisterUser)
	user.POST("/sessions", userHandler.Login)
	user.POST("/email_checkers", userHandler.CheckEmailAvailability)
	user.POST("/avatars", userHandler.UploadAvatar)
}
