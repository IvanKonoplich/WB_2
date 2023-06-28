package systemComponents

import "fmt"

type component3 struct {
}

func NewC3() *component1 {
	return &component1{}
}

func (c *component1) Operation3() {
	fmt.Println("operation dev03 running")
}
