/*
Package triangle take in three lengths and determines if the shape is a triangle.
The type of triangle is returned as a string.
Accepted input of float64
*/
package triangle

import "math"

type Kind string

const (
	NaT Kind = "NaT" // not a triangle
	Equ Kind = "Equ" // equilateral
	Iso Kind = "Iso" // isosceles
	Sca Kind = "Sca" // scalene
)

func KindFromSides(a, b, c float64) Kind {
	if (a > 0) && (b > 0) && (c > 0) && !math.IsInf(a+b+c, 0) {
		if c <= a+b && b <= a+c && a <= b+c {
			if a == b && a == c {
				return Equ
			} else if a == b || b == c || a == c {
				return Iso
			} else {
				return Sca
			}
		}
	}
	return NaT
}
