package CommanD_Bot

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"strconv"
)



func loadChannelCommand() *ChannelCommands {
	// Creates channel command struct //
	c := ChannelCommands{}

	// Creates sub command function map //
	c.commands = make(map[CommandKey]func(RootCommand) error)

	// Create Channel function //
	c.commands["-new"] = createChannel
	c.commands["-c"] = createChannel

	// Delete Channel function //
	c.commands["-delete"] = deleteChannel
	c.commands["-del"] = deleteChannel

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
func createChannel(command RootCommand) error {
	// Check if user is admin //
	// - Returns if user is not an admin
	// - Returns an error if err is not nil
	if ok, err := isAdmin(command.session, command.message); err != nil {
		return err
	} else if !ok {
		_, err := command.session.ChannelMessageSend(command.message.ChannelID, "You do not have permission to do that.")
		return err
	}

	// Get guild to add channel to //
	// - Returns an error if err is not nil
	if guild, err := getGuild(command.session, command.message); err != nil {
		return err
	} else {
		// Check the number of args //
		// - 3 = Create a channel with out giving a type (default text)
		// - 4 = Create a channel with a given type
		switch len(command.args) {
		case 1:
			// Create channel with default type //
			return newChannel(command.session, guild, command.args[0], "text")
		case 2:
			// Create channel with a given type //
			return newChannel(command.session, guild, command.args[0], command.args[1])
		default:
			// Error if the number of arguments is anything above 4 or below 3 //
			return errors.New("length of args was not correct. Length: " + strconv.Itoa(len(command.args)))
		}
		return nil
	}

}

// Create channel with given name an type //
// - Returns an error (nil if non)
func newChannel(session *discordgo.Session, guild *discordgo.Guild, name, cType string) error {
	// Check to make sure the type was correct //
	// - returns and error if it was not
	if cType != "text" && cType != "voice" {
		return errors.New("channel type was not ether text or voice")
	}

	// Create channel with given name and type in guild //
	// - Returns an error if err is not nil
	if _, err := session.GuildChannelCreate(guild.ID, name, cType); err != nil {
		return err
	}

	return nil
}

// Delete a channel function //
// - Returns an error (nil if non)
func deleteChannel(command RootCommand) error {
	// Check if user is an admin //
	// - returns an error if err is not nil
	// - returns if user is not an admin
	if ok, err := isAdmin(command.session, command.message); err != nil {
		return err
	} else if !ok {
		_, err := command.session.ChannelMessageSend(command.message.ChannelID, "You do not have permission to do that.")
		return err
	}

	// Get guild channel the channel exists in //
	// - returns an error if err is not nil
	if guild, err := getGuild(command.session, command.message); err != nil {
		return err
	} else {
		// Check length of message //
		// - returns an error if if args does not have a channel name
		switch len(command.args) {
		// channel name was given //
		case 1:
			// Get channel to delete //
			// - returns an error if err is not nil
			if c, err := getChannelToDel(command.session, guild, command.args[0]); err != nil {
				return err
			} else {
				// Delete channel //
				// - returns error (nil if non)
				_, err := command.session.ChannelDelete(c.ID)
				return err
			}
		// channel name was not given //
		default:
			// returns an error for channel name not being given //
			return errors.New("channel name to delete was not given")
		}
	}
}

// Get the channel to delete //
// - Returns a reference to a channel and an error (nil if non)
func getChannelToDel(session *discordgo.Session, guild *discordgo.Guild, name string) (*discordgo.Channel, error) {
	// Gets all channels with in the guild //
	// - Returns an error if err is not nil
	if chs, err := session.GuildChannels(guild.ID); err != nil {
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

	// Return error if channel was not found in guild //
	return nil, errors.New("could not find channel to delete")
}
