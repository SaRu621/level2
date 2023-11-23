package pkg

import (
	"sort"
	"strings"
)

func isAnagram(str1, str2 string) bool {
	runeStr1 := []rune(strings.ToLower(str1))
	runeStr2 := []rune(strings.ToLower(str2))

	anagramCheck := make(map[rune]int)

	for i := range runeStr1 {
		anagramCheck[runeStr1[i]]++
	}

	for i := range runeStr2 {
		val, exists := anagramCheck[runeStr2[i]]

		if !exists {
			return false
		}

		val--
		anagramCheck[runeStr2[i]] = val

		if anagramCheck[runeStr2[i]] < 0 {
			return false
		}
	}

	for _, key := range anagramCheck {
		if key > 0 {
			return false
		}
	}

	return true
}

//Anagram -- main function for task 4
func Anagram(arr *[]string) map[string][]string {
	anagrams := make(map[string][]string)

	for _, val := range *arr {

		val = strings.ToLower(val)

		flag := true

		for words := range anagrams {
			if isAnagram(val, words) {
				anagrams[words] = append(anagrams[words], val)
				flag = false
				break
			}
		}

		if flag {
			anagrams[val] = []string{}
		}
	}

	for key := range anagrams {
		sort.Strings(anagrams[key])
	}

	return anagrams
}
