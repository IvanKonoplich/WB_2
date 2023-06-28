package internal

type Builder1 struct {
	component1 int
	component2 int
}

func NewBuilder1() *Builder1 {
	return &Builder1{}
}
func (b *Builder1) Prepare1() {
	b.component1 = 1
}
func (b *Builder1) Prepare2() {
	b.component2 = 1
}

func (b *Builder1) GetProduct() Product {
	return Product{component1: b.component1, component2: b.component2}
}
