package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type Flags struct {
	F int
	D string
	S bool
}

func main() {
	flagF := flag.Int("f", 0, "")
	flagD := flag.String("d", " ", "")
	flagS := flag.Bool("s", false, "")

	flag.Parse()

	f := Flags{
		F: *flagF,
		D: *flagD,
		S: *flagS,
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

	for _, line := range lines {
		cols := strings.Split(line, f.D)

		if f.F > 0 && f.F <= len(cols) {
			fmt.Println(cols[f.F-1])
		} else if !f.S {
			fmt.Println(line)
		}
	}
}
