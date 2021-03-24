package routes

import (
	"bwastartup/campaign"
	"bwastartup/db"
	"bwastartup/handler"

	"github.com/gin-gonic/gin"
)

func CampaignRoutes(route *gin.Engine) {
	//setup the handler,service and repo
	campaignRepository := campaign.NewRepository(db.DbConfig())
	campaignService := campaign.NewService(campaignRepository)
	// useAuth := auth.NewService()

	campaignHandler := handler.NewCampaignHandler(campaignService)

	//setup the router
	campaign := route.Group("/api/v1")
	campaign.GET("/campaigns", campaignHandler.GetCampaigns)

}
