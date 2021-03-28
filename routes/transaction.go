package routes

import (
	"bwastartup/auth"
	"bwastartup/campaign"
	"bwastartup/db"
	"bwastartup/handler"
	"bwastartup/middleware"
	"bwastartup/payment"
	"bwastartup/transaction"
	"bwastartup/user"

	"github.com/gin-gonic/gin"
)

func TransactionRoutes(route *gin.Engine) {
	//call campaign repo for the dependency injection
	campaignRepository := campaign.NewRepository(db.DbConfig())

	//setup the handler,service and repo
	transactionRepository := transaction.NewRepository(db.DbConfig())
	//setup payment service
	paymentService := payment.NewService()
	transactionService := transaction.NewService(transactionRepository, campaignRepository, paymentService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	// userService and userRepo for middleware
	userRepository := user.NewRepository(db.DbConfig())
	userService := user.NewService(userRepository)
	useAuth := auth.NewService()

	//setup the router
	transaction := route.Group("/api/v1")
	transaction.GET("/campaigns/:id/transaction", middleware.AuthMiddleware(useAuth, userService), transactionHandler.GetCampaignTransaction)
	transaction.GET("/transactions", middleware.AuthMiddleware(useAuth, userService), transactionHandler.GetUserTransactions)
	transaction.POST("/transactions", middleware.AuthMiddleware(useAuth, userService), transactionHandler.CreateTransaction)
	transaction.POST("/transactions/notifications", middleware.AuthMiddleware(useAuth, userService), transactionHandler.GetNotification)
}
