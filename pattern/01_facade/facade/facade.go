package facade

import (
	systemComponents2 "WB2/pattern/01_facade/system_components"
)

type facade struct {
	cmp1 component1
	cmp2 component2
	cmp3 component3
}

type component1 interface {
	Operation1()
}

type component2 interface {
	Operation2()
}

type component3 interface {
	Operation3()
}

func NewFacade() *facade {
	return &facade{
		cmp1: systemComponents2.NewC1(),
		cmp2: systemComponents2.NewC2(),
		cmp3: systemComponents2.NewC3(),
	}
}

func (f *facade) Op1() {
	f.cmp1.Operation1()
}
func (f *facade) Op2() {
	f.cmp2.Operation2()
}
func (f *facade) Op3() {
	f.cmp3.Operation3()
}
