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
	// campaignService := user.NewService(campaignRepository)
	// useAuth := auth.NewService()

	// campaignHandler := handler.NewUserHandler(userService, useAuth)
	campaign, _ := campaignRepository.FindByUserID(1)

	for _, campaigns := range campaign {
		fmt.Println(campaigns.Name)
		fmt.Println(campaigns.CampaignImages[0].FileName)
	}
	//setup the router
	// campaign := route.Group("/api/v1")
	// campaign.POST("/users")

}
