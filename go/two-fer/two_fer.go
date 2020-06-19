// Package twofer has the simple function ShareWith to solve the twofer problem
package twofer

// ShareWith returns the string "One for name, one for me." where name is the parameter passed
// If empty name is given, returns the string "One for you, one for me."
func ShareWith(name string) string {

	// Check for empty name
	if name == "" {
		name = "you"
	}

	return "One for " + name + ", one for me."
}
