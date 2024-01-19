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
	for indexWord, word := range words {
		if word == "" {
			result = append(result, word)
			continue
		}
		var middleResult string
		for indexHeight := 0; indexHeight < 8; indexHeight++ {
			for _, letter := range word {
				middleResult += alphabet[letter][indexHeight]
			}
			if indexHeight != 7 || indexWord == len(words)-1 {
				middleResult += "\n"
			}
		}
		result = append(result, middleResult)
	}

	result = adjustNewLines(result)

	res := strings.Join(result, "\n")

	return res
}

func adjustNewLines(result []string) []string {
	onlyNewLine := true
	for _, v := range result {
		if v != "" {
			onlyNewLine = false
			break
		}
	}
	if onlyNewLine {
		return result
	}
	toAdd := false
	for _, v := range result {
		if v == "" {
			toAdd = true
			continue
		}
		toAdd = false
	}
	if toAdd {
		result = append(result, "")
	}
	return result
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
