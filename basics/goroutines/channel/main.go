package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	runSimpleExample()
	runComplexExample()
}

func runSimpleExample() {
	msg := make(chan string)

	go func() {
		msg <- "Hello World"
	}()

	result := <- msg
	fmt.Println(result)
}

func runComplexExample() {
	channel := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			// adiciona o valor de i no canal, e depois fica aguardando até o canal ser esvaziado
			channel <- i
		}
	}()

	go func() {
		for {
			// esvazia o canal obtendo o valor de i, e depois fica aguardando até preencher o canal novamente
			fmt.Println(<- channel)
		}
	}()

	time.Sleep(time.Second)
}

func runRangeCloseExample() {
	channel := make(chan int)
	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
		waitGroup.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
		waitGroup.Done()
	}()

	go func() {
		waitGroup.Wait()
		close(channel)
	}()

	for number := range channel {
		fmt.Println(number)
	}
}