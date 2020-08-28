package twofer

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	noname := "you"
	if name != "" {
		noname = name
	}
	return "One for " + noname + ", one for me."
}