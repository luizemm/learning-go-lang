package main

import "fmt"

type Any []interface{}

func (a *Any) load() {
	*a = Any{
		"Luiz",
		1,
		23.5,
	}
}

func (a Any) printSlice() {
	fmt.Println(a)
}

func main() {
	var slice Any

	slice.load()
	slice.printSlice()
}
