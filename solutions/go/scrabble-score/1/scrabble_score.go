package scrabble

import "strings"

var ScoreTable = map[string]int{
	"AEIOULNRST": 1,
	"DG":         2,
	"BCMP":       3,
	"FHVWY":      4,
	"K":          5,
	"JX":         8,
	"QZ":         10,
}

func Score(word string) int {
	score := 0

	letterScores := make(map[string]int)

	for letters, score := range ScoreTable {
		for _, letter := range letters {
			letterScores[string(letter)] = score
		}
	}

	word = strings.ToUpper(word)
	for _, letter := range word {
		score += letterScores[string(letter)]
	}

	return score
}
