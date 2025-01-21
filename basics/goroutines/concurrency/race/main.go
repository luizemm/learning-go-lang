package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup
var result int

/* 
Comando: go run -race main.go

A opção "-race" verifica se está ocorrendo
alguma condição de corrida no código

*/

func main() {
	waitGroup.Add(2)

	go runProcess("P1", 20)
	go runProcess("P2", 20)

	waitGroup.Wait()

	fmt.Println("Final result:", result)
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		z := result
		z++
		fmt.Println(name, "->", i)
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
		result = z
		fmt.Println(name, "->", "Partial result:", result)
	}
	waitGroup.Done() 
}