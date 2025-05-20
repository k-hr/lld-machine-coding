package service

import (
	"errors"
	"lld-machine-coding/parking-lot-problem/models"
)

func (p *parkingLot) ParkCar(car *models.Car) (int, error) {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()
	for _, slot := range p.ParkingSlots {
		if slot.Car == nil {
			slot.Car = car
			return slot.ID, nil
		}
	}
	return -1, errors.New("parking lot is full")
}

func (p *parkingLot) Leave(slotID int) error {
	p.Mutex.Lock()
	defer p.Mutex.Unlock()
	if slotID < 1 || slotID > p.Capacity {
		return errors.New("invalid slot number")
	}
	slot := p.ParkingSlots[slotID-1]
	if slot.Car == nil {
		return errors.New("slot is already empty")
	}
	slot.Car = nil
	return nil
}

func (p *parkingLot) GetSlotByRegistration(registration string) (int, error) {
	for _, slot := range p.ParkingSlots {
		if slot.Car != nil && slot.Car.RegistrationNumber == registration {
			return slot.ID, nil
		}
	}
	return -1, errors.New("car not found")
}

func (p *parkingLot) ListOccupiedSlots() []int {
	var occupied []int
	for _, slot := range p.ParkingSlots {
		if slot.Car != nil {
			occupied = append(occupied, slot.ID)
		}
	}
	return occupied
}

func (p *parkingLot) GetCarsByColor(color string) []string {
	var cars []string
	for _, slot := range p.ParkingSlots {
		if slot.Car != nil && slot.Car.Color == color {
			cars = append(cars, slot.Car.RegistrationNumber)
		}
	}
	return cars
}
