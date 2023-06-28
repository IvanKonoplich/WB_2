package visitor

import (
	"WB2/pattern/03_visitor/interfaces"
	"fmt"
)

type areaCalculator struct {
	area int
}

func NewAreaCalculator() *areaCalculator {
	return &areaCalculator{}
}

func (a *areaCalculator) VisitForSquare(s interfaces.Shape) {
	// Вычисляем площадь для квадрата. После вычисления площади присваиваем её в
	// переменную area экземпляра
	fmt.Println("Calculating area for square")
}

func (a *areaCalculator) VisitForCircle(s interfaces.Shape) {
	// Вычисляем площадь для окружности. После вычисления площади присваиваем её в
	// переменную area экземпляра
	fmt.Println("Calculating area for circle")
}

func (a *areaCalculator) VisitForTriangle(s interfaces.Shape) {
	// Вычисляем площадь для треугольника. После вычисления площади присваиваем её в
	// переменную area экземпляра
	fmt.Println("Calculating area for triangle")
}
