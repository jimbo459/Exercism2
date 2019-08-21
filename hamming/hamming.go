// Package hamming allows you to calculate the hamming difference between two DNA strands.
package hamming

import "errors"

//Distance - given two DNA strands of type string, hamming difference is returned. Strings must be of equal length.
func Distance(a, b string) (int, error) {
	newA := []rune(a)
	newB := []rune(b)

	if len(newA) != len(newB) {
		return 0, errors.New("DNA strands must be of equal length")
	}

	hammingDistance := 0

	for k, v := range newA {
		if newB[k] != v {
			hammingDistance++
		}
	}

	return hammingDistance, nil
}
