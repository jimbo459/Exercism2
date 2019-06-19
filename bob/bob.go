/*
Package bob takes in a string and will respond in the style of bob, depending on the style of input.
*/
package bob

import (
	"strings"
	"unicode"
)

// IsLetter will evaluate a string, returning true as soon as a letter is identified.
func IsLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}

// Hey takes in a string, returning another string based on the user input.
func Hey(remark string) string {
	ts := strings.TrimSpace(remark)
	if len(ts) > 0 {
		if string(ts[len(ts)-1]) == "?" && strings.ToUpper(ts) == ts && IsLetter(ts) {
			return "Calm down, I know what I'm doing!"
		} else if string(ts[len(ts)-1]) == "?" {
			return "Sure."
		} else if strings.ToUpper(ts) == ts && IsLetter(ts) {
			return "Whoa, chill out!"
		}
		return "Whatever."
	}
	return "Fine. Be that way!"
}
