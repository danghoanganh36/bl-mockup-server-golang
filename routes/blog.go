package routes

import (
	"bl-mockup-server-golang/models"
	"bl-mockup-server-golang/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterBlogRoutes(router *gin.Engine) {
	r := router.Group("/api/blogs")
	{
		r.GET("/", func(c *gin.Context) {
			var blogs []models.Blog
			database.DB.Preload("Category").Find(&blogs)
			c.JSON(http.StatusOK, blogs)
		})

		r.POST("/", func(c *gin.Context) {
			var blog models.Blog
			if err := c.ShouldBindJSON(&blog); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			database.DB.Create(&blog)
			c.JSON(http.StatusCreated, blog)
		})
	}
}