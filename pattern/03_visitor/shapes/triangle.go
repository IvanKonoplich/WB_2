package shapes

import "WB2/pattern/03_visitor/interfaces"

type Triangle struct {
	a int
	b int
	c int
}

func NewTriangle(a, b, c int) *Triangle {
	return &Triangle{
		a: a,
		b: b,
		c: c,
	}
}

func (t *Triangle) Accept(v interfaces.Visitor) {
	v.VisitForTriangle(t)
}

func (t *Triangle) GetType() string {
	return "Triangle"
}
