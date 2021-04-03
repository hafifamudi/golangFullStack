package main

import (
	"bwastartup/routes"
	routesWeb "bwastartup/web/routes"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	//load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	//setup the router
	router := gin.Default()
	//set cookie
	cookieStore := cookie.NewStore([]byte(os.Getenv("SECRET_KEY")))
	//using cors, sessions with gin middleware for handle the outgoing request
	router.Use(cors.Default())
	router.Use(sessions.Sessions("samawaengineer", cookieStore))

	//add image routing
	router.Static("/images", "./images")
	// load css, js, images and webfonts static files
	router.Static("/css", "./web/assets/css")
	router.Static("/js", "./web/assets/js")
	router.Static("/webfonts", "./web/assets/webFonts")
	router.Static("/image", "./web/assets/image")

	//register user routes
	routes.UserRoutes(router)
	//load templates
	router.HTMLRender = loadTemplates("./web/templates")
	//register campaign route
	routes.CampaignRoutes(router)
	//registrer transaction routes
	routes.TransactionRoutes(router)
	//register the user web route
	routesWeb.UserWebRoutes(router)
	//register the campaign web route
	routesWeb.CampaignWebRoutes(router)
	//register the transaction web route
	routesWeb.TransactionWebRoutes(router)

	//run the app
	router.Run(":5000")
}

//load template function
func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/**/*.html")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}
