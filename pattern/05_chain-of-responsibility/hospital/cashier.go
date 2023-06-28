package hospital

import (
	"WB2/pattern/chain-of-responsibility/interfaces"
	"WB2/pattern/chain-of-responsibility/patient"
	"fmt"
)

type cashier struct {
	next interfaces.Department
}

func NewCashier() *cashier {
	return &cashier{}
}

func (c *cashier) Execute(p *patient.Patient) {
	if p.PaymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient")
}

func (c *cashier) SetNext(next interfaces.Department) {
	c.next = next
}
