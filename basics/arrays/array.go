package main

import "fmt"

func main() {
	// Inicializa os valores com 0(zero)
	var intArray [10]int
	fmt.Println(intArray)
	fmt.Println(len(intArray))

	intArray[0] = 10
	intArray[1] = 6
	intArray[2] = 12

	fmt.Println(intArray)

	// Inicializa os valores vazios
	var stringArray [10]string
	fmt.Println(stringArray)

	int2Array := [5]int{2, 4, 5, 3, 1}
	fmt.Println(int2Array)
}
