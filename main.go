package main

import (
	"bwastartup/routes"
	routesWeb "bwastartup/web/routes"
	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func main() {

	//setup the router
	router := gin.Default()

	//using cors with gin middleware for handle the outgoing request
	router.Use(cors.Default())

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
