package models

//import "gorm.io/gorm"

type Flight struct {
	//gorm.Model
	DepartureAirport string `gorm:"primarykey;size:16"`
	DepartureTime    string
	DepartureDate    string
	FlightNum        string
	ArrivalAirport   string
	ArrivalDate      string
	ArrivalTime      string
	NumberofAircraft string
	InFlightAircraft string
}
