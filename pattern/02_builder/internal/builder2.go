package internal

type Builder2 struct {
	component1 int
	component2 int
}

func NewBuilder2() *Builder2 {
	return &Builder2{}
}
func (b *Builder2) Prepare1() {
	b.component1 = 2
}
func (b *Builder2) Prepare2() {
	b.component2 = 2
}

func (b *Builder2) GetProduct() Product {
	return Product{component1: b.component1, component2: b.component2}
}
