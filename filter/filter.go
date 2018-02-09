package filter

import (
"strings"
"github.com/bwmarrin/discordgo"
)

/*
WIP

This files contents are currently just temp functions and variables to try and outline what may be needed for the learning algorithm.

TODO - Work on implementing algorithm.
*/

var filter = make(map[string]int)

func StartFilter(s *discordgo.Session, input string) error {
	parcedString := strings.Split(input, " ")
	if parcedString == nil {
		return nil
	}
	return nil
}
