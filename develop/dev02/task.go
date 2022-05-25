package main

import (
	"errors"
	"fmt"
	"strconv"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

//func unpack(str string) string {
//	newStr := ""
//
//	for k, v := range str {
//		_, err := strconv.Atoi(string(v))
//		if err == nil {
//			continue
//		}
//
//		count := 1
//
//		if k+1 < len(str) {
//			nextV := string(str[k+1])
//			num, err := strconv.Atoi(nextV)
//			if err == nil {
//				count = num
//			}
//		}
//
//		for i := 0; i < count; i++ {
//			newStr += string(v)
//		}
//	}
//
//	return newStr
//}

var errWrongSyntax = errors.New("wrong syntax")

func unpack(str string) (string, error) {
	newStr := ""

	isEscape := false
	var prevSym rune = 0

	for _, v := range str {
		if v == '\\' {
			if isEscape {
				newStr += string(v)
				prevSym = v
				isEscape = false
				continue
			}
			isEscape = true
			continue
		}

		amount, err := strconv.Atoi(string(v))
		if err == nil {
			if isEscape {
				newStr += string(v)
				prevSym = v
				isEscape = false
				continue
			}

			if prevSym == 0 {
				return "", errWrongSyntax
			}

			for i := 0; i < amount-1; i++ {
				newStr += string(prevSym)
			}

			prevSym = 0

			continue
		}

		newStr += string(v)
		prevSym = v
	}

	return newStr, nil
}

func main() {
	fmt.Println(unpack("qwe45"))
}
