package main

import (
	"github.com/saifulislam07/go-curd/initializers"
	"github.com/saifulislam07/go-curd/models"
)

func init() {
	//
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	//
	initializers.DB.AutoMigrate(&models.Post{})
}
