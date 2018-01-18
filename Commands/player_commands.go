package Commands

import "github.com/bwmarrin/discordgo"

func PlayerCommands(s *discordgo.Session, m *discordgo.Message, admin bool) error {
	return nil
}

// TODO - Implement KickMember
func KickMember(s *discordgo.Session, m *discordgo.Message, admin bool)error {
	if admin == false {
		_, err := s.ChannelMessageSend(m.ChannelID, "You do not have the permission to kick someone.")
		return err
	}

	return nil
}

// TODO - Implement BanMember
func BanMember(s *discordgo.Session, m *discordgo.Message, admin bool)error {
	if admin == false {
		_, err := s.ChannelMessageSend(m.ChannelID, "You do not have the permission to kick someone.")
		return err
	}

	return nil
}
