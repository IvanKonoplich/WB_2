package internal

type Builder interface {
	Prepare1()
	Prepare2()
	GetProduct() Product
}
