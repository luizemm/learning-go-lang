package main

import "fmt"

func namedFunc(param string) (result string) {
	result = "Resultado: " + param
	return
}

func moreReturns(numA int, numB int) (string, string, int) {
	var stringA, stringB string

	for i := 0; i < numA; i++ {
		stringA += "A"
	}

	for i := 0; i < numB; i++ {
		stringB += "B"
	}

	return stringA, stringB, numA + numB
}

func variadicFunc(nums ...int) int {
	var result int

	for _, value := range nums {
		result += value
	}

	return result
}

func funcInsideFunc(function func() string) func() {
	return func() {
		fmt.Printf("Função que retorna outra: %v", function())
	}
}

func main() {
	fmt.Println(namedFunc("Função com uma variável já definida como sendo o retorno"))

	stringA, stringB, sum := moreReturns(3, 2)
	fmt.Printf("%v %v -> %v\n", stringA, stringB, sum)

	fmt.Printf("Soma: %v\n", variadicFunc(2, 4, 2))

	anonymousFunc := func() string {
		return "Função anônima"
	}

	printResult := funcInsideFunc(anonymousFunc)

	printResult()
}
