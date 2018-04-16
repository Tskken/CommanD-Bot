package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tsukinai/CommanD-Bot/utility"
	"math/rand"
)

// Roles a dice and prints the results //
// Returns an error (nil if non)
func diceRole(s *discordgo.Session, m *discordgo.Message) error {
	// Parce messages on a space //
	args := utility.ParceInput(m.Content)

	// Convert the third argument to an int //
	// - Returns an error if err is not nil
	if val, err := utility.StrToInt(args[2]); err != nil {
		return err
	} else {
		// Get a random number from 0 to the given value //s
		rnd := rand.Intn(val)

		// Print random number to channel //
		// - Returns an error if err is not nil
		if _, err := s.ChannelMessageSend(m.ChannelID, m.Author.Username+" got "+utility.IntToStr(rnd)); err != nil {
			return err
		}
	}
	return nil
}

// TODO - Implement IGN
func ign(s *discordgo.Session, m *discordgo.Message) error {
	return nil
}

// TODO - Implement Trinity
func trinity(s *discordgo.Session, m *discordgo.Message) error {
	return nil
}
