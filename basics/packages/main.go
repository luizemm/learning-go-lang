package main

import (
	"fmt"

	"github.com/luizemm/learning-go-lang/basics/packages/car"
)

func main() {
	car := car.Car{Name: "Gol", Color: "White"}

	fmt.Println(car.Start())
}
