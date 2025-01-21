package main

import "fmt"

func main(){
	channel := make(chan int)
	/*
	Utilizando o código comentado abaixo causa deadlock.
	Pelo menos uma das intruções abaixo devem estar dentro de uma
	go routine
	*/
	// channel <- 10
	// fmt.Println(<-channel)

	// Exemplo 1
	go func() {
		fmt.Println(<-channel)
	}()
	channel <- 10

	// Exemplo 2
	go func() {
		channel <- 10
	}()
	fmt.Println(<-channel)
}