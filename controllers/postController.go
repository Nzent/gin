package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nzent/gin/initializers"
	"github.com/nzent/gin/models"
)

// Create post
func CreatePost(c *gin.Context) {

	// get data from req body
	var body struct {
		ID    uint
		Title string
		Body  string
	}

	c.Bind(&body)
	// create a post
	post := models.Post{Title: body.Title, Body: body.Body}

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

// Get all posts
func GetAllPosts(c *gin.Context) {
	// get posts
	var posts []models.Post
	initializers.DB.Find(&posts)
	// respond with them
	// return post

	c.JSON(200, gin.H{
		"post": posts,
	})
}

// Get posts
func GetSinglePosts(c *gin.Context) {
	// get id from the url
	id := c.Param("id")
	// get post
	var post models.Post
	initializers.DB.Find(&post, id)
	// respond with them
	// return post

	c.JSON(200, gin.H{
		"post": post,
	})
}

// Update post
func UpdatePost(c *gin.Context) {
	// get id from the url
	id := c.Param("id")
	// get data from req body
	var body struct {
		ID    uint
		Title string
		Body  string
	}
	c.Bind(&body)
	// find the post where updating
	var post models.Post
	initializers.DB.Find(&post, id)
	// update the post
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})
	// respond

	c.JSON(200, gin.H{
		"post": post,
	})
}

// Delete a post
func DeletePost(c *gin.Context) {
	// get id from the url
	id := c.Param("id")

	// find the post where updating
	var post models.Post
	initializers.DB.Find(&post, id)
	// delete the post
	initializers.DB.Delete(&models.Post{}, id)
	// respond

	c.Status(200)
}
