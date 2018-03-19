package utility

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

// TODO - Comment
func Parce(input, sep string) []string {
	return strings.Split(input, sep)
}

func ToLower(input []string) []string {
	output := make([]string, 0)
	for _, v := range input {
		s := StrToLower(v)
		output = append(output, s)
	}

	return output
}

// Changes a given value from with in a given list of strings to lowercase //
func StrToLower(input string) string {
	// Sets the specified value in the list to lowercase //
	arg := strings.ToLower(input)

	// Return a reference to the new lowercase string //
	return arg
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

// TODO - Comment
func IsTime(t1 time.Time, t2 time.Time) bool {
	if t1.Year() != t2.Year() || (t1.YearDay()+14) <= t2.YearDay() {
		return true
	} else {
		return false
	}
}
