package main

import (
	"bwastartup/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	//setup the router
	router := gin.Default()

	//register user routes
	routes.Routes(router)

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
