package CommanD_Bot

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"strings"
)

/*
TODO - Fix help commands
*/

func loadHelpCommand() *commands {
	h := commands{}
	h.commandInfo = loadHelpCommandInfo()
	h.subCommands = make(map[string]func(*discordgo.Session, *discordgo.Message) error)
	h.subCommands["!messages"] = helpMessages
	h.subCommands["!ms"] = helpMessages
	h.subCommands["!player"] = helpMessages
	h.subCommands["!pl"] = helpMessages
	h.subCommands["!channel"] = helpMessages
	h.subCommands["!ch"] = helpMessages
	h.subCommands["!utility"] = helpMessages
	h.subCommands["!util"] = helpMessages
	return &h
}

func loadHelpCommandInfo() *commandInfo {
	h := commandInfo{}
	h.detail = "type !help and a command to get more info on each command.\n" +
		"**Commands:**\n" +
		"    **!message:** Commands for messages\n" +
		"    **!player:** Commands for Players\n" +
		"    **!channel:** Commands for channels\n" +
		"    **!utility:** Utility commands"
	return &h
}

func helpMessages(session *discordgo.Session, message *discordgo.Message) error {
	args := toLower(strings.Fields(message.Content))

	if len(args) == 2 {
		c := botCommands[args[1]]
		session.ChannelMessageSend(message.ChannelID, c.commandInfo.help())
	} else if len(args) == 3 {
		c := botCommands[args[1]]
		session.ChannelMessageSend(message.ChannelID, c.commandInfo.helpCommand(args[2]))
	} else {
		return errors.New("not the right number of arguments in help commnad call")
	}
	return nil
}
