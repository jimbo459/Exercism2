package strand

import "strings"

//ToRNA - given DNA string the RNA string is returned
func ToRNA(dna string) string {
	replacer := strings.NewReplacer("G", "C", "C", "G", "T", "A", "A", "U")
	return replacer.Replace(dna)
}
