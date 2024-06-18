package main

import "fmt"

/*
Declarações em outro arquivo mas que possuam o mesmo nome de pacote são enxergadas
em ambos.
Obs.:

	É necessário incluir os dois arquivos ao executar.
	Exemplo: go run scope/*.go (Executando a partir da raiz do projeto)
*/

func main() {
	fmt.Println(otherFilesVariable)
	printOtherFilesVariable()
}
