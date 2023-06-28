package interfaces

import "WB2/pattern/05_chain-of-responsibility/patient"

type Department interface {
	Execute(*patient.Patient)
	SetNext(Department)
}
