//Package scale ...
package scale

import (
	"strings"
)

var (
	returnScale []string
	position    int
	sharpNotes  = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	flatNotes   = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
	flatTonics  = "Fdgcf"
	intervals   = map[rune]int{
		'm': 1,
		'M': 2,
		'A': 3,
	}
)

//Scale - provide starting note and the intervals to return the scale.
func Scale(tonic string, interval string) []string {
	scaleKey := identifyKey(tonic)
	tonicIndex := getTonicIndex(tonic, scaleKey)
	return scaleBuilder(tonicIndex, scaleKey, interval)
}

func identifyKey(tonic string) []string {
	if strings.Index(flatTonics, tonic) > -1 {
		return flatNotes
	} else if len(tonic) > 1 {
		if tonic[1:] == "b" {
			return flatNotes
		}
	}
	return sharpNotes
}

func getTonicIndex(tonic string, scale []string) int {

	if len(tonic) > 1 {
		tonic = strings.ToUpper(tonic[:1]) + tonic[1:]
	} else {
		tonic = strings.ToUpper(tonic)
	}

	for k, v := range scale {
		if v == tonic {
			return k
		}
	}
	return 0
}

func scaleBuilder(tonicIndex int, scale []string, interval string) []string {
	var returnScale []string
	position := tonicIndex

	if len(interval) == 0 {
		for i := 0; i < 12; i++ {
			returnScale = append(returnScale, scale[position])
			position = (position + 1) % 12
		}
	} else {
		for _, v := range interval {
			returnScale = append(returnScale, scale[position])
			position = (position + intervals[v]) % 12
		}
	}
	return returnScale
}
