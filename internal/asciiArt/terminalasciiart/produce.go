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
	t.output = converter(t.input, t.Banner)
}

func converter(input, banner string) string {
	if input == "" {
		return ""
	}

	alphabet := getAlphabet(banner)

	input = strings.ReplaceAll(input, "\\n", "\n")

	words := strings.Split(input, "\n")

	result := make([]string, 0, len(words))

	for j, word := range words {

		if word == "" {
			result = append(result, "")
			continue
		}

		var middleResult string

		for i := 0; i < 8; i++ {
			for _, letter := range word {
				middleResult += alphabet[letter][i]
			}

			if i != 7 || j == len(words)-1 {
				middleResult += "\n"
			}
		}
		result = append(result, middleResult)
	}

	res := strings.Join(result, "\n")

	if len(result) > 1 {
		if result[len(result)-2] != "" && result[len(result)-1] == "" {
			res += "\n"
		}
	}

	return res
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
