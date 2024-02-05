package models

import "gorm.io/gorm"

type Aircaft struct {
	gorm.Model

	Manufacturer string
	UniqueSerial string
}
