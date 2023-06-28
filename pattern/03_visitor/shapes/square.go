package shapes

import "WB2/pattern/03_visitor/interfaces"

type square struct {
	side int
}

func NewSquare(side int) *square {
	return &square{
		side: side,
	}
}

func (s *square) Accept(v interfaces.Visitor) {
	v.VisitForSquare(s)
}

func (s *square) GetType() string {
	return "Square"
}
