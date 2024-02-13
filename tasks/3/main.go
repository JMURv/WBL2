package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fpath := flag.String("f", "", "Путь к файлу для сортировки")
	column := flag.Int("k", 0, "Номер колонки для сортировки (по умолчанию 0, разделитель - пробел)")
	numeric := flag.Bool("n", false, "Сортировать по числовому значению")
	reverse := flag.Bool("r", false, "Сортировать в обратном порядке")
	unique := flag.Bool("u", false, "Не выводить повторяющиеся строки")

	flag.Parse()

	if *fpath == "" {
		fmt.Println("Необходимо указать путь к файлу -f")
		return
	}

	lines, err := readFileLines(*fpath)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	sortLines(&lines, *column, *numeric, *reverse, *unique)
	err = writeLinesToFile(*fpath, lines)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
		return
	}
}

func sortLines(lines *[]string, col int, num, rev, uni bool) {
	compareFunc := func(i, j int) bool {
		if num {
			numI, errI := strconv.ParseInt(strings.Fields((*lines)[i])[col], 10, 64)
			numJ, errJ := strconv.ParseInt(strings.Fields((*lines)[j])[col], 10, 64)

			if errI == nil && errJ == nil {
				return numI < numJ
			}
		}

		return (*lines)[i] < (*lines)[j]
	}

	if rev {
		sort.SliceStable(*lines, func(i, j int) bool {
			return !compareFunc(i, j)
		})
	} else {
		sort.SliceStable(*lines, compareFunc)
	}

	if uni {
		removeDuplicates(lines)
	}
}

// Удаляем дубликаты из массива по месту
func removeDuplicates(lines *[]string) {
	seen := make(map[string]struct{}, len(*lines))

	idx := 0
	for _, v := range *lines {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			(*lines)[idx] = v
			idx++
		}
	}
	*lines = (*lines)[:idx]
}

// Вспомогательные функции для чтения/записи файла
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

func writeLinesToFile(fpath string, lines []string) error {
	f, err := os.Create("result.txt")
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, line := range lines {
		_, err := w.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return w.Flush()
}
