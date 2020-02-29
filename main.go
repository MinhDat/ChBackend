package main

import (
	"log"

	"github.com/gin-gonic/gin"

	controllers "ChGo/controllers"
	"ChGo/db"
	middlewares "ChGo/middlewares"
)

func main() {
	log.Println("Starting server..")

	db.Init()

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		tasks := v1.Group("/tasks", middlewares.Auth)
		{
			tasks.GET("/", controllers.GetTasks)
			tasks.POST("/", controllers.CreateTask)
			tasks.PUT("/:id", controllers.UpdateTask)
			tasks.DELETE("/:id", controllers.DeleteTask)
		}

		noauth := v1.Group("/noauth")
		{
			noauth.POST("/login", middlewares.Sign)
			noauth.POST("/register", controllers.Register)
			noauth.GET("/refresh", middlewares.Refresh)
		}

	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	r.Run()
}
