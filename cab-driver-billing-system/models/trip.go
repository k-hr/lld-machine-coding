package models

import "time"

type TripType string

type TripStatus string

const (
	IntraCity  TripType = "INTRACITY"
	OutStation TripType = "OUTSTATION"

	Completed        TripStatus = "COMPLETED"
	CanceledDriver   TripStatus = "CANCELED_DRIVER"
	CanceledCustomer TripStatus = "CANCELED_CUSTOMER"
)

type Trip struct {
	DriverName string
	Distance   float64
	StartTime  time.Time
	EndTime    time.Time
	Type       TripType
	TripStatus TripStatus
}
