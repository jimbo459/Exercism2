/*
Package gigasecond takes in one time as an input and returns the time with a gigasecond added to it.
*/
package gigasecond

import (
	"math"
	"time"
)

// AddGigasecond adds a gigasecond to the time input
func AddGigasecond(t time.Time) time.Time {
	var tAdd = t.Add(time.Second * time.Duration(math.Pow(10, 9)))
	return tAdd
}
