package models

//import "gorm.io/gorm"

type Flight struct {
	//gorm.Model

	DepartureAirport string
	DepartureTime    string
	DepartureDate    string
	FlightNum        string
	ArrivalAirport   string
	ArrivalDate      string
	ArrivalTime      string
	NumberofAircraft string
	InFlightAircraft string
}
