package CommanD_Bot

import (
	"strings"
)

// Parce given string on a space //
// input: Given input string to parce
// Returns a string array of parced input
func ParceInput(input string) []string {
	// Return parced input //
	return strings.Split(input, " ")
}

// TODO - Comment
func ToLower(input []string, i int) (*string, error) {
	if len(input) <= i {
		return nil, NewError("Given location to ToLower is outside of bounds of given array", "util.go")
	}

	arg := strings.ToLower(input[i])

	return &arg, nil
}

// TODO - Comment
func ToString(input []string) string {
	return strings.Join(input, " ")
}

