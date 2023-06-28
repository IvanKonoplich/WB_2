package internal

type Director struct {
	Builder Builder
}

func NewDirector(b Builder) *Director {
	return &Director{b}
}

func (d *Director) SetBuilder(b Builder) {
	d.Builder = b
}

func (d *Director) MakeProduct() Product {
	d.Builder.Prepare1()
	d.Builder.Prepare2()
	return d.Builder.GetProduct()
}
