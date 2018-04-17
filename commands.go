package CommanD_Bot

import (
	"github.com/bwmarrin/discordgo"
)

/*
 - TODO : Refactor command structure with interfaces
*/

var botCommands map[string]interface{}

type command interface {
	command(s *discordgo.Session, m *discordgo.Message) error
}

type commands struct {
	commandInfo *CommandInfo
	subCommands map[string]func(*discordgo.Session, *discordgo.Message) error
}

func RunCommand(s *discordgo.Session, m *discordgo.Message, c command) error {
	return c.command(s, m)
}

func LoadCommands() map[string]interface{} {
	cmd := make(map[string]interface{})

	MessageCommands := LoadMessageCommand()
	ChannelCommands := LoadChannelCommand()
	PlayerCommands := LoadPlayerCommand()
	UtilityCommands := LoadUtilityCommand()

	cmd["!message"] = MessageCommands
	cmd["!ms"] = MessageCommands
	cmd["!channel"] = ChannelCommands
	cmd["!ch"] = ChannelCommands
	cmd["!player"] = PlayerCommands
	cmd["!pl"] = PlayerCommands
	cmd["!utility"] = UtilityCommands
	cmd["!util"] = UtilityCommands
	return cmd
}
