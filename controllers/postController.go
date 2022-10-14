package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nzent/gin/initializers"
	"github.com/nzent/gin/models"
)

func PostCreate(c *gin.Context) {

	// get data from req body
	var body struct {
		ID uint
		Title string
		Body  string
	}

	c.Bind(&body)
	// create a post
	post := models.Post{ Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	// return post
	c.JSON(200, gin.H{
		"post": post.ID,
	})
}
