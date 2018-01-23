package CommanD_Bot

import (
	"github.com/bwmarrin/discordgo"
	"strconv"
)

// TODO - comment
func PlayerCommands(s *discordgo.Session, m *discordgo.Message, admin bool) error {
	arg := ToLower(ParceInput(m.Content), 1)

	channel, err := GetChannel(s, m)
	if err != nil {
		return err
	}
	guild, err := GetGuild(s, channel)
	if err != nil {
		return err
	}

	switch *arg {
	case "-kick":
		if channel.Name != "terminal" {
			_, err := s.ChannelMessageSend(m.ChannelID, "You can not enter that command with in this channel.")
			return err
		}
		if admin != true {
			_, err := s.ChannelMessageSend(m.ChannelID, "You do not have the permission to kick someone.")
			return err
		}
		return KickMember(s, m, guild)
	case "-k":
		if channel.Name != "terminal" {
			_, err := s.ChannelMessageSend(m.ChannelID, "You can not enter that command with in this channel.")
			return err
		}
		if admin != true {
			_, err := s.ChannelMessageSend(m.ChannelID, "You do not have the permission to kick someone.")
			return err
		}
		return KickMember(s, m, guild)
	case "-ban":
		if channel.Name != "terminal" {
			_, err := s.ChannelMessageSend(m.ChannelID, "You can not enter that command with in this channel.")
			return err
		}
		if admin != true {
			_, err := s.ChannelMessageSend(m.ChannelID, "You do not have the permission to kick someone.")
			return err
		}
		return BanMember(s, m, guild)
	case "-b":
		if channel.Name != "terminal" {
			_, err := s.ChannelMessageSend(m.ChannelID, "You can not enter that command with in this channel.")
			return err
		}
		if admin != true {
			_, err := s.ChannelMessageSend(m.ChannelID, "You do not have the permission to kick someone.")
			return err
		}
		return BanMember(s, m, guild)
	case "-bantime":
		if channel.Name != "terminal" {
			_, err := s.ChannelMessageSend(m.ChannelID, "You can not enter that command with in this channel.")
			return err
		}
		if admin != true {
			_, err := s.ChannelMessageSend(m.ChannelID, "You do not have the permission to kick someone.")
			return err
		}
		return NewBanTimer(s, m, guild)
	case "-bt":
		if channel.Name != "terminal" {
			_, err := s.ChannelMessageSend(m.ChannelID, "You can not enter that command with in this channel.")
			return err
		}
		if admin != true {
			_, err := s.ChannelMessageSend(m.ChannelID, "You do not have the permission to kick someone.")
			return err
		}
		return NewBanTimer(s, m, guild)
	default:
		_, err := s.ChannelMessageSend(m.ChannelID, *arg + " is not a recognized option with in !player.  Type !help -player for a list of supported options.")
		return err
	}

	return nil
}

// TODO - Comment
// TODO - Fix kick with reason functionality
func KickMember(s *discordgo.Session, m *discordgo.Message, guild *discordgo.Guild) error {

	args := ParceInput(m.Content)

	var guildMember *discordgo.Member

	for _, member := range guild.Members {
		if member.Nick == args[2] || member.User.Username == args[2] {
			guildMember = member
		}
	}

	if len(args) == 3 {
		return s.GuildMemberDelete(guild.ID, guildMember.User.ID)
	} else if len(args) >= 4 {
		return s.GuildMemberDeleteWithReason(guild.ID, guildMember.User.ID, ToString(args[3:]))
	} else {
		_, err := s.ChannelMessageSend(m.ChannelID, "Incorect arguments with in -kick call. !help -kick for more info on the option.")
		return err
	}

	return nil
}

// TODO - Fix Ban with reason functionality
// TODO - Comment
func BanMember(s *discordgo.Session, m *discordgo.Message, guild *discordgo.Guild)error {
	args := ParceInput(m.Content)

	if len(args) == 3 {
		for _, member := range guild.Members {
			if member.Nick == args[2] || member.User.Username == args[2] {
				return s.GuildBanCreate(guild.ID, member.User.ID, BanTime[guild.Name])
			}
		}
	} else if len(args) >= 4 {
		for _, member := range guild.Members {
			if member.Nick == args[2] || member.User.Username == args[2] {
				return s.GuildBanCreateWithReason(guild.ID, member.User.ID, ToString(args[3:]), BanTime[guild.Name])
			}
		}
	} else {
		_, err := s.ChannelMessageSend(m.ChannelID, "Incorect arguments with in -ban call. !help -ban for more info on the option.")
		return err
	}

	return nil
}

// TODO - Comment
func NewBanTimer(s *discordgo.Session, m *discordgo.Message, guild *discordgo.Guild)error{
	_, err := s.ChannelMessageSend(m.ChannelID, "Ban time was " + strconv.Itoa(BanTime[guild.Name]))
	if err != nil {
		return err
	}
	args := ParceInput(m.Content)
	time, err := strconv.Atoi(args[2])
	if err != nil {
		return err
	}
	BanTime[guild.Name] = time

	_, err = s.ChannelMessageSend(m.ChannelID, "Ban time has been set to " + strconv.Itoa(BanTime[guild.Name]))
	return err
}