package proverb

import (
	"fmt"
)

// For want of a nail the shoe was lost.
// For want of a shoe the horse was lost.
// For want of a horse the rider was lost.
// For want of a rider the message was lost.
// For want of a message the battle was lost.
// For want of a battle the kingdom was lost.
// And all for the want of a nail.

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {

	if len(rhyme) <= 0 {
		return []string{}
	}

	result := []string{}

	if len(rhyme) > 1 {
		for i := range len(rhyme) - 1 {
			current := rhyme[i]
			next := rhyme[i+1]
			result = append(result, fmt.Sprintf("For want of a %s the %s was lost.", current, next))
		}
	}

	result = append(result, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))

	return result
}
