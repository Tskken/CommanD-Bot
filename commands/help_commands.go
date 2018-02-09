package commands

import (
	"github.com/bwmarrin/discordgo"
	"strings"
	"github.com/tsukinai/CommanD-Bot/utility"
	"github.com/tsukinai/CommanD-Bot/botErrors"
	"github.com/tsukinai/CommanD-Bot/servers"
)

// TODO - Implement



type cmdInfo struct {
	id   string
	args []string
	info string
}

func NewCmd(id, info string, args []string) error {
	c := newCmdInfo(id, info, args)
	return addToMap(c)
}

func newCmdInfo(id, info string, args []string) *cmdInfo {
	c := cmdInfo{id, args, info}
	return &c
}

func addToMap(c *cmdInfo) error {
	HelpMap[c.id] = *c
	_, ok := HelpMap[c.id]
	if ok != true {
		return botErrors.NewError(c.id+" was not added to the map correctly.", "help_commands.go")
	}
	return nil
}

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

func printMap(s *discordgo.Session, c *discordgo.Channel, cmd *cmdInfo) error {
	outputString := cmd.info + " \n**Possible Arguments:** " + strings.Join(cmd.args, ", ")
	_, err := s.ChannelMessageSend(c.ID, outputString)
	return err
}

/*
func (c *cmdInfo) setID(id string){}
func (c *cmdInfo) getId()string{return c.id}

func (c *cmdInfo) setArgs(args []string){}
func (c *cmdInfo) getArgs()[]string{return c.args}

func (c *cmdInfo) setInfo(info string){}
func (c *cmdInfo) getInfo()string{return c.info}
*/

func Help(s *discordgo.Session, m *discordgo.Message, admin bool) error {
	cmd, err := utility.ToLower(utility.ParceInput(m.Content), 1)
	if err != nil {
		cmd, err = utility.ToLower(utility.ParceInput(m.Content), 0)
		if err != nil {
			return err
		}
	}
	cmdInfo, ok := HelpMap[*cmd]
	if ok != true {
		errorText := *cmd + " does not exist with in my understood options.  Type !help for a list of all commands you can do."
		_, err = s.ChannelMessageSend(m.ChannelID, errorText)
		if err != nil {
			return err
		}
	}
	chn, err := servers.GetChannel(s, m)
	if err != nil {
		return err
	}
	return printMap(s, chn, &cmdInfo)
}
