package CommanD_Bot

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
	"log"
)

// Parce given string on a space //
// input: Given input string to parce
// Returns a string array of parced input
func ParceInput(input string) []string {
	// Return parced input //
	return strings.Split(input, " ")
}

// TODO - Comment
func ToLower(input []string, i int) *string {
	if len(input) <= i {
		log.Println("ToLower:", "given location to ToLower is outside of bounds of given array.")
		return nil
	}

	arg := strings.ToLower(input[i])

	return &arg
}
