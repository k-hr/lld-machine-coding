package service

import (
	"lld-machine-coding/parking-lot-problem/models"
	"sync"
)

type parkingLot struct {
	Capacity     int
	ParkingSlots []*models.Slot
	Mutex        sync.Mutex
}

func NewParkingLot(capacity int) ParkingLotImpl {
	parkingSlots := make([]*models.Slot, capacity)
	for i := 0; i < capacity; i++ {
		parkingSlots[i] = &models.Slot{ID: i + 1, Car: nil}
	}
	return &parkingLot{
		Capacity:     capacity,
		ParkingSlots: parkingSlots,
	}
}
