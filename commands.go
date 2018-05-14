package CommanD_Bot

import (
	"errors"
	"github.com/bwmarrin/discordgo"
)

// Master list of commands structures //
var botCommands map[string]*commands

// Command interface //
type command interface {
	command(session *discordgo.Session, message *discordgo.Message, args []string) error
}

// Commands structure for holding all sub commands and command info //
type commands struct {
	commandInfo *commandInfo
	subCommands map[string]func(*discordgo.Session, *discordgo.Message, []string) error
}

// command interface implementation for commands structure //
// - returns an error (nil if non)
func (c *commands) command(session *discordgo.Session, message *discordgo.Message, args []string) error {
	// Sends help info if no sub command is given //
	// - returns an error if err is not nil
	if len(args) < 2 {
		_, err := session.ChannelMessageSend(message.ChannelID, c.commandInfo.help())
		return err
	} else {
		// get sub command //
		// - returns an error if sub command does not exist in list //
		if cmd, ok := c.subCommands[args[1]]; !ok {
			return errors.New("could not understand given command")
		} else {
			// Run sub command //
			// - returns an error (nil if non)
			return cmd(session, message, args)
		}
	}
}

// Run interface command //
// - returns an error (nil if non)
func runCommand(session *discordgo.Session, message *discordgo.Message, c command, args []string) error {
	// Run command for given command structure //
	// - returns an error (nil if non)
	return c.command(session, message, args)
}

// Set all commands for command structures with in master command list //
func loadCommands() {
	// Create master command map //
	botCommands = make(map[string]*commands)

	// Load command structure for message commands //
	messageCommands := loadMessageCommand()

	// Load command structure for channel commands //
	channelCommands := loadChannelCommand()

	// Load command structure for player commands //
	playerCommands := loadPlayerCommand()

	// Load command structure for utility commands //
	utilityCommands := loadUtilityCommand()

	// Load command structure for help commands //
	helpCommands := loadHelpCommand()

	// Load command structure for guild commands //
	guildCommands := loadGuildCommand()

	// Set message commands to master command map //
	botCommands["!message"] = messageCommands
	botCommands["!ms"] = messageCommands

	// Set channel commands to master command map //
	botCommands["!channel"] = channelCommands
	botCommands["!ch"] = channelCommands

	// Set player commands to master command map //
	botCommands["!player"] = playerCommands
	botCommands["!pl"] = playerCommands

	// Set utility commands to master command map //
	botCommands["!utility"] = utilityCommands
	botCommands["!util"] = utilityCommands

	// Set help commands to master command map //
	botCommands["!help"] = helpCommands
	botCommands["!h"] = helpCommands

	// Set guild commands to master command map //
	botCommands["!guild"] = guildCommands
	botCommands["!g"] = guildCommands
}
