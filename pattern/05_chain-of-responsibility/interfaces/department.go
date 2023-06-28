package interfaces

import "WB2/pattern/chain-of-responsibility/patient"

type Department interface {
	Execute(*patient.Patient)
	SetNext(Department)
}
