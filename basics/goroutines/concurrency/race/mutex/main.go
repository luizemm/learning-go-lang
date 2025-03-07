package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup
var result int
var m sync.Mutex

func main() {
	waitGroup.Add(2)

	go runProcess("P1", 20)
	go runProcess("P2", 20)

	waitGroup.Wait()

	fmt.Println("Final result:", result)
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
		m.Lock()
		result++
		fmt.Println(name, "->", "Partial result:", result)
		m.Unlock()
	}
	waitGroup.Done() 
}