package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tsukinai/CommanD-Bot/botErrors"
	"github.com/tsukinai/CommanD-Bot/servers"
	"github.com/tsukinai/CommanD-Bot/utility"
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

// TODO - Test and refine code

// TODO - Cleanup code
func createChannel(s *discordgo.Session, m *discordgo.Message) error {
	args := utility.ParceInput(m.Content)
	switch len(args) {
	case 3:
		newChannel(s, m, args[2], "")
	case 4:
		newChannel(s, m, args[2], args[3])
	default:
		return botErrors.NewError("Length of args was not correct. Length: "+utility.IntToStr(len(args)), "channel_commands.go")
	}
	return nil
}

// TODO - Cleanup code
func newChannel(s *discordgo.Session, m *discordgo.Message, name, cType string) error {
	if g, err := servers.GetGuild(s, m); err != nil {
		return err
	} else {
		if _, err := s.GuildChannelCreate(g.ID, name, cType); err != nil {
			return err
		}
	}
	return nil
}

// TODO - Finish code
func deleteChannel(s *discordgo.Session, m *discordgo.Message) error {
	args := utility.ParceInput(m.Content)
	switch len(args) {
	case 3:
		c, _ := getChannelToDel(s, m)
		s.ChannelDelete(c.ID)
	default:

	}

	return nil
}

// TODO - Cleanup code
func getChannelToDel(s *discordgo.Session, m *discordgo.Message) (*discordgo.Channel, error) {
	g, _ := servers.GetGuild(s, m)
	args := utility.ParceInput(m.Content)
	chs, _ := s.GuildChannels(g.ID)
	for _, c := range chs {
		if c.Name == args[2] {
			return c, nil
		}
	}
	return nil, nil
}
