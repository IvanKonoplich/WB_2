package main

import (
	"fmt"
	"sync"
	"time"
)

//Реализовать функцию, которая будет объединять один или более done-каналов
//в single-канал, если один из его составляющих каналов закроется.
//Очевидным вариантом решения могло бы стать выражение
//при использованием select, которое бы реализовывало эту связь,
//однако иногда неизвестно общее число done-каналов,
//с которыми вы работаете в рантайме. В этом случае удобнее использовать вызов
//единственной функции, которая, приняв на вход один или более or-каналов, реализовывала бы весь функционал.
//
//Определение функции:
//var or func(channels ...<- chan interface{}) <- chan interface{}
//Пример использования функции:
//sig := func(after time.Duration) <- chan interface{} {
//	c := make(chan interface{})
//	go func() {
//		defer close(c)
//		time.Sleep(after)
//}()
//return c
//}
//
//start := time.Now()
//<-or (
//	sig(dev02*time.Hour),
//	sig(dev05*time.Minute),
//	sig(dev01*time.Second),
//	sig(dev01*time.Hour),
//	sig(dev01*time.Minute),
//)
//
//fmt.Printf(“fone after %v”, time.Since(start))

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	fmt.Println(<-or(
		sig(1*time.Second),
		sig(2*time.Second),
		sig(3*time.Second),
		sig(4*time.Second),
	))

	fmt.Printf("fone after %v", time.Since(start))
}

func or(channels ...<-chan interface{}) <-chan interface{} {
	resultChan := make(chan interface{}) //создаем канал в который будут писать все остальные каналы
	var wg sync.WaitGroup                //используем waitgroup чтобы гарантировано отработали все горутины
	for _, j := range channels {         //для каждого канала создаем горутину которая будет его слушать
		wg.Add(1)
		go func(input <-chan interface{}) {
			for val := range input { //слушаем канал, важно то, что если канал закроется и в него не будут переданы никакие данные то код в цикле не выполнится
				resultChan <- val
			}
			close(resultChan) //закрываем канал даже в том случае если не получили значение из канала
			wg.Done()
		}(j)
	}
	go func() { wg.Wait() }() //ждем все горутины в отдельной горутине, так как мы хотим сразу вернуть созданные объединяющий канал в вызывающую функцию
	return resultChan         //сразу возвращаем созданный канал
}
