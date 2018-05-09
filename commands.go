package CommanD_Bot

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"strings"
)

/*
TODO - Comment
*/

var botCommands map[string]*commands

type command interface {
	command(session *discordgo.Session, message *discordgo.Message) error
}

type commands struct {
	commandInfo *commandInfo
	subCommands map[string]func(*discordgo.Session, *discordgo.Message) error
}

func (c *commands) command(session *discordgo.Session, message *discordgo.Message) error {
	args := strings.Fields(message.Content)
	if len(args) < 2 {
		_, err := session.ChannelMessageSend(message.ChannelID, c.commandInfo.help())
		return err
	} else {
		if cmd, ok := c.subCommands[args[1]]; !ok {
			return errors.New("could not understand given command")
		} else {
			return cmd(session, message)
		}
	}
}

func runCommand(session *discordgo.Session, message *discordgo.Message, c command) error {
	return c.command(session, message)
}

func loadCommands() {
	botCommands = make(map[string]*commands)

	messageCommands := loadMessageCommand()
	channelCommands := loadChannelCommand()
	playerCommands := loadPlayerCommand()
	utilityCommands := loadUtilityCommand()
	helpCommands := loadHelpCommand()
	guildCommands := loadGuildCommand()

	botCommands["!message"] = messageCommands
	botCommands["!ms"] = messageCommands
	botCommands["!channel"] = channelCommands
	botCommands["!ch"] = channelCommands
	botCommands["!player"] = playerCommands
	botCommands["!pl"] = playerCommands
	botCommands["!utility"] = utilityCommands
	botCommands["!util"] = utilityCommands
	botCommands["!help"] = helpCommands
	botCommands["!h"] = helpCommands
	botCommands["!guild"] = guildCommands
	botCommands["!g"] = guildCommands
}
