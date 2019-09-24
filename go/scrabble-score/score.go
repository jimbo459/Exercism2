package scrabble

import "fmt"

func Score(word string) (int) {
	var score int

	var valueTable [][]rune

	valueTable = [][]rune{
		[]rune{},
		[]rune{'a','e', 'i', 'o', 'u','l', 'n', 'r','s','t', 'A','E', 'I', 'O', 'U','L', 'N', 'R','S','T'},
		[]rune{'d', 'g', 'D', 'G'},
		[]rune{'b', 'c', 'm', 'p', 'B', 'C', 'M', 'P'},
		[]rune{'f', 'h', 'v', 'w', 'y', 'F', 'H', 'V', 'W', 'Y'},
		[]rune{'k', 'K'},
		[]rune{},
		[]rune{},
		[]rune{'j', 'x', 'J', 'X'},
		[]rune{},
		[]rune{'q', 'z', 'Q', 'Z'},
	}

	for _, char := range word {
		for entryIndex, entry := range valueTable {
			for _, letter := range entry {
					if char == letter {
						score += entryIndex
				}
			}
		}
	}

	return score
}

func print(letter rune, index int) {
	fmt.Printf("%v, %v\n", string(letter), index)
}
