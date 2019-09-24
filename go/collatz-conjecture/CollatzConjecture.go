package collatzconjecture

import "errors"

// CollatzConjecture - input int, get number of steps required to return 1 using Collatz Conjecture.
func CollatzConjecture(number int) (int, error) {
	collatzCount := 0
	if number > 0 {
		for number > 1 {
			if number%2 != 0 {
				number = (number * 3) + 1
			} else {
				number = number / 2
			}
			collatzCount++
		}
		return collatzCount, nil
	}
	return collatzCount, errors.New("Number must be greater than 0")
}
