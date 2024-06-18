package main

import "fmt"

// Somente declaração
var integer int

// Declaração e inicialização sem definir o tipo explicitamente
var integerChar = 'w'

// Declaração e inicialização de 2 variáveis em somente uma linha
var string1, string2 string = "Hello", "World"

func main() {
	// Declaração e inicialização de variável. Funciona somente dentro de funções
	float := 10.123
	boolean := false
	stringCrase := `Crase
		para
		multiplas
		linhas
	`

	fmt.Printf("float = %v \n", float)
	fmt.Printf("integer = %v \n", integer)
	fmt.Printf("string1 = %v \n", string1)
	fmt.Printf("string2 = %v \n", string2)
	fmt.Printf("boolean = %v \n", boolean)
	fmt.Printf("char = %v \n", integerChar)
	fmt.Printf("stringCrase = %v \n", stringCrase)
}
