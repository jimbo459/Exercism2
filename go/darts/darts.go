// Package Darts contains a method score to return the score of a dart at a given coordinate.
package darts

// Score - given two co-ordinates, a score is returned as an integer.
func Score(x, y float64) int {

	distance := x*x + y*y

	if distance <= 1 {
		return 10
	} else if distance <= 25 {
		return 5
	} else if distance <= 100 {
		return 1
	} else {
		return 0
	}

}
