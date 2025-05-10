package main

import (
	"github.com/gin-gonic/gin"
	"github.com/saifulislam07/go-curd/controllers"
	"github.com/saifulislam07/go-curd/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)

	r.Run()
}
