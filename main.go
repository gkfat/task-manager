package main

import (
	"go/controllers"
	"go/data"
	"go/middleware/header"
	"go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.New()

	// Connect to DB & Migrate
	db := data.ConnectDB()
	tableList := []interface{}{
		&models.User{},
		&models.Task{},
		&models.Category{},
	}
	db.AutoMigrate(tableList...)

	// Middlewares
	g.Use(header.Nocache)
	g.Use(header.Secure)

	// 404 Handler
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// Setting routes
	user := g.Group("/user")
	{
		user.POST("/login", controllers.Login)
		user.POST("/logout", controllers.Logout)
	}
	set := g.Group("/set")
	// set.Use(middleware.IdentityHandler)
	{
		set.POST("category", controllers.SetCategory)
		set.POST("task", controllers.SetTask)
	}
	fetch := g.Group("/fetch")
	// fetch.Use(middleware.IdentityHandler)
	{
		fetch.POST("/category", controllers.FetchCategory)
		fetch.POST("/task", controllers.FetchTask)
	}

	g.Run(":7777")
}
