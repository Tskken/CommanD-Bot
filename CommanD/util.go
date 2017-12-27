package CommanD

/*
Last Updated: 11/7/27
Author: Dylan Blanchard

util.go

All non bot related functions
*/

import (
	// Golang imports //
	"strings"
	// External imports //
)

// Parce given string on a space //
// input: Given input string to parce
// Returns a string array of parced input
func ParceInput(input string) []string {
	// Return parced input //
	return strings.Split(input, " ")
}
