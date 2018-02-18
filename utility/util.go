package utility

import (
	"github.com/tsukinai/CommanD-Bot/botErrors"
	"strconv"
	"strings"
)

// Parce user input on a space //
func ParceInput(input string) []string {
	// Return parced input as a list of strings //
	return strings.Split(input, " ")
}

// Changes a given value from with in a given list of strings to lowercase //
func ToLower(input []string, i int) (*string, error) {
	// If the length of the list is less then or equal to i then return an error //
	// Returns an error that the given value to set to lowercase is outside the bounds of the list
	if len(input) <= i {
		return nil, botErrors.NewError("Given location to ToLower is outside of bounds of given array", "util.go")
	}

	// Sets the specified value in the list to lowercase //
	arg := strings.ToLower(input[i])

	// Return a reference to the new lowercase string //
	return &arg, nil
}

// Converts a list of strings to a string //
func ToString(input []string, sep string) string {
	// Convert the list to a string with a given separator //
	return strings.Join(input, sep)
}

// Convert string to int //
// Returns an error if conversion fails
func StrToInt(input string) (int, error) {
	return strconv.Atoi(input)
}

// Convert int to string //
func IntToStr(input int) string {
	return strconv.Itoa(input)
}
