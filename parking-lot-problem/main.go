package main

import (
	"awesomeProject/models"
	"awesomeProject/service"
	"fmt"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {

	parkingLot := service.NewParkingLot(10)

	//park car
	_, _ = parkingLot.ParkCar(&models.Car{RegistrationNumber: "KA-01-HH-1234", Color: "White"})
	_, _ = parkingLot.ParkCar(&models.Car{RegistrationNumber: "KA-01-HH-9999", Color: "Black"})
	_, _ = parkingLot.ParkCar(&models.Car{RegistrationNumber: "KA-01-HH-1999", Color: "Silver"})
	_, _ = parkingLot.ParkCar(&models.Car{RegistrationNumber: "KA-01-HH-1998", Color: "Grey"})

	// list filled slots
	fmt.Println("Occupied ParkingSlots:", parkingLot.ListOccupiedSlots())
	// get slot by reg number
	slot, err := parkingLot.GetSlotByRegistration("KA-01-HH-1234")
	if err == nil {
		fmt.Println("Car KA-01-HH-1234 is parked at slot:", slot)
	}

	err = parkingLot.Leave(2)
	if err != nil {
		fmt.Println(err)
	}

	// list filled slots
	fmt.Println("Occupied ParkingSlots:", parkingLot.ListOccupiedSlots())

	_, _ = parkingLot.ParkCar(&models.Car{RegistrationNumber: "KA-01-HH-1996", Color: "Orange"})

	// list filled slots
	fmt.Println("Occupied ParkingSlots:", parkingLot.ListOccupiedSlots())

	// get car by given colour
	slot, _ = parkingLot.GetSlotByRegistration("KA-01-HH-1996")
	fmt.Println("Cars of color Orange:", parkingLot.GetCarsByColor("Orange"), "and slot: ", slot)
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
