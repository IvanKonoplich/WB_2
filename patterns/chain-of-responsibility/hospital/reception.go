package hospital

import (
	"WB2/patterns/chain-of-responsibility/interfaces"
	"WB2/patterns/chain-of-responsibility/patient"
	"fmt"
)

type reception struct {
	next interfaces.Department
}

func NewReception() *reception {
	return &reception{}
}

func (r *reception) Execute(p *patient.Patient) {
	if p.RegistrationDone {
		fmt.Println("Patient registration already done")
		r.next.Execute(p)
		return
	}
	fmt.Println("Reception registering patient")
	p.RegistrationDone = true
	r.next.Execute(p)
}

func (r *reception) SetNext(next interfaces.Department) {
	r.next = next
}
