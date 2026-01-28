package isogram

import "strings"

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	letterMap := make(map[string]int)
	for _, letter := range word {
		if letter != ' ' && letter != '-' {
			if _, ok := letterMap[string(letter)]; ok {
				return false
			}
			letterMap[string(letter)]++
		}
	}
	return true
}
