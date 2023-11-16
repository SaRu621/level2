package pkg

import (
	"strconv"
	"unicode"
)

func Extraction(archive string) []rune {
	extract := make([]rune, 0)

	runeArchive := []rune(archive)

	var symbol rune
	var count []rune
	var haveSlash bool = false

	for _, val := range runeArchive {
		if val == '\\' && !haveSlash {
			haveSlash = true
			continue
		}

		if !unicode.IsDigit(val) || haveSlash {
			num, _ := strconv.Atoi(string(count))

			if symbol != rune(0) {
				extract = append(extract, symbol)
			}

			for i := 0; i < num-1; i++ {
				if symbol != rune(0) {
					extract = append(extract, symbol)
				}
			}

			count = []rune{}

			haveSlash = false
			symbol = val
		} else {
			count = append(count, val)
		}
	}

	num, _ := strconv.Atoi(string(count))

	if symbol != rune(0) {
		extract = append(extract, symbol)
	}

	for i := 0; i < num-1; i++ {
		if symbol != rune(0) {
			extract = append(extract, symbol)
		}
	}
	return extract
}
