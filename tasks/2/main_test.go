package main

import (
	"strings"
	"testing"
)

var EscapeChar = "\\"

func TestUnpackString(t *testing.T) {
	cases := []struct {
		input    string
		expected string
		err      bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"45", "", true},
		{"", "", false},
		{"qwe" + EscapeChar + "4" + EscapeChar + "5", "qwe45", false},
		{"qwe" + EscapeChar + "45", "qwe44444", false},
		{"qwe" + EscapeChar + EscapeChar + "5", "qwe" + strings.Repeat(EscapeChar, 5), false},
	}

	for _, tc := range cases {
		result, err := unpackString(tc.input)

		if (err != nil) != tc.err {
			t.Errorf("Для данных %s ожидалась ошибка: %v, фактически получено: %v", tc.input, tc.err, err)
		}

		if result != tc.expected {
			t.Errorf("Для данных %s ожидалось %s, получено %s", tc.input, tc.expected, result)
		}
	}
}
