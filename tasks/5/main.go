package main

import (
	"bufio"
	"os"
)

// Обчный поиск: go run main.go -i "enim" test.txt
// Обычный поиск с контекстом 2: go run main.go -i -C 2 "enim" test.txt
// Инвертирование: go run main.go -i -v "enim" test.txt
// С номером строк: go run main.go -i -n "enim" test.txt
// Подсчёт строк: go run main.go -i -c "enim" test.txt

import (
	"flag"
	"fmt"
	"regexp"
)

func readFileLines(fpath string) ([]string, error) {
	f, err := os.Open(fpath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	r := bufio.NewScanner(f)
	for r.Scan() {
		lines = append(lines, r.Text())
	}
	return lines, r.Err()
}

func main() {
	after := flag.Int("A", 0, "печатать N строк после совпадения")
	before := flag.Int("B", 0, "печатать N строк до совпадения")
	context := flag.Int("C", 0, "печатать N строк вокруг совпадения")
	count := flag.Bool("c", false, "количество строк")
	ignoreCase := flag.Bool("i", false, "игнорировать регистр")
	invert := flag.Bool("v", false, "вместо совпадения, исключать")
	fixed := flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	lineNum := flag.Bool("n", false, "напечатать номер строки")
	flag.Parse()

	var regexPattern *regexp.Regexp
	pattern := flag.Arg(0)

	if *fixed {
		pattern = regexp.QuoteMeta(pattern)
	}
	if *ignoreCase {
		regexPattern, _ = regexp.Compile("(?i)" + pattern)
	} else {
		regexPattern, _ = regexp.Compile(pattern)
	}

	lines, _ := readFileLines(flag.Arg(1))
	lineCount := 0
	for _, line := range lines {
		lineCount++

		matched := regexPattern.MatchString(line)
		if matched && !*invert {
			if *count {
				continue
			}
			printGrep(lines, lineCount, *before, *after, *context, *lineNum)
		} else if !matched && *invert {
			if *count {
				lineCount++
				continue
			}
			printGrep(lines, lineCount, *before, *after, *context, *lineNum)
		}
	}

	if *count {
		fmt.Println(lineCount)
	}
}

func printGrep(lines []string, lineCount, before, after, context int, printLineNum bool) {
	start := lineCount - before - context - 1
	if start < 0 {
		start = 0
	}

	end := lineCount + after + context
	if end > len(lines) {
		end = len(lines)
	}

	for i := start; i < end; i++ {
		line := lines[i]
		if printLineNum {
			fmt.Printf("%d: %s\n", i+1, line)
		} else {
			fmt.Println(line)
		}
	}
}
