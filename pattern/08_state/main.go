package main

/*
	Паттерн «состояние».
Шаблон "Состояние" - это поведенческий шаблон проектирования, основанный на конечном состоянии.
Используется, когда объект может быть в различных состояниях. В зависимости от текущего запроса
объекту необходимо изменять свое текущее состояние. Также когда объект может по-разному отвечать
на один и тот же запрос в зависимости от текущего состояния.
В этом случае использование шаблона проектирования "Состояние" заменит множество условных операторов.
*/
import (
	"WB2/pattern/08_state/machine"
	"fmt"
	"log"
)

func main() {
	vendingMachine := machine.NewVendingMachine(1, 10)
	err := vendingMachine.RequestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}
	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()
	err = vendingMachine.AddItem(2)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println()

	err = vendingMachine.RequestItem()
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatalf(err.Error())
	}
}
