// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"bytes"
	"regexp"
	"strings"
)

// Abbreviate - input a string, returns Acronym from string.
func Abbreviate(s string) string {
	var buffer bytes.Buffer

	s = strings.Replace(s, "-", " ", -1)

	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString

	buffer.WriteByte(s[0])

	for pos, char := range s {
		if !isAlpha(string(char)) && string(char) != "'" {
			if isAlpha(string(s[pos+1])) {
				buffer.WriteByte(s[pos+1])
			}
		}
	}

	return strings.ToUpper(buffer.String())
}
