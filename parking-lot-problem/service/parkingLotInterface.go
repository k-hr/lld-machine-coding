package service

import "lld-machine-coding/parking-lot-problem/models"

type ParkingLotImpl interface {
	ParkCar(car *models.Car) (int, error)
	Leave(slotID int) error
	GetSlotByRegistration(registration string) (int, error)
	ListOccupiedSlots() []int
	GetCarsByColor(color string) []string
}
