package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tsukinai/CommanD-Bot/utility"
)

// Wrapper function to call all Utility commands //
func UtilityCommands(s *discordgo.Session, m *discordgo.Message, admin bool) error {
	// Get the argument passed to !utility and make sure its lowercase //
	// Returns an error if err is nil
	if arg, err := utility.ToLower(utility.ParceInput(m.ChannelID), 1); err != nil {
		return err
	} else {
		// Get the arguments command //
		// Prints an error if the command does not exist //
		if cmd, ok := utilityCommands[*arg]; !ok {
			_, err := s.ChannelMessageSend(m.ChannelID, *arg+" is not a recognized option with in !utility.  Type !help -utility for a list of supported options.")
			return err
		} else {
			// Run command //
			return cmd(s, m)
		}
	}
}

// TODO - Implement DiceRole
func diceRole(s *discordgo.Session, m *discordgo.Message) error {
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
