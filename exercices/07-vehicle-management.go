package main

import (
	"fmt"
)

type Insurable interface {
	CalculateInsurance() float64
}

type Printable interface {
	PrintDetails()
}

type Vehicle struct {
	Make  string
	Model string
	Year  int
}

func (v *Vehicle) PrintDetails() {
	fmt.Printf("%d %s %s\n", v.Year, v.Make, v.Model)
}

type Car struct {
	Vehicle
	NumberOfDoors int
}

func (v *Car) PrintDetails() {
	v.Vehicle.PrintDetails()
	fmt.Printf("Number of doors: %d\n", v.NumberOfDoors)
}

func (v *Car) CalculateInsurance() float64 {
	return 100.0
}

type Truck struct {
	Vehicle
	PayloadCapacity float64
}

func (v *Truck) PrintDetails() {
	v.Vehicle.PrintDetails()
	fmt.Printf("Payload capacity: %f\n", v.PayloadCapacity)
}

func (v *Truck) CalculateInsurance() float64 {
	return 200.0
}

func main() {
	vehicles := Code()

	for _, v := range vehicles {
		v.PrintDetails()
	}
}

func Code() []Printable {
	var vehicles []Printable

	vehicles = append(vehicles, &Car{
		Vehicle: Vehicle{
			Make:  "Toyota",
			Model: "Corolla",
			Year:  2022,
		},
		NumberOfDoors: 8,
	})

	vehicles = append(vehicles, &Truck{
		Vehicle: Vehicle{
			Make:  "Iveco",
			Model: "apah",
			Year:  2022,
		},
		PayloadCapacity: 1000.0,
	})

	vehicles = append(vehicles, &Truck{
		Vehicle: Vehicle{
			Make:  "Ford",
			Model: "F150",
			Year:  2018,
		},
		PayloadCapacity: 7000.0,
	})

	vehicles = append(vehicles, &Truck{
		Vehicle: Vehicle{
			Make:  "MITSUBISHI",
			Model: "L200",
			Year:  2018,
		},
		PayloadCapacity: 5000.0,
	})

	vehicles = append(vehicles, &Car{
		Vehicle: Vehicle{
			Make:  "Dacia",
			Model: "Logan",
			Year:  2024,
		},
		NumberOfDoors: 4,
	})

	vehicles = append(vehicles, &Car{
		Vehicle: Vehicle{
			Make:  "Bugatti",
			Model: "Veyron",
			Year:  2024,
		},
		NumberOfDoors: 4,
	})

	return vehicles
}
