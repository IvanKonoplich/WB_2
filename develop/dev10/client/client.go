package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

//Утилита telnet
//
//Реализовать простейший telnet-клиент.
//
//Примеры вызовов:
//go-telnet --timeout=10s host port
//go-telnet mysite.ru 8080
//go-telnet --timeout=3s dev01.dev01.dev01.dev01 123
//
//
//Требования:
//Программа должна подключаться к указанному хосту (ip или доменное имя + порт) по протоколу TCP.
//После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
//Опционально в программу можно передать таймаут на подключение к серверу
//(через аргумент --timeout, по умолчанию 10s)
//При нажатии Ctrl+D программа должна закрывать сокет и завершаться.
//Если сокет закрывается со стороны сервера, программа должна также завершаться.
//При подключении к несуществующему сервер, программа должна завершаться через timeout

func main() {
	var timeout time.Duration                                        //определяем переменную для таймаута
	flag.DurationVar(&timeout, "timeout", 10*time.Second, "timeout") //получаем таймаут
	flag.Parse()
	if len(flag.Args()) < 2 {
		panic("передано недостаточное количество аргументов")
	}
	conn, err := net.DialTimeout("tcp", flag.Args()[0]+":"+flag.Args()[1], timeout) //подключаемся к серверу с переданным таймаутом
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			if _, err := io.Copy(os.Stdout, conn); err != nil { //копируем данные из сокета в stdout
				log.Fatal(err)
			}
		}
	}()

	// stdin -> socket
	go func() {
		for {
			if _, err := io.Copy(conn, os.Stdin); err != nil { //копируем данные из stdin в сокет
				log.Fatal(err)
			}
		}
	}()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan)
	for range sigChan {
		os.Exit(0)
	}
}
