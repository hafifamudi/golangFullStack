package routes

import (
	"bwastartup/campaign"
	"bwastartup/db"
	"bwastartup/middleware"
	"bwastartup/user"
	webHandler "bwastartup/web/handler"

	"github.com/gin-gonic/gin"
)

func CampaignWebRoutes(route *gin.Engine) {
	//load campaign repo and service
	campaignRepository := campaign.NewRepository(db.DbConfig())
	campaignService := campaign.NewService(campaignRepository)
	//load user service depedency
	userRepository := user.NewRepository(db.DbConfig())
	userService := user.NewService(userRepository)

	CampaignWebHandler := webHandler.NewCampaignHandler(campaignService, userService)
	route.GET("/campaigns", middleware.AuthAdminMiddleware(), CampaignWebHandler.Index)
	route.GET("/campaigns/new", middleware.AuthAdminMiddleware(), CampaignWebHandler.New)
	route.POST("/campaigns", middleware.AuthAdminMiddleware(), CampaignWebHandler.Create)
	route.GET("/campaigns/image/:id", middleware.AuthAdminMiddleware(), CampaignWebHandler.NewImage)
	route.POST("/campaigns/image/:id", middleware.AuthAdminMiddleware(), CampaignWebHandler.CreateImage)
	route.GET("/campaigns/edit/:id", middleware.AuthAdminMiddleware(), CampaignWebHandler.Edit)
	route.POST("/campaigns/update/:id", middleware.AuthAdminMiddleware(), CampaignWebHandler.Update)
	route.GET("/campaigns/show/:id", middleware.AuthAdminMiddleware(), CampaignWebHandler.Show)
}
