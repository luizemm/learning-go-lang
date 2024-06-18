package main

import (
	"fmt"

	"github.com/luizemm/learning-go-lang/visibility"
)

func main() {
	fmt.Println(visibility.PublicVariable)
	visibility.PublicMethod()
}
