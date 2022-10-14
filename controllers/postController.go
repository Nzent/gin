package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nzent/gin/initializers"
	"github.com/nzent/gin/models"
)

func PostCreate(c *gin.Context) {

	// get data from req body

	// create a post
	post := models.Post{Title: "First title", Body: "First body"}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	// return post
	c.JSON(200, gin.H{
		"post": post,
	})
}
