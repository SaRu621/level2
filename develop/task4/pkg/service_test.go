package pkg

import (
	"fmt"
	"testing"
)

var testArr = [][]string{
	{"Пятка", "Тяпка", "пятак", "листок", "столик", "слиток"},
}

var testAns = []map[string][]string{
	{
		"листок": {
			"слиток",
			"столик",
		},
		"пятка": {
			"пятак",
			"тяпка",
		},
	},
}

func slicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func mapsEqual(a, b map[string][]string) bool {
	if len(a) != len(b) {
		return false
	}

	for key, valueA := range a {
		valueB, ok := b[key]
		if !ok || !slicesEqual(valueA, valueB) {
			return false
		}
	}

	return true
}

func TestPkg(t *testing.T) {
	for i, val := range testArr {
		if mapsEqual(testAns[i], Anagram(&val)) {
			fmt.Printf("Test %d passed.\n", i)
		} else {
			t.Errorf("Test %d not passed.\n", i)
		}
	}
}
