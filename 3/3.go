package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// Отсортировать строки в файле по аналогии с консольной утилитой sort (man sort — смотрим описание и основные параметры):
// на входе подается файл из несортированных строками, на выходе — файл с отсортированными.
//
// Реализовать поддержку утилитой следующих ключей:
//
// -k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
// -n — сортировать по числовому значению
// -r — сортировать в обратном порядке
// -u — не выводить повторяющиеся строки
//
// # Дополнительно
//
// Реализовать поддержку утилитой следующих ключей:
//
// -M — сортировать по названию месяца
// -b — игнорировать хвостовые пробелы
// -c — проверять отсортированы ли данные
// -h — сортировать по числовому значению с учетом суффиксов
func main() {
	//вызываем функцию сортировки
	manSort()
}

func manSort() {
	//создаем переменные для флагов
	var (
		fn string
		k  int
		n  bool
		r  bool
		u  bool
	)
	//считываем флаги
	flag.StringVar(&fn, "fn", "input1.txt", "имя файла")
	flag.IntVar(&k, "k", 0, "указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)")
	flag.BoolVar(&n, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&r, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&u, "u", false, "не выводить повторяющиеся строки")
	flag.Parse()

	//читаем данные из файла
	data, err := os.ReadFile("3/" + fn)
	if err != nil {
		log.Fatal(err)
	}
	//если указан флаг k то, сразу делим слайс байт на строки и перемещаем указанную колонку вперед
	if k != 0 {
		dataSplit := bytes.Split(data, []byte("\n"))
		for _, j := range dataSplit {
			j[0], j[(k-1)*2] = j[(k-1)*2], j[0] //используем такую запись потому что в слайсе байты не только значений, но и пробелов, соответственно их вдвое больше
		}
		data = bytes.Join(dataSplit, []byte("\n"))
	}
	fmt.Println(data)
	fmt.Println("----------------------------------------------")

	var str []string
	//сортировка по числовому значению
	if n {
		//отделяем значения начинающиеся с чисел от буквенных
		digits := make([][]byte, 0, len(data))
		words := make([]string, 0, len(data))
		dataSplit := bytes.Split(data, []byte("\n"))
		for _, j := range dataSplit {
			if unicode.IsDigit(rune(j[0])) {
				digits = append(digits, j)
			} else {
				words = append(words, string(j))
			}
		}
		//сортируем два слайса отдельно и совмещаем
		sort.Strings(words)
		digits = sortDigits(digits)
		str = append(words, bytesToStrings(digits)...)
	} else {
		//создаем слайс строк для сортировки
		str = strings.Split(string(data), "\n")
		sort.Strings(str)
	}
	//сортировка в обратном порядке
	if r {
		for i := 0; i < len(str)/2; i++ {
			str[i], str[len(str)-1-i] = str[len(str)-1-i], str[i]
		}
	}
	//не выводить повторяющиеся строки
	if u {
		str = removeDuplicateStr(str)
	}

	//выводим результат
	var result string
	for i, j := range str {
		if i == len(str)-1 {
			result += j
		} else {
			result += j + "\n"
		}
	}
	//меняем колонки обратно
	if k != 0 {
		data := []byte(result)
		dataSplit := bytes.Split(data, []byte("\n"))
		for _, j := range dataSplit {
			j[0], j[(k-1)*2] = j[(k-1)*2], j[0]
		}
		result = string(bytes.Join(dataSplit, []byte("\n")))
	}
	fmt.Println(result)
}

// сортирует строки на основе числа в начале
func sortDigits(inc [][]byte) [][]byte {
	if len(inc) < 2 { //если один элемент, то его и возвращаем
		return inc
	}
	pivot := findDigit(inc[0])    //отпределяем от чего отталкиваемся
	var less, greater [][]byte    //создаем слайсы для больших и меньших элементов относительно опоры
	for _, num := range inc[1:] { //распределяем значения
		if findDigit(num) < pivot {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}
	result := append(sortDigits(less), inc[0]) //рекурсивно сортируем полученные слайсы
	result = append(result, sortDigits(greater)...)
	return result

}

// находит число в начале строки, если например 123kwfbwe, чтобы отсортировать с флагом n
func findDigit(inc []byte) int {
	var result string
	for _, j := range inc {
		if unicode.IsDigit(rune(j)) {
			result += string(j)
		} else {
			break
		}
	}
	resultInt, _ := strconv.Atoi(result)
	return resultInt
}

// Преображает байты в строки
func bytesToStrings(inc [][]byte) []string {
	result := make([]string, 0, len(inc))
	for _, j := range inc {
		result = append(result, string(j))
	}
	return result
}

// удаляет дубликаты
func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
