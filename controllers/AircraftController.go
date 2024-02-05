package controllers

import (
	"initializers"
	"models"

	"github.com/gin-gonic/gin"
)

func AircraftPostsCreate(c *gin.Context) {

	//Get data off req body

	var body struct {
		Manufacturer string
		UniqueSerial string
	}
	c.Bind(&body)

	//Create a post

	AircraftPost := models.Aircaft{Manufacturer: body.Manufacturer, UniqueSerial: body.UniqueSerial}
	result := initializers.DB.Create(&AircraftPost)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"Aircraft": AircraftPost,
	})
}

func AircraftGet(c *gin.Context) {
	//Get all Aircraft data
	var GetAircraft []models.Aircaft
	initializers.DB.Find(&GetAircraft)

	//Respond with them
	c.JSON(200, gin.H{
		"Aircraft": GetAircraft,
	})

}

func AircraftSingle(c *gin.Context) {
	//Get the id of the url
	id := c.Param("id")
	//Get each Aircraft by id
	var SingleAircraft models.Aircaft
	initializers.DB.First(&SingleAircraft, id)

	//Respond with them
	c.JSON(200, gin.H{
		"Aircraft": SingleAircraft,
	})

}

func AircraftUpdate(c *gin.Context) {
	//Get the id off the url
	id := c.Param("id")

	//Get the data off req body
	var body struct {
		Manufacturer string
		UniqueSerial string
	}
	c.Bind(&body)

	//find the table  were to  updating
	var SingleAircraft models.Aircaft
	initializers.DB.First(&SingleAircraft, id)

	//Update
	initializers.DB.Model(&SingleAircraft).Updates(models.Aircaft{Manufacturer: body.Manufacturer, UniqueSerial: body.UniqueSerial})

	//Respond with
	c.JSON(200, gin.H{
		"Aircraft": SingleAircraft,
	})

}

func AircraftDelete(c *gin.Context) {

	//Get the id off the url
	id := c.Param("id")

	//Delete  Aircraft
	initializers.DB.Delete(&models.Aircaft{}, id)
	c.Status(200)
}
