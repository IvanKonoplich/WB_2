package systemComponents

import "fmt"

type component1 struct {
}

func NewC1() *component1 {
	return &component1{}
}

func (c *component1) Operation1() {
	fmt.Println("operation 1 running")
}
