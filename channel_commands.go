package CommanD_Bot

import "github.com/bwmarrin/discordgo"

// TODO - Comment
func ChannelCommands(s *discordgo.Session, m *discordgo.Message, admin bool) error {
	arg := ToLower(ParceInput(m.Content),1)

	switch *arg {
	case "-new":
		return CreateChannel(s, m, admin)
	case "-n":
		return CreateChannel(s, m, admin)
	case "-delete":
		return DeleteChannel(s, m, admin)
	case "-d":
		return DeleteChannel(s, m, admin)
	default:
		_, err := s.ChannelMessageSend(m.ChannelID, *arg + " is not a recognized option with in !channel.  Type !help -channel for a list of supported options.")
		return err
	}
}

// TODO - Implement CreateChannel
func CreateChannel(s *discordgo.Session, m *discordgo.Message, admin bool) error {
	return nil
}

// TODO - Implement DeleteChannel
func DeleteChannel(s *discordgo.Session, m *discordgo.Message, admin bool)error {
	return nil
}