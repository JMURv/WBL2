package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func allDigits(str string) bool {
	for _, char := range str {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func unpackString(str string) (string, error) {
	if len(str) == 0 {
		return "", nil
	}

	if allDigits(str) {
		return "", fmt.Errorf("некорректная строка")
	}

	var r strings.Builder
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		char := runes[i]
		if char == '\\' {
			if i+1 <= len(runes)-1 {
				r.WriteRune(runes[i+1])
				i++
				continue
			} else {
				return "", fmt.Errorf("некорректная escape-последовательность в конце строки")
			}
		}

		if unicode.IsDigit(char) && i-1 >= 0 {
			prevChar := runes[i-1]
			c, err := strconv.Atoi(string(char))
			if err != nil {
				return "", err
			}
			r.WriteString(strings.Repeat(string(prevChar), c-1))
		} else {
			r.WriteRune(char)
		}
	}

	return r.String(), nil
}

func main() {
	strs := []string{"a4bc2d5e", "abcd", "45", "", "qwe\\4\\5", "qwe\\45", "qwe\\\\5"}

	for _, str := range strs {
		r, err := unpackString(str)
		if err != nil {
			fmt.Printf("Ошибка: %v\n", err)
		} else {
			fmt.Printf("Исходная строка: %s\nРезультат: %s\n", str, r)
		}
	}
}
