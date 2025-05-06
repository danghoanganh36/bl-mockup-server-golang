package main

import (
	"bl-mockup-server-golang/config"
	"bl-mockup-server-golang/database"
	"bl-mockup-server-golang/routes"
	"bl-mockup-server-golang/scripts"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnvVariables()
	database.ConnectDB()
	scripts.ImportBlogsFromJSON()
	scripts.ImportMetricsFromJSON()
	scripts.ImportCategoriesFromJSON()

	r := gin.Default()
	routes.RegisterBlogRoutes(r)

	r.Run(":2808")
}
