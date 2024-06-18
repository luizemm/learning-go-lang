package main

import "fmt"

func main() {
	value := "Luiz"

	// Não precisa do break
	switch value {
	case "Luiz":
		fmt.Printf("Olá Luiz")
	case "Maria":
		fmt.Printf("Olá Maria")
	case "João":
		fmt.Printf("Olá João")
	default:
		fmt.Printf("Nome não encontrado")
	}
}
