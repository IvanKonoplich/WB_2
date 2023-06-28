package main

/*
	Паттерн «Фабричный метод».
Шаблон "Фабрика" - это порождающий шаблон проектирования, а также один из наиболее часто используемых шаблонов.
Этот шаблон позволяет скрыть логику создания генерируемых экземпляров.

Клиент взаимодействует только с фабричной структурой и сообщает, какие экземпляры необходимо создать.
Класс фабрики взаимодействует с соответствующими конкретными структурами и возвращает правильный экземпляр.
*/

import (
	"WB2/pattern/06_factory_method/gun"
	"WB2/pattern/06_factory_method/interfaces"
	"fmt"
	"log"
)

func main() {
	ak47, err := gun.GetGun("ak47")
	if err != nil {
		log.Fatalf("Cannot create ak47 gun. Error %v", err)
	}
	maverick, err := gun.GetGun("maverick")
	if err != nil {
		log.Fatalf("Cannot create maverick gun. Error %v", err)
	}
	printDetails(ak47)
	printDetails(maverick)
}

func printDetails(g interfaces.Gun) {
	fmt.Printf("Gun: %s\n", g.GetName())
	fmt.Printf("Power: %d\n", g.GetPower())
}
