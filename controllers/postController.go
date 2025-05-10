package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/saifulislam07/go-curd/initializers"
	"github.com/saifulislam07/go-curd/models"
)

func PostsCreate(c *gin.Context) {

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
