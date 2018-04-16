package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tsukinai/CommanD-Bot/botErrors"
	"github.com/tsukinai/CommanD-Bot/servers"
	"github.com/tsukinai/CommanD-Bot/utility"
)

// Create new channel function //
// - Returns an error (nil if non)
func createChannel(s *discordgo.Session, m *discordgo.Message) error {
	// Parce message on a space //
	args := utility.ParceInput(m.Content)

	// Check the number of args passed //
	// - 3 = Create a channel with out giving a type (default text)
	// - 4 = Create a channel with a given type
	switch len(args) {
	case 3:
		// Create channel with default type //
		newChannel(s, m, args[2], "text")
	case 4:
		// Create channel with a given type //
		newChannel(s, m, args[2], args[3])
	default:
		// Error if the number of arguments is anything above 4 or below 3 //
		return botErrors.NewError("Length of args was not correct. Length: "+utility.IntToStr(len(args)), "channel_commands.go")
	}
	return nil
}

// Create channel with given name an type //
// - Returns an error (nil if non)
func newChannel(s *discordgo.Session, m *discordgo.Message, name, cType string) error {
	// Get guild to create channel in //
	// - Returns an error if err is not nil
	if g, err := servers.GetGuild(s, m); err != nil {
		return err
	} else {
		// Create channel with given name and type in guild //
		// - Returns an error if err is not nil
		if _, err := s.GuildChannelCreate(g.ID, name, cType); err != nil {
			return err
		}
	}
	return nil
}

// Delete a channel function //
// - Returns an error (nil if non)
func deleteChannel(s *discordgo.Session, m *discordgo.Message) error {
	// Parce massage on a space //
	args := utility.ParceInput(m.Content)

	// Check length of message //
	switch len(args) {
	case 3:
		// Get channel to delete //
		c, _ := getChannelToDel(s, m, args[2])
		s.ChannelDelete(c.ID)
	default:
		return botErrors.NewError("Length of arguments passed to delete message is "+utility.IntToStr(len(args)), "channel_commands.go")
	}

	return nil
}

// Get the channel to delete //
// - Returns a reference to a channel and an error (nil if non)
func getChannelToDel(s *discordgo.Session, m *discordgo.Message, name string) (*discordgo.Channel, error) {
	// Gets guild to delete channel with in //
	// - Returns an error if err is not nil
	if g, err := servers.GetGuild(s, m); err != nil {
		return nil, err
	} else {
		// Gets all channels with in the guild //
		// - Returns an error if err is not nil
		if chs, err := s.GuildChannels(g.ID); err != nil {
			return nil, err
		} else {
			// Check list of channels for given name to delete //
			for _, c := range chs {
				// Return channel if channel name = given name //
				if c.Name == name {
					return c, nil
				}
			}
		}
	}
	// Something broke if this ever happens ... //
	return nil, botErrors.NewError("GetChannelToDel failed", "channel_commands.go")
}
