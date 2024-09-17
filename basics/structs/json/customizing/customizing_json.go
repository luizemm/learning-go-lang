package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Name  string `json:"Model"` //Campo "Name" imprime como "Model" no json
	Year  int    `json:"-"`     //Não adiciona o campo "Year" no json
	Color string
}

func main() {
	car := Car{"Gol", 2017, "White"}

	result, _ := json.Marshal(car)

	fmt.Println(string(result))
}
