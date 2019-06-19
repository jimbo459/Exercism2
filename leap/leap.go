package leap

// IsLeapYear returns true if input year is leap year.
func IsLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			} else {
				return false
			}
		}
		return true
	}
	return false
}
