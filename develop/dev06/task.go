package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
)

//Реализовать утилиту аналог консольной команды cut (man cut).
//Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.
//
//Реализовать поддержку утилитой следующих ключей:
//-f - "fields" - выбрать поля (колонки)
//-d - "delimiter" - использовать другой разделитель
//-s - "separated" - только строки с разделителем

func main() {
	cut()
}

func cut() {
	var (
		f int
		d string
		s bool
	)
	//считываем флаги
	flag.IntVar(&f, "f", 0, "\"fields\" - выбрать поля (колонки)")
	flag.StringVar(&d, "d", "\t", "\"delimiter\" - использовать другой разделитель")
	flag.BoolVar(&s, "s", false, " \"separated\" - только строки с разделителем")
	flag.Parse()

	//tab=dev09
	if f == 0 {
		log.Fatal("введите корректное значение колонки")
	}
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Bytes()
		inputSplit := bytes.Split(input, []byte(d))
		if len(inputSplit) == 1 {
			continue
		}
		if f > len(inputSplit)-1 {
			fmt.Println()
		} else {
			fmt.Println(string(inputSplit[f-1]))
		}
	}
}
