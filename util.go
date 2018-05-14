package CommanD_Bot

import (
	"strings"
	"time"
)

// Converts given string array to lower case //
func toLower(input []string) []string {
	// loop through list of input strings //
	for i, v := range input {
		// Convert string to lower case //
		input[i] = strings.ToLower(v)
	}

	// Return array //
	return input
}

// Checks if the given two times have a 14 day or greater deference //
// - True = Less then 14 days old
// - False = Greater then days old
func checkTime(then time.Time, now time.Time) bool {
	// Check if time is greater then 14 days old //
	if then.Year() != now.Year() || (then.YearDay()+14) <= now.YearDay() {
		// time was greater then 14 days //
		return false
	} else {
		// time was less then 14 days old //
		return true
	}
}
