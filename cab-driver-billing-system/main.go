package main

import (
	"fmt"
	"test/service"
)

func main() {
	// define a billing service
	bs := service.NewBillingSystemService()

	// onboard driver
	bs.AddDriver("Sachin", "+91-9998887776", "DL_01")
	bs.AddDriver("Ramesh", "+91-9998887775", "DL_02")
	bs.AddDriver("Manjunath", "+91-9998887774", "DL_03")

	// add vehicle
	bs.AddVehicle("KA-01-2222", "Maruti", "ECONOMY")
	bs.AddVehicle("KA-01-2223", "Ertiga", "PREMIUM")

	// map driver to vehicle
	err := bs.MapDriverToVehicle("Ramesh", "KA-01-2222")
	if err != nil {
		fmt.Println("Mapping Error: ", err)
	}
	err = bs.MapDriverToVehicle("Manjunath", "KA-01-2223")
	if err != nil {
		fmt.Println("Mapping Error: ", err)
	}

	// add trips
	err = bs.AddTrip(50, 1723601123, 1723701123, "INTRACITY", "COMPLETED", "Ramesh")
	if err != nil {
		fmt.Println("Trip Error: ", err)
	}

	err = bs.AddTrip(1050, 1723601123, 1723701123, "OUTSTATION", "COMPLETED", "Ramesh")
	if err != nil {
		fmt.Println("Trip Error: ", err)
	}

	err = bs.AddTrip(50, 1723601123, 1723701123, "INTRACITY", "CANCELED_CUSTOMER", "Ramesh")
	if err != nil {
		fmt.Println("Trip Error: ", err)
	}

	err = bs.AddTrip(50, 1723601123, 1723701123, "OUTSTATION", "CANCELED_DRIVER", "Manjunath")
	if err != nil {
		fmt.Println("Trip Error: ", err)
	}

	err = bs.AddTrip(70, 1723601123, 1723701123, "OUTSTATION", "COMPLETED", "Manjunath")
	if err != nil {
		fmt.Println("Trip Error: ", err)
	}

	// calc bill
	bill, err := bs.GetBill("Ramesh")
	if err != nil {
		fmt.Println("Billing Error: ", err)
	}
	fmt.Printf("Bill for Driver Ramesh: Rs %.2f\n", bill)

	bill, err = bs.GetBill("Manjunath")
	if err != nil {
		fmt.Println("Billing Error: ", err)
	}
	fmt.Printf("Bill for Driver Manjunath: Rs %.2f\n", bill)
}
