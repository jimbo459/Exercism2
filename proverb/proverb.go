/*
Package proverb takes in arguments of type string, returning a proverb with the given input.
*/
package proverb

/*Proverb will take a string array or slice as input and return a string array. */
func Proverb(rhyme []string) []string {
	var s []string
	if len(rhyme) > 0 {
		var baseText = []string{"For want of a ", " was lost.", "And all for the want of a ", "."}
		for x := 0; x < len(rhyme)-1; x++ {
			s = append(s, baseText[0]+rhyme[x]+" the "+rhyme[x+1]+baseText[1])
		}
		s = append(s, baseText[2]+rhyme[0]+baseText[3])
	}
	return s
}
