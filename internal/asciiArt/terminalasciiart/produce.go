package terminalasciiart

import (
	"bufio"
	"os"
	"strings"
)

type textProduce struct {
	*Text
}

func (t *textProduce) produce() {
	t.output = Converter(t.input, t.Banner)
}

func Converter(input, banner string) string {
	alphabet := getAlphabet(banner)

	words := strings.Split(input, "\\n")

	result := make([]string, 0, len(words))

	for j, word := range words {
		var middleResult string

		for i := 0; i < 8; i++ {

			if word == "" {

				if j == len(words)-1 {
					middleResult = "\n"
				}
				break
			}

			for _, v := range word {
				if v == '\n' {
					continue
				}
				middleResult += alphabet[v][i]
			}

			if i != 7 {
				middleResult += "\n"
			}

			if i == 7 && j == len(words)-1 {
				middleResult += "\n"
			}
		}
		result = append(result, middleResult)
	}

	return strings.Join(result, "\n")
}

func getAlphabet(banner string) map[rune][]string {
	cwd, _ := os.Getwd()
	cwd = trimCwd(cwd)
	file, _ := os.Open(cwd + "/banners/" + banner)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	alphabet := make(map[rune][]string, 110)

	var indexRune rune = 32

	skip := true

	for scanner.Scan() {
		if skip {
			skip = false
			continue
		}
		slice := make([]string, 8)
		for i := 0; i < 8; i++ {
			line := scanner.Text()
			slice[i] = line
			if i != 7 {
				scanner.Scan()
			}
		}

		alphabet[indexRune] = slice
		indexRune++
		skip = true
	}

	return alphabet
}
