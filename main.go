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
	r.POST("/posts", controllers.PostCreate)
	r.Run() // listen and serve on 0.0.0.0:8080
}
