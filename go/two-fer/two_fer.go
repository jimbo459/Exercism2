package twofer

// ShareWith usage: supply a name to amend the returned string.
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	} else {
		return "One for " + name + ", one for me."
	}
}
