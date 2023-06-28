package main

import (
	"WB2/patterns/facade/facade"
)

func main() {
	fc := facade.NewFacade()
	fc.Op1()
	fc.Op2()
	fc.Op3()
}
