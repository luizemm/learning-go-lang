package visibility

import "fmt"

/*
Variáveis, métodos e etc que iniciam com letra minúscula são privados
e que iniciam com letra maiúscula são publicas
*/

var privateVariable string = "Variável privada"
var PublicVariable string = "Variável pública"

func privateMethod() {
	fmt.Printf("Métodos privados iniciam com letra minúscula")
}

func PublicMethod() {
	fmt.Printf("Métodos públicos iniciam com letra Maiúscula")
}
