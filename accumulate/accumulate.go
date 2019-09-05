package accumulate


type operation func(string) string

func Accumulate(givenStrings []string, op operation) []string {
	var newStrings []string

	for _,v := range givenStrings {
		newStrings = append(newStrings,op(v))
	}

	return newStrings
}