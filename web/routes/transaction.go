package routes

import (
	"bwastartup/campaign"
	"bwastartup/db"
	"bwastartup/payment"
	"bwastartup/transaction"
	webHandler "bwastartup/web/handler"

	"github.com/gin-gonic/gin"
)

func TransactionWebRoutes(route *gin.Engine) {
	//call campaign repo for the dependency injection
	campaignRepository := campaign.NewRepository(db.DbConfig())

	//setup the handler,service and repo
	transactionRepository := transaction.NewRepository(db.DbConfig())
	//setup payment service
	paymentService := payment.NewService()

	//load transacition service depedency
	transactionService := transaction.NewService(transactionRepository, campaignRepository, paymentService)

	TransactionWebHandler := webHandler.NewTransactionHandler(transactionService)
	route.GET("/transactions", TransactionWebHandler.Index)
}
