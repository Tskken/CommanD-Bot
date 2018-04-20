package CommanD_Bot

import (
	"strconv"
	"strings"
	"time"
)

// Parce user input on a space //
func ParceInput(input string) []string {
	// Return parced input as a list of strings //
	return Parce(input, " ")
}

// Parce given string on the given separator //
func Parce(input, sep string) []string {
	return strings.Split(input, sep)
}

// Converts given string array to lower case //
func ToLower(input []string) []string {
	// Create output string array //
	output := make([]string, 0)
	// loop through list of input strings //
	for _, v := range input {
		// Convert string to lower case //
		s := StrToLower(v)
		// Add new string to output array //
		output = append(output, s)
	}

	// Return output array //
	return output
}

// Converts given string to lowercase //
func StrToLower(input string) string {
	// Sets the specified value to lowercase and return it //
	return strings.ToLower(input)
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

// Checks if the given two times have a 14 day or greater deference //
// - True = is greater then 14 days
// - False = is less then 14 days
func IsTime(then time.Time, now time.Time) bool {
	// Check if time is greater then 14 days old //
	if then.Year() != now.Year() || (then.YearDay()+14) <= now.YearDay() {
		// time was greater then 14 days //
		return true
	} else {
		// time was less then 14 days old //
		return false
	}
}
