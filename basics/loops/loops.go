package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("For comum: %v\n", i)
	}

	cont := 0

	for cont < 5 {
		fmt.Printf("while (no go usa o for também): %v\n", cont)
		cont++
	}

	cont = 0
	for {
		cont++
		fmt.Println("For infinito")
		if cont < 5 {
			continue
		}
		break
	}

	list := [3]string{"a", "b", "c"}

	for index, value := range list {
		fmt.Printf("list[%v] = %v\n", index, value)
	}
}
