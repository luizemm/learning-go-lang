package main

import "fmt"

func main() {
	value := 10

	if value > 10 {
		fmt.Println("maior que 10")
	} else if value < 5 {
		fmt.Println("menor que 5")
	} else {
		fmt.Println("nenhuma das condições")
	}

	boolean := true

	if boolean {
		fmt.Println("sem parentêses")
	}

	if boolean {
		fmt.Println("com parentêses")
	}

	if value2 := "Hello"; boolean {
		fmt.Println(value2)
	}

	// value2 = "Teste" - É acessível somente dentro do escopo do if

	if ifFunc(); false {
		fmt.Println("Não imprime")
	}
}

func ifFunc() {
	fmt.Println(
		"Também pode executar funções. E executa independente da condição do if",
	)
}
