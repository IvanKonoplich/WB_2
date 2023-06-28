package interfaces

import "WB2/patterns/chain-of-responsibility/patient"

type Department interface {
	Execute(*patient.Patient)
	SetNext(Department)
}
