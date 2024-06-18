package main

import "fmt"

func main() {
	value1 := 10
	pointer1 := &value1

	// Imprime o endereço da memória
	fmt.Println("Endereços:")
	fmt.Println(&value1)
	fmt.Println(pointer1)

	// Imprime o valor
	fmt.Println("Valores:")
	fmt.Println(value1)
	fmt.Println(*pointer1)

	// Alterando o valor no endereço da memória da variável value
	*pointer1 = 20
	fmt.Println("Valores alterados:")
	fmt.Println(value1)
	fmt.Println(*pointer1)

	// Não altera o valor da variável value
	fmt.Println("Valores diferentes:")
	fmt.Println(paramByValue(value1))
	fmt.Println(value1)

	// Não altera o valor da variável value
	fmt.Println("Valores iguais:")
	fmt.Println(paramByReference(&value1))
	fmt.Println(value1)
}

func paramByValue(param int) int {
	param += param

	return param
}

func paramByReference(param *int) int {
	*param += *param

	return *param
}
