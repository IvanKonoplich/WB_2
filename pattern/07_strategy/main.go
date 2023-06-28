package main

/*
	Паттерн «Стратегия».
Шаблон "Стратегия" - это поведенческий шаблон проектирования. Он позволяет изменять поведение объекта
во время выполнения программы без каких-либо изменений в классе этого объекта.

Стоит использовать когда объекту необходимо поддерживать различное поведение, и вы хотите изменить поведение во время выполнения.
Когда вы хотите избежать множества условий при выборе поведения во время выполнения.
Когда существуют различные похожие друг на друга алгоритмы и они отличаются только какой-то определенной частью.
*/
import "WB2/pattern/07_strategy/cache"

func main() {
	lfu := cache.NewLfu()
	cache1 := cache.InitCache(lfu)
	cache1.Add("a", "dev01")
	cache1.Add("b", "dev02")
	cache1.Add("c", "dev03")
	lru := cache.NewLru()
	cache1.SetEvictionAlgo(lru)
	cache1.Add("d", "dev04")
	fifo := cache.NewFifo()
	cache1.SetEvictionAlgo(fifo)
	cache1.Add("e", "dev05")
}
