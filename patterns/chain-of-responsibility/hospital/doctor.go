package hospital

import (
	"WB2/patterns/chain-of-responsibility/interfaces"
	"WB2/patterns/chain-of-responsibility/patient"
	"fmt"
)

type doctor struct {
	next interfaces.Department
}

func NewDoctor() *doctor {
	return &doctor{}
}

func (d *doctor) Execute(p *patient.Patient) {
	if p.DoctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.Execute(p)
		return
	}
	fmt.Println("Doctor checking patient")
	p.DoctorCheckUpDone = true
	d.next.Execute(p)
}

func (d *doctor) SetNext(next interfaces.Department) {
	d.next = next
}
