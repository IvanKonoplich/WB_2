package main

import (
	"WB2/patterns/builder/internal"
	"fmt"
)

func main() {
	b1 := internal.NewBuilder1()
	b2 := internal.NewBuilder2()
	d := internal.NewDirector(b1)
	fmt.Println(d.MakeProduct())
	d.SetBuilder(b2)
	fmt.Println(d.MakeProduct())
}
