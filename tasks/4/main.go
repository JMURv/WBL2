package main

import (
	"fmt"
	"sort"
	"strings"
)

func findAnagrams(words []string) map[string][]string {
	anagrams := make(map[string][]string)

	for _, v := range words {
		// Сплитуем строку побуквенно, каждую приводим к нижнему регистру
		sorted := strings.Split(strings.ToLower(v), "")

		// Сортируем строку по символам
		sort.Strings(sorted)

		// Это будет наш ключ в мапе
		sortedWord := strings.Join(sorted, "")

		if set, ok := anagrams[sortedWord]; !ok {
			anagrams[sortedWord] = []string{v}
		} else {
			anagrams[sortedWord] = append(set, v)
		}
	}

	for key, set := range anagrams {
		if len(set) <= 1 {
			delete(anagrams, key)
		} else {
			sort.Strings(set)
		}
	}

	return anagrams
}

func main() {
	dictionary := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	fmt.Println(findAnagrams(dictionary))
}
