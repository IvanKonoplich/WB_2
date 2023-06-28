package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// Реализовать утилиту фильтрации по аналогии с консольной утилитой (man grep — смотрим описание и основные параметры).
//
// Реализовать поддержку утилитой следующих ключей:
// -A - "after" печатать +N строк после совпадения
// -B - "before" печатать +N строк до совпадения
// -C - "context" (A+B) печатать ±N строк вокруг совпадения
// -c - "count" (количество строк)
// -i - "ignore-case" (игнорировать регистр)
// -v - "invert" (вместо совпадения, исключать)
// -F - "fixed", точное совпадение со строкой, не паттерн
// -n - "line num", напечатать номер строки
func main() {
	grep()
}

func grep() {
	var (
		fn  string
		exp string
		A   int
		B   int
		C   int
		c   bool
		i   bool
		v   bool
		F   bool
		n   bool
	)
	//считываем флаги
	flag.StringVar(&fn, "fn", "input1.txt", "имя файла")
	flag.StringVar(&exp, "exp", "", "шаблон")
	flag.IntVar(&A, "A", 0, "\"after\" печатать +N строк после совпадения")
	flag.IntVar(&B, "B", 0, " \"before\" печатать +N строк до совпадения")
	flag.IntVar(&C, "C", 0, " \"context\" (A+B) печатать ±N строк вокруг совпадения")
	flag.BoolVar(&c, "c", false, " \"count\" (количество строк)")
	flag.BoolVar(&i, "i", false, "\"ignore-case\" (игнорировать регистр)")
	flag.BoolVar(&v, "v", false, "\"invert\" (вместо совпадения, исключать)")
	flag.BoolVar(&F, "F", false, "\"fixed\", точное совпадение со строкой, не паттерн")
	flag.BoolVar(&n, "n", false, " \"line num\", напечатать номер строки")

	flag.Parse()

	if exp == "" {
		log.Fatal("шаблон не введен")
	}

	//читаем данные из файла
	data, err := os.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}
	dataSplit := bytes.Split(data, []byte("\n"))

	//сравнивать будем по dataActual, а записывать в результат по dataOriginal. Это необходимо если указывается флаг i (игнорирование регистра)
	dataActual := make([][]byte, 0, len(dataSplit))
	dataOriginal := make([][]byte, 0, len(dataSplit))
	//если регистр не важен, то создаем копию и приводим шаблон к нижнему регистру
	if i {
		exp = strings.ToLower(exp)
		dataOriginal = dataSplit
		for _, j := range dataSplit { //если регистр не важен, то приводим dataActual по которой сравниваем к нижнему регистру
			dataActual = append(dataActual, []byte(strings.ToLower(string(j))))
		}
	} else { //если регистр важен, то записываем и сравниваем одно и то же
		dataActual = dataSplit
		dataOriginal = dataSplit
	}

	//создаем переменную для результата. Хранить будем индексы, причем слайс заранее заполним nil для простоты реализации флагов A, B и C
	result := make([]int, len(dataSplit))

	var zeroNil bool //переменная нужна, чтобы проверить является ли первое значение в слайсе nil значением
	//начинаем проверку
	for i, j := range dataActual {
		if F { //точное совпадение
			if v { //исключение или нет
				if exp != string(j) {
					result[i] = i //индексы в слайсе-результате и изначальной выборке совпадают
					if i == 0 {   //проверяем первое значение на nil
						zeroNil = true
					}
				}

			} else {
				if exp == string(j) {
					result[i] = i
					if i == 0 {
						zeroNil = true
					}
				}
			}
		} else { //наличие подстроки
			if v { //исключение или нет
				if !checker(j, []byte(exp)) {
					result[i] = i
					if i == 0 {
						zeroNil = true
					}
				}
			} else {
				if checker(j, []byte(exp)) {
					result[i] = i
					if i == 0 {
						zeroNil = true
					}
				}
			}
		}
	}
	if A != 0 || B != 0 || C != 0 { //необходимо скопировать слайс с индексами после фильтрации. Чтобы при указании нескольких флагов (например A и B) не копировать несколько раз делаем доп проверку
		resultCopy := make([]int, len(dataSplit))
		copy(resultCopy, result)
		if A != 0 {
			for i, j := range resultCopy {
				if j != 0 || (zeroNil && i == 0) { //если индекс добавлен в слайс
					if j+A > len(result)-1 {
						for i := j + 1; i <= len(result)-1; i++ {
							result[i] = i //добавляем дополнительные значение в результат
						}
					} else {
						for i := j + 1; i <= j+A; i++ {
							result[i] = i
						}
					}
				}
			}
		}
		if B != 0 {
			for i, j := range resultCopy {
				if j != 0 || (zeroNil && i == 0) { //если индекс добавлен в слайс
					if j-B < 0 {
						for i := j - 1; i >= 0; i-- {
							result[i] = i
						}
					} else {
						for i := j - 1; i >= j-B; i-- {
							result[i] = i
						}
					}
				}
			}
		}

		if C != 0 {
			for i, j := range resultCopy {
				if j != 0 || (zeroNil && i == 0) { //если индекс добавлен в слайс
					if j+C > len(result)-1 {
						for i := j + 1; i <= len(result)-1; i++ {
							result[i] = i
						}
					} else {
						for i := j + 1; i <= j+C; i++ {
							result[i] = i
						}
					}
				}
			}
			for i, j := range resultCopy {
				if j != 0 || (zeroNil && i == 0) { //если индекс добавлен в слайс
					if j-C < 0 {
						for i := j - 1; i >= 0; i-- {
							result[i] = i
						}
					} else {
						for i := j - 1; i >= j-C; i-- {
							result[i] = i
						}
					}
				}
			}
		}
	}
	var filterFmt string //определяем что пишем в результат в зависимости от флага n
	for i, j := range result {
		if j != 0 || (zeroNil && i == 0) {
			if n { //либо пишем с указанием номера строки
				filterFmt = fmt.Sprintf("%d:%s", j+1, string(dataOriginal[j]))
			} else { //либо просто строку
				filterFmt = string(dataOriginal[j])
			}
			fmt.Println(filterFmt) //выводим результат
		}
	}
}

// проверка на наличие в строке подстроки
func checker(str, substr []byte) bool {
	mc := make(map[byte]int, len(str)+len(substr))
	for _, j := range str {
		mc[j]++
	}
	for _, j := range substr {
		if val, ok := mc[j]; !ok || val < 0 {
			return false
		}
		mc[j]--
	}
	return true
}
