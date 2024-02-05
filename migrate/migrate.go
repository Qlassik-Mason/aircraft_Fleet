package main

import (
	"initializers"
	"models"
)

func init() {

	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {

	initializers.DB.AutoMigrate(&models.Aircaft{})
	initializers.DB.AutoMigrate(&models.Flight{})
}
