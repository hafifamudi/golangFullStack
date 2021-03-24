package routes

import (
	"bwastartup/campaign"
	"bwastartup/db"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CampaignRoutes(route *gin.Engine) {
	//setup the handler,service and repo
	campaignRepository := campaign.NewRepository(db.DbConfig())
	campaignService := campaign.NewService(campaignRepository)
	// useAuth := auth.NewService()

	// campaignHandler := handler.NewUserHandler(userService, useAuth)

	campaigns, _ := campaignService.FindCampaigns(8)
	fmt.Println(len(campaigns))
	//setup the router
	// campaign := route.Group("/api/v1")
	// campaign.POST("/users")

}
