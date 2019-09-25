package romannumerals


// ToRomanNumeral - converts Arabic number to Roman Numeral
func ToRomanNumeral(number int) (string, error) {
	var romanNumber string

	if number == 0 {
		return "", nil
	} 

	if number < 4 {
		upperLimit =
		for i := 0; i < number; i++ {
			romanNumber = romanNumber + "I"
			
		}

	} else if number == 4 {
		romanNumber =  romanNumber +  "IV"
		recursionResult, _ := ToRomanNumeral(number - 4)
		romanNumber += recursionResult

	} else if number == 5 {
		romanNumber = romanNumber +  "V"
		recursionResult, _ := ToRomanNumeral(number - 5)
		romanNumber += recursionResult

	} else if number < 9 {
		romanNumber = romanNumber + "V"
		recursionResult, _ := ToRomanNumeral(number - 5)
		romanNumber += recursionResult
	} else if number == 9 {
		romanNumber = romanNumber + "IX"
		recursionResult, _ := ToRomanNumeral(number - 9)
		romanNumber += recursionResult

	} else if number < 40 {
		for i := 0; i < (number / 10); i++ {
			romanNumber = romanNumber + "X" 
			number -= 10
		}

		recursionResult, _ := ToRomanNumeral(number)
		romanNumber += recursionResult
	}

	return romanNumber, nil
}

