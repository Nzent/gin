package main

import (
	"github.com/nzent/gin/initializers"
	"github.com/nzent/gin/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
