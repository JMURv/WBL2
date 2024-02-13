package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Test str: aaa:sss:ddd
// go run . -f 1,3 -d ":" | Out: "aaa ddd"
// go run . -f 1,3 -d ":" -s | Out: "aaa ddd"
// Test str: aaa	sss	ddd
// go run . -f 1,3 | Out: "aaa ddd"

func contains(slice []int, v int) bool {
	for _, s := range slice {
		if v == s {
			return true
		}
	}
	return false
}

func cut(msg string, fields []int, delimiter string, sep bool) {
	if !sep || strings.Contains(msg, delimiter) {
		splited := strings.Split(msg, delimiter)
		r := make([]string, 0, len(splited))
		for i, v := range splited {
			if contains(fields, i+1) {
				r = append(r, v)
			}
		}
		fmt.Println(strings.Join(r, " "))
	}
}

func main() {
	fieldStr := flag.String("f", "", "fields")
	delimiter := flag.String("d", "\t", "delimiter")
	sep := flag.Bool("s", false, "separated")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Type strings for cut: ")
	msg, _ := reader.ReadString('\n') // Читаем инпут пока не встретим line breaker(корректно обрабатывает табуляции)
	msg = strings.TrimSpace(msg)

	// Получаем все необходимые для вывода поля
	var fields []int
	for _, field := range strings.Split(*fieldStr, ",") {
		if field == "" {
			continue
		}
		f, _ := strconv.Atoi(field)
		fields = append(fields, f)
	}

	cut(msg, fields, *delimiter, *sep)
}
