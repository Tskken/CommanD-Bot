package CommanD_Bot

import (
	"github.com/bwmarrin/discordgo"
)

// TODO - comment
func PlayerCommands(s *discordgo.Session, m *discordgo.Message, admin bool) error {
	arg := ToLower(ParceInput(m.Content), 1)

	switch *arg {
	case "-kick":
		return KickMember(s, m, admin)
	case "-k":
		return KickMember(s, m, admin)
	case "-ban":
		return BanMember(s, m, admin)
	case "-b":
		return BanMember(s, m, admin)
	default:
		_, err := s.ChannelMessageSend(m.ChannelID, *arg + " is not a recognized option with in !player.  Type !help -player for a list of supported options.")
		return err
	}

	return nil
}

// TODO - Comment
// TODO - Fix kick with reason functionality
func KickMember(s *discordgo.Session, m *discordgo.Message, admin bool) error {
	if admin == false {
		_, err := s.ChannelMessageSend(m.ChannelID, "You do not have the permission to kick someone.")
		return err
	}

	channel, err := GetChannel(s, m)
	if err != nil {
		return err
	}
	guild, err := GetGuild(s, channel)
	if err != nil {
		return err
	}

	args := ParceInput(m.Content)

	if len(args) == 3 {
		for _, member := range guild.Members {
			if member.Nick == args[2] || member.User.Username == args[2] {
				return s.GuildMemberDelete(guild.ID, member.User.ID)
			}
		}
	} else if len(args) == 4 {
		for _, member := range guild.Members {
			if member.Nick == args[2] {
				return s.GuildMemberDeleteWithReason(guild.ID, member.User.ID, args[3])
			}
		}
	} else {
		_, err := s.ChannelMessageSend(m.ChannelID, "Incorect arguments with in -kick call. !help -kick for more info on the option.")
		return err
	}

	return nil
}

// TODO - Implement BanMember
func BanMember(s *discordgo.Session, m *discordgo.Message, admin bool)error {
	if admin == false {
		_, err := s.ChannelMessageSend(m.ChannelID, "You do not have the permission to ban some one.")
		return err
	}

	return nil
}
