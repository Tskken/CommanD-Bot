package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tsukinai/CommanD-Bot/botErrors"
	"github.com/tsukinai/CommanD-Bot/servers"
	"github.com/tsukinai/CommanD-Bot/utility"
)

// Command info struct //
type cmdInfo struct {
	id   string   // name
	args []string // arguments
	info string   // info about command
}

// Wrapper to create a new command in help map //
func NewCmd(id, info string, args []string) error {
	// Creates new cmdInfo //
	c := newCmdInfo(id, info, args)
	// Add struct to map //
	return addToMap(c)
}

// Creates new cmdInfo struct with given information //
func newCmdInfo(id, info string, args []string) *cmdInfo {
	return &cmdInfo{id, args, info}
}

// Adds given struct to helpMap //
func addToMap(c *cmdInfo) error {
	// Sets map key to command name and value to cmdInfo struct //
	helpMap[c.id] = *c
	// Checks if command was added correctly //
	// Returns an error if it is not
	if _, ok := helpMap[c.id]; !ok {
		return botErrors.NewError(c.id+" was not added to the map correctly.", "help_commands.go")
	} else {
		return nil
	}
}

// Temporary loadHelp map function to load the helpMap on startup.  Will be changed later to read from a file //
func loadHelp() {
	args := []string{"!player", "!message"}
	NewCmd("!help", "All possible main wrapper commands.  Type !help -<command name> to see info on each sub set of commands.", args)

	// !player commands
	args = []string{"-kick", "-k", "-ban", "-b", "-bantime", "-bt"}
	NewCmd("-player", "!player or !pl is the main interface to issue commands for the users with in your server.", args)
	args = []string{"<player name>"}
	NewCmd("-kick", "-kick is a sub command of !player and will kick the given user from the server.  You must type !player before -kick or -k to issue the command.  This can only be issued by admins of the server.", args)
	args = []string{"<player name>"}
	NewCmd("-ban", "-ban is a sub command of !player and will ban the given user from the server for the given amount of time.  Default time is 30 days.  You must type !player before -ban or -b to issue the command.  This can only be issued by admins of the server.", args)
	args = []string{"<number of days>"}
	NewCmd("-bantime", "-bantime sets the ban time when you call -ban on a user.  The default if this is not set is 30 days.", args)

	// !message commands
	args = []string{"-delete", "-del"}
	NewCmd("-message", "!message or !ms is the main interfac eto issue commands to manipulate messages with in chat.", args)
	args = []string{" ", "<given number to delete>", "<player name to delete messages from>"}
	NewCmd("-delete", "-delete deletes messages with in the channel the command is called.  The options differ based on user permissions.", args)
}

// Print the info of a command //
func printMap(s *discordgo.Session, c *discordgo.Channel, cmd *cmdInfo) error {
	outputString := cmd.info + " \n**Possible Arguments:** " + utility.ToString(cmd.args, ", ")
	_, err := s.ChannelMessageSend(c.ID, outputString)
	return err
}

// Wrapper for calling help command info //
func Help(s *discordgo.Session, m *discordgo.Message) error {
	// Parce message content on a space //
	args := utility.ParceInput(m.Content)

	// Check args length for which case to run //
	switch len(args) {
	// Length of 1: !help called with out an argument //
	case 1:
		// Make sure first argument is lowercase //
		// Returns an error if err is not nil
		if cmd, err := utility.ToLower(args, 0); err != nil {
			return err
		} else {
			// Get info about given argument and print to channel //
			return getHelp(s, m, *cmd)
		}
		// Length of 2: !help called with an argument //
	case 2:
		// Make sure second argument is lowercase //
		// Returns an error if err is not nil
		if cmd, err := utility.ToLower(args, 1); err != nil {
			return err
		} else {
			// Get info about given argument and print to channel //
			return getHelp(s, m, *cmd)
		}
	default:
		// Length of args is either 0 or greater then 2 which is not any understood function //
		// Created and returns an error
		return botErrors.NewError("Switch case in Help did not complete correctly.", "help_commands.go")
	}
}

// Get the given command and print it to the channel //
func getHelp(s *discordgo.Session, m *discordgo.Message, cmd string) error {
	// Get struct info on given argument //
	// Print an error if the given argument does not exist with in the map
	if cmdInfo, ok := helpMap[cmd]; !ok {
		errorText := cmd + " does not exist with in my understood options.  Type !help for a list of all commands you can do."
		_, err := s.ChannelMessageSend(m.ChannelID, errorText)
		return err
	} else {
		// Gets the channel the original command call was sent in //
		// Returns an error if err is not nil
		if chn, err := servers.GetChannel(s, m); err != nil {
			return err
		} else {
			// Prints command info to the channel //
			return printMap(s, chn, &cmdInfo)
		}
	}
}
