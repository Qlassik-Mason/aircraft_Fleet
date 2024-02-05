package controllers

import (
	"initializers"
	"models"

	"github.com/gin-gonic/gin"
)

var body struct {
	departure_airport string `gorm:"primarykey;size:16"`
	DepartureTime     string
	DepartureDate     string
	FlightNum         string
	ArrivalAirport    string
	ArrivalDate       string
	ArrivalTime       string
	NumberofAircraft  string
	InFlightAircraft  string
}

func PostFlightCreate(c *gin.Context) {
	// Get data of req body
	c.Bind(&body)

	//create  Flight post
	FlightPost := models.Flight{DepartureAirport: body.departure_airport, DepartureTime: body.DepartureTime, DepartureDate: body.DepartureDate, FlightNum: body.FlightNum,
		ArrivalAirport: body.ArrivalAirport, ArrivalDate: body.ArrivalDate, ArrivalTime: body.ArrivalTime, NumberofAircraft: body.NumberofAircraft,
		InFlightAircraft: body.InFlightAircraft}

	result := initializers.DB.Create(&FlightPost)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"FlightCreated": FlightPost,
	})
}

func FlightGetAll(c *gin.Context) {
	// Get all flight data
	var GetFlight []models.Flight
	initializers.DB.Find(&GetFlight)

	//Respond with
	c.JSON(200, gin.H{
		"Get all FLight": GetFlight,
	})
}

func FlightGetByDepartureAirport(c *gin.Context) {

	//Get each Aircraft by departure airport
	var SingleFlight models.Flight
	initializers.DB.Where("departure_airport = ?", c.Param("departure_airport")).Find(&SingleFlight)

	//Respond with
	c.JSON(200, gin.H{
		"Get all Flight by departure airport": SingleFlight,
	})

}

func FlightGetByArrivalAirport(c *gin.Context) {

	//Get each Aircraft by arrival airport
	var SingleFlight models.Flight
	initializers.DB.Where("arrival_airport = ?", c.Param("arrival_airport")).Find(&SingleFlight)

	//Respond with
	c.JSON(200, gin.H{
		"Get all Flight by arrival airport": SingleFlight,
	})

}

func FlightGetByDepartureDate(c *gin.Context) {

	//Get each Aircraft by departure date
	var SingleFlight models.Flight
	initializers.DB.Where("departure_date = ?", c.Param("departure_date")).Find(&SingleFlight)

	//Respond with
	c.JSON(200, gin.H{
		"Get all Flight by departure date": SingleFlight,
	})
}

func FlightGetByDepartureTime(c *gin.Context) {

	//Get each Aircraft by Departure time
	var SingleFlight models.Flight
	initializers.DB.Where("departure_time = ?", c.Param("departure_time")).Find(&SingleFlight)

	//Respond with
	c.JSON(200, gin.H{
		"Get all Flight by departure time": SingleFlight,
	})

}

func FlightGetByArrivalDate(c *gin.Context) {

	//Get each Aircraft by arrival date
	var SingleFlight models.Flight
	initializers.DB.Where("arrival_date = ?", c.Param("arrival_date")).Find(&SingleFlight)

	//Respond with
	c.JSON(200, gin.H{
		"Get all Flight by arrival date": SingleFlight,
	})

}

func FlightGetByArrivalTime(c *gin.Context) {

	//Get each Aircraft by arrival time
	var SingleFlight models.Flight
	initializers.DB.Where("arrival_time = ?", c.Param("arrival_time")).Find(&SingleFlight)

	//Respond with
	c.JSON(200, gin.H{
		"Get all Flight by arrival time": SingleFlight,
	})

}

// Get list of departure airport
func DepartureAirportListGet(c *gin.Context) {

	// Get first matched record
	var GetFlight models.Flight
	initializers.DB.Model(&models.Flight{}).Select("departure_airport, numberof_aircraft").Find(&GetFlight)

	//Respond with
	c.JSON(200, gin.H{
		"Get all departure airport and aircraft in flight list": GetFlight,
	})

}

func GetInflightAircraftWithIntimeFlight(c *gin.Context) {

	// Get first matched record
	var GetFlight []models.Flight
	initializers.DB.Model(&models.Flight{}).Select("in_flight_aircraft").Find(&GetFlight)

	//Respond with
	c.JSON(200, gin.H{
		"Inflight aircraft list and time in minutes": GetFlight,
	})

}

func FlightUpdate(c *gin.Context) {

	//find the table  were to  updating
	var SingleFlight []models.Flight
	//initializers.DB.First(&SingleFlight)

	//Update the value by arrival time
	initializers.DB.Model(&SingleFlight).Updates(models.Flight{DepartureAirport: body.departure_airport, DepartureTime: body.DepartureTime, DepartureDate: body.DepartureDate, FlightNum: body.FlightNum,
		ArrivalAirport: body.ArrivalAirport, ArrivalDate: body.ArrivalDate, ArrivalTime: body.ArrivalTime, NumberofAircraft: body.NumberofAircraft,
		InFlightAircraft: body.InFlightAircraft}).Where("arrival_time = ?", c.Param("arrival_time")).Find(&SingleFlight)

	//Respond with
	c.JSON(200, gin.H{
		"Flights": SingleFlight,
	})
}

func FlightDelete(c *gin.Context) {

	//Delete  Flight
	initializers.DB.Delete(&models.Flight{}).Where("arrival_time = ?", c.Param("arrival_time")).Find(&models.Flight{})
	c.Status(200)
}
