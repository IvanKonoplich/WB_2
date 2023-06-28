package main

/*
	Паттерн «строитель».
Шаблон "Строитель" - это порождающий шаблон проектирования, используемый для создания сложных объектов.
Шаблон "Строитель" подходит, когда создаваемый объект очень большой и состоит из нескольких стадий.
Когда необходимо создать другую версию того же объекта.
Когда не может существовать частично инициализированного объекта.
*/

import (
	"WB2/pattern/02_builder/internal"
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
