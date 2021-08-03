package main

import (
	"go/controllers"
	"go/data"
	"go/middleware/header"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	data.Migrate()
	data.CreateInitUser()

	// Middlewares
	g.Use(header.Nocache)
	g.Use(header.Secure)

	// 404 Handler
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// Setting routes
	// user := g.Group("/user")
	// {
	// 	user.POST("/login", controllers.Login)
	// 	user.POST("/logout", controllers.Logout)
	// }
	set := g.Group("api/set")
	{
		set.POST("category", controllers.SetCategory)
		set.POST("task", controllers.SetTask)
	}
	fetch := g.Group("api/fetch")
	{
		fetch.GET("/category", controllers.FetchCategory)
		fetch.POST("/task", controllers.FetchTask)
	}

	g.Run(":7777")
}
