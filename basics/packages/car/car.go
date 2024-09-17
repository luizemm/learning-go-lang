package car

import "fmt"

type Car struct {
	Name  string
	Color string
}

func (c Car) Start() string {
	return c.Name + " has been started"
}

func (c Car) nonExportedMethod() {
	fmt.Println("Method is not exported")
}
