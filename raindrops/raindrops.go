// Package raindrops returns a string if the given input has a factor of 3,5 or 7.
package raindrops

import (
	"bytes"
	"strconv"
)

// Convert - accepts int as input, returns string if input has factor of 3,5 or 7.
func Convert(i int) string {
	var wordMap = map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}

	keys := [3]int{3, 5, 7}

	var buffer bytes.Buffer

	for _, k := range keys {
		if i%k == 0 {
			buffer.WriteString(wordMap[k])
		}
	}
	if buffer.String() == "" {
		return strconv.Itoa(i)
	}

	return buffer.String()
}
