package main

import (
	"fmt"

	"github.com/luizemm/learning-go-lang/basics/visibility"
)

func main() {
	fmt.Println(visibility.PublicVariable)
	visibility.PublicMethod()
}
