package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/saifulislam07/go-curd/initializers"
	"github.com/saifulislam07/go-curd/models"
)

func PostsCreate(c *gin.Context) {
	// insert data
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// get post data
	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// get post by id

	var post models.Post
	id := c.Param("id")
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// update post data
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	var post models.Post
	id := c.Param("id")
	initializers.DB.First(&post, id)

	post.Title = body.Title
	post.Body = body.Body

	initializers.DB.Model(&post).Updates(models.Post{
		Title: post.Title,
		Body:  post.Body,
	})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context) {
	// delete post data
	var post models.Post
	id := c.Param("id")
	initializers.DB.Delete(&post, id)

	c.Status(200)
}
