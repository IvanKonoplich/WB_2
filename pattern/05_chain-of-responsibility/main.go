package main

/*
	Паттерн «Цепочка Обязанностей».
Цепочка Обязанностей — это поведенческий шаблон проектирования.
Он позволяет создать цепочку обработчиков запроса. Каждый входящий запрос передаётся по цепочке и каждый обработчик:
Обрабатывает запрос или нет.
Решает передавать ли запрос следующему обработчику в цепочке или нет.

Шаблон применим, когда существуют разные способы обработки одного и того же запроса.
А также когда не нужно, чтобы клиент выбирал обработчика, поскольку несколько обработчиков
могут обрабатывать запрос. Кроме того, нужно отделить клиента от обработчиков. Клиенту нужно знать только первый
элемент в цепочке.
*/
import (
	"WB2/pattern/05_chain-of-responsibility/hospital"
	"WB2/pattern/05_chain-of-responsibility/patient"
)

func main() {
	cashier := hospital.NewCashier()
	medical := hospital.NewMedical()
	medical.SetNext(cashier)
	doctor := hospital.NewDoctor()
	doctor.SetNext(medical)
	reception := hospital.NewReception()
	reception.SetNext(doctor)
	hospitalPatient := patient.NewPatient("abc")
	reception.Execute(hospitalPatient)
}
