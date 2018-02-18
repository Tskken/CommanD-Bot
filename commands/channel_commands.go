package commands

import (
	"github.com/bwmarrin/discordgo"
)

/*
// Wrapper function called by BotCommands to run channelCommands //
func ChannelCommands(s *discordgo.Session, m *discordgo.Message) error {
	admin, err := servers.IsAdmin(s, m)
	if err != nil {
		return err
	}

	// Get argument passed to !channel //
	// returns error if err != nil
	if arg, err := utility.ToLower(utility.ParceInput(m.Content), 1); err != nil {
		return err
	} else {
		// Get command from channelCommands //
		// Prints errors to server if argument passed did not exist in map
		if cmd, ok := channelCommands[*arg]; !ok {
			_, err := s.ChannelMessageSend(m.ChannelID, *arg+" is not a recognized option with in !channel.  Type !help -channel for a list of supported options.")
			return err
		} else {
			// Run command //
			return cmd(s, m, admin)
		}
	}
}*/

// TODO - Implement CreateChannel
func createChannel(s *discordgo.Session, m *discordgo.Message) error { return nil }

// TODO - Implement DeleteChannel
func deleteChannel(s *discordgo.Session, m *discordgo.Message) error { return nil }
