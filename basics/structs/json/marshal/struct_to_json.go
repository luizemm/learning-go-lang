package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Name  string
	Year  int
	color string
}

func main() {
	car := Car{"Gol", 2017, "White"}

	result, _ := json.Marshal(car)

	// Irá imprimir somente o nome e o ano, pois a cor é privada (começa com letra minúscula)

	// Imprime json como array de byte
	fmt.Println(result)

	// Imprime json como string
	fmt.Println(string(result))
}
