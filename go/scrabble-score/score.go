// Package scrabble is used to determine the srabble score of a word.
package scrabble

import "unicode"

// Score - takes a word and returns the scrabble score as an int.
func Score(word string) int {
	if len(word) == 0 {
		return 0
	}

	var score int
	var valueTable = [][]rune{
		[]rune{},
		[]rune{'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't'},
		[]rune{'d', 'g'},
		[]rune{'b', 'c', 'm', 'p'},
		[]rune{'f', 'h', 'v', 'w', 'y'},
		[]rune{'k'},
		[]rune{},
		[]rune{},
		[]rune{'j', 'x'},
		[]rune{},
		[]rune{'q', 'z'},
	}

	for _, char := range word {
		for entryIndex, entry := range valueTable {
			for _, letter := range entry {
				if unicode.ToLower(char) == letter {
					score += entryIndex
				}
			}
		}
	}

	return score
}
