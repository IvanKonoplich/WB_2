package main

import (
	"WB2/pattern/visitor/shapes"
	"WB2/pattern/visitor/visitor"
	"fmt"
)

func main() {
	square := shapes.NewSquare(2)
	circle := shapes.NewCircle(3)
	triangle := shapes.NewTriangle(1, 2, 3)

	areaCalculator := visitor.NewAreaCalculator()
	square.Accept(areaCalculator)
	circle.Accept(areaCalculator)
	triangle.Accept(areaCalculator)

	fmt.Println()
	middleCoordinates := visitor.NewMiddleCoordinates()
	square.Accept(middleCoordinates)
	circle.Accept(middleCoordinates)
	triangle.Accept(middleCoordinates)
}
