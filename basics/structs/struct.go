package main

import "fmt"

type Car struct {
	Name  string
	Year  int
	Color string
}

func (c Car) info() string {
	return fmt.Sprintf("Name: %s\nYear: %d\nColor: %s", c.Name, c.Year, c.Color)
}

func main() {
	car1 := Car{"Corolla", 2017, "White"}
	car2 := Car{
		Name:  "BMW 320i",
		Year:  2017,
		Color: "Black",
	}

	fmt.Println(car1.Name, car1.Year, car1.Color)
	fmt.Println("-------------")
	fmt.Println(car2.info())
}
