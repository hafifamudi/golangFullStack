package routes

import (
	"bwastartup/auth"
	"bwastartup/campaign"
	"bwastartup/db"
	"bwastartup/handler"
	"bwastartup/middleware"
	"bwastartup/user"

	"github.com/gin-gonic/gin"
)

func CampaignRoutes(route *gin.Engine) {
	//setup the handler,service and repo
	campaignRepository := campaign.NewRepository(db.DbConfig())
	campaignService := campaign.NewService(campaignRepository)
	campaignHandler := handler.NewCampaignHandler(campaignService)

	// userService and userRepo for middleware
	userRepository := user.NewRepository(db.DbConfig())
	userService := user.NewService(userRepository)
	useAuth := auth.NewService()

	//setup the router
	campaign := route.Group("/api/v1")
	campaign.GET("/campaigns", campaignHandler.GetCampaigns)
	campaign.GET("/campaigns/:id", campaignHandler.GetCampaign)
	campaign.POST("/campaigns", middleware.AuthMiddleware(useAuth, userService), campaignHandler.CreateCampaign)
	campaign.PUT("/campaigns/:id", middleware.AuthMiddleware(useAuth, userService), campaignHandler.UpdateCampaign)
	campaign.POST("/campaign-images", middleware.AuthMiddleware(useAuth, userService), campaignHandler.UploadImage)
}
