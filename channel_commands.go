package CommanD_Bot

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"strconv"
)

type ChannelCommands struct {
	commands map[string]func(*Root)error
}

func (m *ChannelCommands) RunCommand(root *Root) error {
	return m.commands[root.CommandType()](root)
}

func LoadChannelCommand() *ChannelCommands {
	// Creates channel command struct //
	c := ChannelCommands{
		make(map[string]func(*Root)error),
	}

	// Create Channel function //
	c.commands["-new"] = CreateChannel
	c.commands["-c"] = CreateChannel

	// Delete Channel function //
	c.commands["-delete"] = DeleteChannel
	c.commands["-del"] = DeleteChannel

	// Return reference to Channel commands instance //
	return &c
}

// Create Channel command help data //
func loadChannelCommandInfo() *commandInfo {
	// Create new commandInfo struct //
	c := commandInfo{}
	// Channel command default help info //
	c.detail = "**!channel** or **!ch** : All commands that pertain to manipulating text and voice channels with in a server."

	// Create sub command info map //
	c.commands = make(map[string]string)

	// Create Create Channel help info //
	c.commands["-new"] = "**-new** or **-c**.\n" +
		"**Info**: Creates a new channel in the guild.\n" +
		"**Arguments:**\n" +
		"    **<channel name>**: Creates a text channel by default with given name.\n" +
		"    **<Channel Name><Channel Type>:** Creates a channel with given name and type (text or voice).\n"

	// Create Delete Channel help info //
	c.commands["-delete"] = "**-delete** or **-del**: Deletes a channel from the guild.\n" +
		"**Arguments:**\n" +
		"    **<channel name>**: Deletes the channel with given name.\n"

	// Return reference to Channel help info instance //
	return &c
}

// Create new channel function //
// - Returns an error (nil if non)
func CreateChannel(root *Root) error {
	// Check if user is admin //
	// - Returns if user is not an admin
	// - Returns an error if err is not nil
	if ok, err := root.IsAdmin(); err != nil {
		return err
	} else if !ok {
		return root.MessageSend("You do not have permission to do that.")
	}

	// Check the number of args //
	// - 3 = Create a channel with out giving a type (default text)
	// - 4 = Create a channel with a given type
	switch len(root.CommandArgs()) {
	case 1:
		// Create channel with default type //
		return root.NewChannel(root.CommandArgs()[0], "text")
	case 2:
		// Create channel with a given type //
		return root.NewChannel(root.CommandArgs()[0], root.CommandArgs()[1])
	default:
		// Error if the number of arguments is anything above 4 or below 3 //
		return errors.New("length of args was not correct. Length: " + strconv.Itoa(len(root.CommandArgs())))
	}
	return nil


}

// Create channel with given name an type //
// - Returns an error (nil if non)
func (r *Root) NewChannel(name, cType string) error {
	// Get guild to add channel to //
	// - Returns an error if err is not nil
	if guild, err := r.GetGuild(); err != nil {
		return err
	} else {
		// Check to make sure the type was correct //
		// - returns and error if it was not
		if cType != "text" && cType != "voice" {
			return errors.New("channel type was not ether text or voice")
		}

		// Create channel with given name and type in guild //
		// - Returns an error if err is not nil
		if _, err := r.GuildChannelCreate(guild.ID, name, cType); err != nil {
			return err
		}
	}

	return nil
}

// Delete a channel function //
// - Returns an error (nil if non)
func DeleteChannel(root *Root) error {
	// Check if user is an admin //
	// - returns an error if err is not nil
	// - returns if user is not an admin
	if ok, err := root.IsAdmin(); err != nil {
		return err
	} else if !ok {
		return root.MessageSend( "You do not have permission to do that.")
	}

	// Check length of message //
	// - returns an error if if args does not have a channel name
	switch len(root.CommandArgs()) {
	// channel name was given //
	case 1:
		// Get channel to delete //
		// - returns an error if err is not nil
		if c, err := root.GetChannelToDel(root.CommandArgs()[0]); err != nil {
			return err
		} else {
			// Delete channel //
			// - returns error (nil if non)
			_, err := root.ChannelDelete(c.ID)
			return err
		}
		// channel name was not given //
	default:
		// returns an error for channel name not being given //
		return errors.New("channel name to delete was not given")
	}
}

// Get the channel to delete //
// - Returns a reference to a channel and an error (nil if non)
func (r *Root) GetChannelToDel(name string) (*discordgo.Channel, error) {
	// Get guild channel the channel exists in //
	// - returns an error if err is not nil
	if guild, err := r.GetGuild(); err != nil {
		return nil, err
	} else {
		// Gets all channels with in the guild //
		// - Returns an error if err is not nil
		if chs, err := r.GuildChannels(guild.ID); err != nil {
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

	// Return error if channel was not found in guild //
	return nil, errors.New("could not find channel to delete")
}
