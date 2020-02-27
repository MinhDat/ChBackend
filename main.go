package main

import (
	"log"

	"github.com/gin-gonic/gin"

	TaskController "ChGo/controllers"
	"ChGo/db"
	auth "ChGo/middlewares"
)

func main() {
	log.Println("Starting server..")

	db.Init()

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		tasks := v1.Group("/tasks")
		{
			tasks.GET("/", TaskController.GetTasks)
			tasks.POST("/", TaskController.CreateTask)
			tasks.PUT("/:id", TaskController.UpdateTask)
			tasks.DELETE("/:id", TaskController.DeleteTask)
		}

		noauth := v1.Group("/noauth")
		{
			noauth.POST("/login", auth.Sign)

			// noauth.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
			// 	claims := jwt.ExtractClaims(c)
			// 	log.Printf("NoRoute claims: %#v\n", claims)
			// 	c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
			// })

		}

	}

	// if err := http.ListenAndServe(":"+port, r); err != nil {
	// 	log.Fatal(err)
	// }

	r.Run()
}
