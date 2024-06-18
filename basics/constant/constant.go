package main

import "fmt"

const constant1 = "Constante sem definir o tipo explicitamente"
const constant2 string = "Constante com definição de tipo"

const (
	PublicConstant  = "Constante pública"
	privateConstant = "Constante privada"
)

func main() {
	const constant = 123

	fmt.Println(constant1)
	fmt.Println(constant2)
	fmt.Println(PublicConstant)
	fmt.Println(privateConstant)
	fmt.Println(constant)
}
