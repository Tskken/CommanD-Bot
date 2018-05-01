package CommanD_Bot

import (
	"github.com/bwmarrin/discordgo"
)

/*
TODO - Comment
*/

var botCommands map[string]*commands

type command interface {
	command(s *discordgo.Session, m *discordgo.Message) error
}

type commands struct {
	CommandInfo *CommandInfo
	SubCommands map[string]func(*discordgo.Session, *discordgo.Message) error
}

func (c *commands) command(s *discordgo.Session, m *discordgo.Message) error {
	args := ParceInput(m.Content)
	if len(args) < 2 {
		_, err := s.ChannelMessageSend(m.ChannelID, c.CommandInfo.Help())
		return err
	} else {
		if cmd, ok := c.SubCommands[args[1]]; !ok {
			return NewError("Could not understand given command", "command.go")
		} else {
			return cmd(s, m)
		}
	}
}

func RunCommand(s *discordgo.Session, m *discordgo.Message, c command) error {
	return c.command(s, m)
}

func loadCommands() {
	botCommands = make(map[string]*commands)

	messageCommands := loadMessageCommand()
	channelCommands := loadChannelCommand()
	playerCommands := loadPlayerCommand()
	utilityCommands := loadUtilityCommand()
	helpCommands := loadHelpCommand()

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
}
