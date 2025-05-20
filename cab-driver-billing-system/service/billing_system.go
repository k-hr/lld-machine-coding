package service

import (
	"errors"
	"fmt"
	"lld-machine-coding/cab-driver-billing-system/models"
	"time"
)

// manage the billing process

// map carType -> fare
var intraCityPricing = map[models.CarCategory]float64{
	models.Economy: 5,
	models.Premium: 12,
	models.Luxury:  15,
}

type BillingSystemService struct {
	drivers map[string]models.Driver
	vehicle map[string]models.Vehicle
	trips   []models.Trip
}

func NewBillingSystemService() *BillingSystemService {
	return &BillingSystemService{
		drivers: make(map[string]models.Driver),
		vehicle: make(map[string]models.Vehicle),
		trips:   []models.Trip{},
	}
}

// AddDriver onboard Driver
func (bs *BillingSystemService) AddDriver(name, phone, id string) {
	bs.drivers[name] = models.Driver{
		Name:  name,
		Phone: phone,
		ID:    id,
	}

	fmt.Printf("Driver Onboarded: %s\n", name)
}

// AddVehicle onboard vehicle
func (bs *BillingSystemService) AddVehicle(id, model string, category models.CarCategory) {
	bs.vehicle[id] = models.Vehicle{
		ID:       id,
		Model:    model,
		Category: category,
	}
	fmt.Printf("Vehicle Onboarded: %s (%s)\n", id, model)
}

// MapDriverToVehicle Driver<>Vehicle
func (bs *BillingSystemService) MapDriverToVehicle(driverName, vehicleId string) error {
	driver, ok := bs.drivers[driverName]
	if !ok {
		return errors.New("driver not found")
	}

	if _, ok := bs.vehicle[vehicleId]; !ok {
		return errors.New("vehicle not found")
	}

	driver.VehicleID = vehicleId
	bs.drivers[driverName] = driver

	fmt.Printf("Driver %s Mapped to vehicle: %s\n", driverName, vehicleId)

	return nil
}

// AddTrip add the trip
func (bs *BillingSystemService) AddTrip(distance float64, start, end int64, tripType models.TripType, tripStatus models.TripStatus, driverName string) error {
	if _, ok := bs.drivers[driverName]; !ok {
		return errors.New("driver not found")
	}

	trip := models.Trip{
		DriverName: driverName,
		Distance:   distance,
		StartTime:  time.Unix(start, 0),
		EndTime:    time.Unix(end, 0),
		Type:       tripType,
		TripStatus: tripStatus,
	}

	bs.trips = append(bs.trips, trip)
	fmt.Printf("Trip Added for driver: %s\n", trip.DriverName)

	return nil
}

// GetBill calculates driver bill
func (bs *BillingSystemService) GetBill(driverName string) (float64, error) {
	var totalBill float64
	driver, ok := bs.drivers[driverName]
	if !ok {
		return 0, errors.New("driver not found")
	}

	vehicle, ok := bs.vehicle[driver.VehicleID]
	if !ok {
		return 0, errors.New("vehicle not found")
	}

	totalBill = 0.0
	for _, trip := range bs.trips {
		if trip.DriverName != driverName {
			continue
		}

		ratePerKm := intraCityPricing[vehicle.Category]
		if trip.Type == models.OutStation {
			ratePerKm *= 2
		}

		if trip.TripStatus == models.CanceledCustomer {
			totalBill += 10
		} else if trip.TripStatus == models.CanceledDriver {
			totalBill -= 10
		} else {
			totalBill += ratePerKm * trip.Distance
		}
	}

	return totalBill, nil
}
