package main

import (
	"bwastartup/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	//setup the router
	router := gin.Default()

	//using cors with gin middleware for handle the outgoing request
	router.Use(cors.Default())

	//add image routing
	router.Static("/images", "./images")
	//register user routes
	routes.UserRoutes(router)
	//register campaign routes
	routes.CampaignRoutes(router)
	//registrer transaction routes
	routes.TransactionRoutes(router)

	//run the app
	router.Run(":5000")
}

/**
input dari user
handler, mapping input dari user -> struct input
service : melakukan mapping dari struct input struct User
repository -> interaksi dengan DB
db
**/
