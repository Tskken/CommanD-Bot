package core

import (
	"strings"
	"time"
)

// Checks if the given two times have a 14 day or greater deference //
// - True = Less then 14 days old
// - False = Greater then days old
func CheckTime(then time.Time, now time.Time) bool {
	// Check if time is greater then 14 days old //
	if then.Year() != now.Year() || (then.YearDay()+14) <= now.YearDay() {
		// time was greater then 14 days //
		return false
	} else {
		// time was less then 14 days old //
		return true
	}
}

func IsMentioned(member, mention string) bool {
	if mention[2] == '!' {
		mention = strings.Join(strings.Split(mention, "!"), "")
	}

	return member == mention
}
