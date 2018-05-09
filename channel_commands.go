package CommanD_Bot

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"strconv"
	"strings"
)

func loadChannelCommand() *commands {
	c := commands{}
	c.commandInfo = loadChannelCommandInfo()
	c.subCommands = make(map[string]func(*discordgo.Session, *discordgo.Message) error)
	c.subCommands["-new"] = createChannel
	c.subCommands["-c"] = createChannel
	c.subCommands["-delete"] = deleteChannel
	c.subCommands["-del"] = deleteChannel

	return &c
}

// Create CommandInfo struct data //
func loadChannelCommandInfo() *commandInfo {
	m := commandInfo{}
	m.detail = "**!channel** or **!ch** : All commands that pertain to manipulating text and voice channels with in a server."
	m.commands = make(map[string]string)
	m.commands["-new"] = "**-new** or **-c**.\n" +
		"**Info**: Creates a new channel in the guild.\n" +
		"**Arguments:**\n" +
		"    **<channel name>**: Creates a text channel by default with given name.\n" +
		"    **<Channel Name><Channel Type>:** Creates a channel with given name and type (text or voice).\n"
	m.commands["-delete"] = "**-delete** or **-del**: Deletes a channel from the guild.\n" +
		"**Arguments:**\n" +
		"    **<channel name>**: Deletes the channel with given name.\n"
	return &m
}

// Create new channel function //
// - Returns an error (nil if non)
func createChannel(session *discordgo.Session, message *discordgo.Message) error {
	ok, err := isAdmin(session, message)
	if err != nil {
		return err
	}
	if !ok {
		_, err := session.ChannelMessageSend(message.ChannelID, "You do not have permission to do that.")
		return err
	}

	// Parce message on a space //
	args := strings.Fields(message.Content)

	guild, err := getGuild(session, message)
	if err != nil {
		return err
	}

	// Check the number of args passed //
	// - 3 = Create a channel with out giving a type (default text)
	// - 4 = Create a channel with a given type
	switch len(args) {
	case 3:
		// Create channel with default type //
		return newChannel(session, guild, args[2], "text")
	case 4:
		// Create channel with a given type //
		return newChannel(session, guild, args[2], args[3])
	default:
		// Error if the number of arguments is anything above 4 or below 3 //
		return errors.New("length of args was not correct. Length: " + strconv.Itoa(len(args)))
	}
	return nil
}

// Create channel with given name an type //
// - Returns an error (nil if non)
func newChannel(session *discordgo.Session, guild *discordgo.Guild, name, cType string) error {
	// Create channel with given name and type in guild //
	// - Returns an error if err is not nil
	if _, err := session.GuildChannelCreate(guild.ID, name, cType); err != nil {
		return err
	}

	return nil
}

// Delete a channel function //
// - Returns an error (nil if non)
func deleteChannel(session *discordgo.Session, message *discordgo.Message) error {
	ok, err := isAdmin(session, message)
	if err != nil {
		return err
	}
	if !ok {
		_, err := session.ChannelMessageSend(message.ChannelID, "You do not have permission to do that.")
		return err
	}

	// Parce massage on a space //
	args := strings.Fields(message.Content)

	guild, err := getGuild(session, message)
	if err != nil {
		return err
	}

	// Check length of message //
	switch len(args) {
	case 3:
		// Get channel to delete //
		if c, err := getChannelToDel(session, guild, args[2]); err != nil {
			return err
		} else {
			_, err := session.ChannelDelete(c.ID)
			return err
		}

	default:
		return errors.New("length of arguments passed to delete message is " + strconv.Itoa(len(args)))
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
	return nil, errors.New("could not find channel to delete")
}
