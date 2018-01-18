package CommanD_Bot

import "github.com/bwmarrin/discordgo"

// TODO - Comment
// TODO - Implement functions
func UtilityCommands(s *discordgo.Session, m *discordgo.Message, admin bool) error {
	arg := ToLower(ParceInput(m.ChannelID), 1)

	switch *arg{
	default:
		_, err := s.ChannelMessageSend(m.ChannelID, *arg + "Is not a recognized option with in !util.  !help -util for a list of options that are supported.")
		return err
	}

	return nil
}
