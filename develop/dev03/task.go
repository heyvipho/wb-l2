package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Flags struct {
	K int
	N bool
	R bool
	U bool
	M bool
	B bool
	C bool
	H bool
	O string
}

func clearDuplicates(lines []string) []string {
	m := make(map[string]struct{})
	for _, line := range lines {
		m[line] = struct{}{}
	}
	s := make([]string, 0)
	for line := range m {
		s = append(s, line)
	}
	return s
}

func getMonthNum(month string) (int, bool) {
	switch month {
	case "jan":
		return 0, true
	case "feb":
		return 1, true
	case "mar":
		return 2, true
	case "apr":
		return 3, true
	case "may":
		return 4, true
	case "jun":
		return 5, true
	case "jul":
		return 6, true
	case "aug":
		return 7, true
	case "sep":
		return 8, true
	case "oct":
		return 9, true
	case "nov":
		return 10, true
	case "dec":
		return 11, true
	default:
		return 0, false
	}
}

func getSufixNum(strNum string) (int, int, bool) {
	strNumLower := strings.ToLower(strNum)

	re := regexp.MustCompile("[-0-9]+")
	nStr := re.FindString(strNumLower)
	n, _ := strconv.Atoi(nStr)

	re = regexp.MustCompile("[mgtprezy]{0,1}")
	letter := re.FindString(strNumLower)

	switch letter {
	case "m":
		return n, 0, true
	case "g":
		return n, 1, true
	case "t":
		return n, 2, true
	case "p":
		return n, 3, true
	case "r":
		return n, 4, true
	case "e":
		return n, 5, true
	case "z":
		return n, 6, true
	case "y":
		return n, 7, true
	}

	return 0, 0, false
}

func getSortFunc(lines []string, f Flags) func(i, j int) bool {
	return func(i, j int) bool {
		a := lines[i]
		b := lines[j]

		if f.K > 0 {
			a = ""
			b = ""
			if len(lines[i]) >= f.K {
				re := regexp.MustCompile("[ ]+")
				split := re.Split(lines[i], -1)
				a = split[f.K-1]
			}
			if len(lines[i]) >= f.K {
				re := regexp.MustCompile("[ ]+")
				split := re.Split(lines[j], -1)
				b = split[f.K-1]
			}
		}

		if f.N {
			re := regexp.MustCompile("^[-0-9]+")

			aNum := 0
			aExist := true
			bNum := 0
			bExist := true

			n := re.FindString(a)
			if len(n) > 0 {
				num, err := strconv.Atoi(n)
				if err != nil {
					aExist = false
				}
				aNum = num
			}

			n = re.FindString(b)
			if len(n) > 0 {
				num, err := strconv.Atoi(n)
				if err != nil {
					bExist = false
				}
				bNum = num
			}

			if aExist || bExist {
				if f.R {
					return aNum > bNum
				}
				return aNum < bNum
			}
		}

		if f.H {
			re := regexp.MustCompile("(?i)^[-0-9]+[mgtprezy]{0,1}")

			aNum := 0
			aLetter := 0
			aExist := true
			bNum := 0
			bLetter := 0
			bExist := true

			n := re.FindString(a)
			if len(n) > 0 {
				aNum, aLetter, aExist = getSufixNum(n)
			}

			n = re.FindString(b)
			if len(n) > 0 {
				bNum, bLetter, bExist = getSufixNum(n)
			}

			if aExist || bExist {
				if f.R {
					return aLetter > bLetter || aNum > bNum
				}
				return aLetter < bLetter || aNum < bNum
			}
		}

		if f.M {
			aNum := 0
			aExist := true
			bNum := 0
			bExist := true

			re := regexp.MustCompile("(?i)^jan|feb|mar|apr|may|jun|jul|aug|sep|oct|nov|dec")

			month := strings.ToLower(re.FindString(a))
			aNum, aExist = getMonthNum(month)

			month = strings.ToLower(re.FindString(b))
			bNum, bExist = getMonthNum(month)

			if aExist || bExist {
				if f.R {
					return aNum > bNum
				}
				return aNum < bNum
			}
		}

		if f.B {
			a = strings.Trim(a, " ")
			b = strings.Trim(b, " ")
		}

		if f.R {
			return a > b
		}
		return a < b
	}
}

func sortLines(lines []string, f Flags) []string {
	sort.Slice(lines, getSortFunc(lines, f))

	return lines
}

func checkLines(lines []string, f Flags) bool {
	sorter := getSortFunc(lines, f)

	isSorted := true

	for i := 1; i < len(lines); i++ {
		isSorted = sorter(i-1, i)
	}

	return isSorted
}

func main() {
	flagK := flag.Int("k", 0, "")
	flagN := flag.Bool("n", false, "")
	flagR := flag.Bool("r", false, "")
	flagU := flag.Bool("u", false, "")
	flagM := flag.Bool("M", false, "")
	flagB := flag.Bool("b", false, "")
	flagC := flag.Bool("c", false, "")
	flagH := flag.Bool("h", false, "")
	flagO := flag.String("o", "", "")

	flag.Parse()

	f := Flags{
		K: *flagK,
		N: *flagN,
		R: *flagR,
		U: *flagU,
		M: *flagM,
		B: *flagB,
		C: *flagC,
		H: *flagH,
		O: *flagO,
	}

	filename := ""
	args := flag.Args()
	if len(args) > 1 {
		log.Fatalln(ErrIncorrectFile)
	}
	if len(args) == 1 {
		filename = args[0]
	}

	input := ""
	if filename != "" {
		bytes, err := os.ReadFile(filename)
		if err != nil {
			log.Fatalln(err)
		}
		input = string(bytes)
	} else {
		bytes, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatalln(err)
		}
		input = string(bytes)
	}

	lines := strings.Split(input, "\n")

	if f.C {
		if checkLines(lines, f) {
			fmt.Println("Yes, it's sorted")
		} else {
			fmt.Println("No, it isn't sorted")
		}

		return
	}

	if f.U {
		lines = clearDuplicates(lines)
	}

	lines = sortLines(lines, f)

	output := strings.Join(lines, "\n")

	if f.O != "" {
		err := os.WriteFile(f.O, []byte(output), 0644)
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		fmt.Println(output)
	}
}
