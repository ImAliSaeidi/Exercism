package isogram

func IsIsogram(word string) bool {
	letterMap := make(map[string]int)
	for _, letter := range word {
		if _, ok := letterMap[string(letter)]; ok {
			return false
		}
		letterMap[string(letter)]++
	}
	return true
}
