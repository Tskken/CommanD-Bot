package CommanD_Bot

import (
	"strings"
	"time"
)

// Converts given string array to lower case //
func toLower(input []string) []string {
	// Create output string array //
	output := make([]string, 0)
	// loop through list of input strings //
	for _, v := range input {
		// Convert string to lower case //
		s := strings.ToLower(v)
		// Add new string to output array //
		output = append(output, s)
	}

	// Return output array //
	return output
}

// Checks if the given two times have a 14 day or greater deference //
// - True = is greater then 14 days
// - False = is less then 14 days
func isTime(then time.Time, now time.Time) bool {
	// Check if time is greater then 14 days old //
	if then.Year() != now.Year() || (then.YearDay()+14) <= now.YearDay() {
		// time was greater then 14 days //
		return true
	} else {
		// time was less then 14 days old //
		return false
	}
}
