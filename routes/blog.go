package routes

import (
	"bl-mockup-server-golang/database"
	"bl-mockup-server-golang/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterBlogRoutes(router *gin.Engine) {
	blogs := router.Group("/api/blogs")
	{
		blogs.GET("/", func(c *gin.Context) {
			var blogs []models.Blog
			database.DB.Preload("Category").Find(&blogs)
			c.JSON(http.StatusOK, blogs)
		})

		blogs.POST("/", func(c *gin.Context) {
			var blog models.Blog
			if err := c.ShouldBindJSON(&blog); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			database.DB.Create(&blog)
			c.JSON(http.StatusCreated, blog)
		})
	}

	categories := router.Group("/api/categories")
	{
		categories.GET("/", func(c *gin.Context) {
			var categories []models.Category
			database.DB.Find(&categories)
			c.JSON(http.StatusOK, categories)
		})
	}
}
