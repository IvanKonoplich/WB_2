package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

// Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:
// "a4bc2d5e" => "aaaabccddddde"
// "abcd" => "abcd"
// "45" => "" (некорректная строка)
// "" => ""
//
// Дополнительно
// Реализовать поддержку escape-последовательностей.
// Например:
// qwe\dev04\dev05 => qwe45 (*)
// qwe\45 => qwe44444 (*)
// qwe\\dev05 => qwe\\\\\ (*)
//
// В случае если была передана некорректная строка, функция должна возвращать ошибку. Написать unit-тесты.
func main() {
	var inc string
	fmt.Scan(&inc)
	result, err := unpacking(inc)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
		fmt.Println(len([]rune(result)))
	}

}

func unpacking(inc string) (string, error) {
	incRune := []rune(inc)
	var result string
	//проверка на пустую строку
	if len(incRune) == 0 {
		return "", nil
	}
	//проверка на цифру в начале
	if unicode.IsDigit(incRune[0]) {
		return "", errors.New("первое значение неэкранированная цифра")
	}
	for i := 0; i < len(incRune); i++ {
		//если значение \, то переходим через одно значение вперед и добавляем его
		if (incRune[i]) == 92 {
			//если оно в конце, то возвращаем ошибку
			if i == len(incRune)-1 {
				return "", errors.New("экранирование в конце")
			}
			i++
			result += string(incRune[i])
			continue
		}
		//если цифра, то умножаем предыдущее значение на нее -dev01 ведь dev01 раз его уже добавили
		if unicode.IsDigit(incRune[i]) {
			digit, lastIndex := findFullDigit(incRune, i)
			for q := 0; q < digit-1; q++ {
				result += string(incRune[i-1])
			}
			i = lastIndex
		}

		//если значение буква, то добавляем его
		if unicode.IsLetter(incRune[i]) {
			result += string(incRune[i])
		}
	}
	return result, nil
}

// находим число полностью если например "a45"
func findFullDigit(r []rune, index int) (digit int, lastIndex int) {
	result := runeToInt(r[index])
	index++
	for ; index < len(r); index++ {
		if unicode.IsDigit(r[index]) {
			result = result * 10
			result += runeToInt(r[index])
		} else {
			return result, index
		}
	}
	return result, index - 1
}
func runeToInt(r rune) int {
	result, _ := strconv.Atoi(string(r))
	return result
}
