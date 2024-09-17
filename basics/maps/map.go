package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["a"] = 0
	m["b"] = 20
	m["c"] = 30
	fmt.Println(m)

	delete(m, "c")
	fmt.Println("map[c]:", m["c"])

	_, exists := m["c"]
	fmt.Println("exists key c: ", exists)

	value, exists := m["a"]
	fmt.Println("exists key a: ", exists)
	fmt.Println("map[a]:", value)

	newMap := map[string]int{"x": 1, "y": 2}
	fmt.Println(newMap)

	if value, exists := m["a"]; exists {
		fmt.Println("map[a]:", value)
	} else {
		fmt.Println("Not exists value")
	}
}
