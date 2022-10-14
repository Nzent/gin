package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nzent/gin/controllers"
	"github.com/nzent/gin/initializers"
)

// init function \ run before main
func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

// main function
func main() {
	r := gin.Default()
	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts", controllers.GetAllPosts)
	r.GET("/post/:id", controllers.GetSinglePosts)
	r.PUT("/post/:id", controllers.UpdatePost)
	r.DELETE("/post/:id", controllers.DeletePost)
	
	r.Run() // listen and serve on 0.0.0.0:8080
}
