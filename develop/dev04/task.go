package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func sortAnagrams(words []string) map[string][]string {
	m := make(map[string]map[string]struct{})

	for _, word := range words {
		lowerWord := strings.ToLower(word)

		slice := []rune(lowerWord)
		sort.Slice(slice, func(i, j int) bool {
			return slice[i] > slice[j]
		})
		key := string(slice)

		if _, ok := m[key]; !ok {
			m[key] = make(map[string]struct{})
		}

		m[key][lowerWord] = struct{}{}
	}

	result := make(map[string][]string)

	for _, v := range m {
		if len(v) < 2 {
			continue
		}

		slice := make([]string, 0)
		for w := range v {
			slice = append(slice, w)
		}
		sort.Strings(slice)
		key := slice[0]
		result[key] = slice
	}

	return result
}

func main() {
	words := []string{"столик", "пятак", "листок", "пятка", "слиток", "тяпка", "кот", "слиток"}

	fmt.Println(sortAnagrams(words))
}
