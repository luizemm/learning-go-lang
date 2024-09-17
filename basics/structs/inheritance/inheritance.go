package main

import "fmt"

type Car struct {
	Name  string
	Year  int
	Color string
}

type SuperCar struct {
	Car
	CanFly bool
}

func main() {
	sCar := SuperCar{
		Car: Car{
			Name:  "BMW 320i",
			Year:  2017,
			Color: "Black",
		},
		CanFly: true,
	}

	fmt.Println(sCar.Name)
	fmt.Println(sCar.Car.Name)
}
