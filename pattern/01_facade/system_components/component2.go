package systemComponents

import "fmt"

type component2 struct {
}

func NewC2() *component1 {
	return &component1{}
}

func (c *component1) Operation2() {
	fmt.Println("operation dev02 running")
}
