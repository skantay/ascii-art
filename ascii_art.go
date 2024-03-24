package asciiart

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func GetASCII(word string, file io.Reader) (string, error) {
	return getASCII(word, file)
}

func getASCII(input string, file io.Reader) (string, error) {
	font, err := getFont(file)
	if err != nil {
		return "", fmt.Errorf("failed to get font: %w", err)
	}

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
				middleResult += font[letter][indexHeight]
			}
			if indexHeight != 7 || indexWord == len(words)-1 {
				middleResult += "\n"
			}
		}
		result = append(result, middleResult)
	}

	result = adjustNewLines(result)

	res := strings.Join(result, "\n")

	return res, nil
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

func getFont(file io.Reader) (map[rune][]string, error) {
	scanner := bufio.NewScanner(file)

	alphabet := make(map[rune][]string, 150)

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

	return alphabet, nil
}
