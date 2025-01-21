package main

import "fmt"

func main() {
	c := generate(5, 10)

	d1 := divide(c)

	for t := range d1 {
		fmt.Println("d1: ", t)
	}

	d2 := divide(c)

	for t := range d2 {
		fmt.Println("d2: ", t)
	}
}

func generate(numbers ...int) chan int {
	channel := make(chan int)
	go func() {
		for _, n := range numbers {
			channel <- n
		}
		close(channel)
	}()

	return channel
}

func divide(input chan int) chan int {
	channel := make(chan int)

	go func() {
		channel <- <- input / 2
		close(channel)
	}()

	return channel
}