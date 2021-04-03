package routes

import (
	"bwastartup/campaign"
	"bwastartup/db"
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
	route.GET("/campaigns", CampaignWebHandler.Index)
	route.GET("/campaigns/new", CampaignWebHandler.New)

}
