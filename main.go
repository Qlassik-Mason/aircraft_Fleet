package main

import (
	"controllers"
	"initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	r := gin.Default()

	//Routing for Aircraft
	r.POST("/models/AircraftCreate", controllers.AircraftPostsCreate)
	r.GET("/models/GetAllAircraft", controllers.AircraftGet)
	r.GET("/models/AircraftById/:id", controllers.AircraftSingle)
	r.PUT("/models/AircraftUpdateById/:id", controllers.AircraftUpdate)
	r.DELETE("/models/AircraftDeleteById/:id", controllers.AircraftDelete)

	//Routing for Flight
	r.POST("/models/FlightCreate", controllers.PostFlightCreate)
	r.GET("/models/FlightAll", controllers.FlightGetAll)
	r.GET("/models/FlightByDepartureAirport/:departure_airport", controllers.FlightGetByDepartureAirport)
	r.GET("/models/FlightByArrivalAirport/:arrival_airport", controllers.FlightGetByArrivalAirport)
	r.GET("/models/FlightByDepartureDate/:departure_date", controllers.FlightGetByDepartureDate)
	r.GET("/models/FlightByDepartureTime/:departure_time", controllers.FlightGetByDepartureTime)
	r.GET("/models/FlightByArrivalDate/:arrival_date", controllers.FlightGetByArrivalDate)
	r.GET("/models/FlightByArrivalTime/:arrival_time", controllers.FlightGetByArrivalTime)
	r.PUT("/models/FlightUpdate", controllers.FlightUpdate)
	r.DELETE("/models/FlightDelete", controllers.FlightDelete)

	// Get list of depature airport

	r.GET("/models/ListOfDepartureAirport/:departure_date", controllers.DepartureAirportListGet)
	r.GET("/models/ListOfInflightAircraft/:departure_date", controllers.GetInflightAircraftWithIntimeFlight)

	r.Run()
}
