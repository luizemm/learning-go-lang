package main

import (
	"fmt"
	"strconv"
)

func main() {
	// O length define o tamanho inicial do slice
	length := 2
	/*
		A capacity define o tamanho máximo que o slice pode alcançar
		até que seja necessário criar um novo array internamente
		com 2x vezes o tamanho máximo do slice anterior
	*/
	capacity := 5

	slice := make([]int, length, capacity)
	fmt.Println(slice)
	fmt.Println(getSliceInfo(slice))

	slice[1] = 1
	// slice[2] = 1 error: index out of range

	slice = append(slice, 2, 3, 4, 5)
	fmt.Println(slice)
	fmt.Println(getSliceInfo(slice))

	fmt.Println("--------------------")

	for i := 6; i < 23; i++ {
		slice = append(slice, i)
		fmt.Println(getSliceInfo(slice))
	}

	fmt.Println("--------------------")

	sliceA := make([]int, 2, 5)
	sliceB := sliceA

	fmt.Println("sliceA:", sliceA)
	fmt.Println("sliceB:", sliceB)

	// Altera também o sliceB, por estar apontando para o mesmo slice
	sliceA[0] = 10

	fmt.Println("sliceA:", sliceA, getSliceInfo(sliceA))
	fmt.Println("sliceB:", sliceB, getSliceInfo(sliceB))

	/*
		Porém ao exceder a capacidade do slice é criado um novo.
		Logo, o sliceB não irá mais apontar para o sliceA
	*/

	sliceB = append(sliceB, 1, 2, 3, 4, 5)

	sliceA[1] = 20

	fmt.Println("sliceA:", sliceA, getSliceInfo(sliceA))
	fmt.Println("sliceB:", sliceB, getSliceInfo(sliceB))

	fmt.Println("--------------------")

	sliceString := []string{
		"Hello",
		"World",
		"Much",
		"Better",
	}

	// Os números abaixo são quantidade e não o índice
	fmt.Println(sliceString[:2])
	fmt.Println(sliceString[1:2])
	fmt.Println(sliceString[2:4])
	fmt.Println(sliceString[1:])
}

func getSliceInfo(slice []int) string {
	return "Len " + strconv.Itoa(len(slice)) + " Cap " + strconv.Itoa(cap(slice))
}
