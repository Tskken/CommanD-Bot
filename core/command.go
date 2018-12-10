package core

import (
	"github.com/Tskana/CommanD-Bot/mc"
	"github.com/bwmarrin/discordgo"
)

type Commander interface {
	Init(command *Command)Commander
	Run()error
}

type Command struct {
	*Root
	*ParsedCommand
}

func NewCommand(session *discordgo.Session, message *discordgo.Message, command *ParsedCommand) *Command {
	return &Command{
		&Root{
			session,
			message,
		},
		command,
	}
}

var (
	MessageCommand = new(mc.MessageCommand)
)

var Commands = map[string]Commander {
	"!message":MessageCommand,
	"!ms":MessageCommand,
}

func (c *Command) Run() error {
	if command, ok := Commands[c.Command]; !ok {
		return NewError("Command Run()", "command does not exist in map")
	} else {
		return command.Init(c).Run()
	}
}




