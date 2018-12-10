package core

import (
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

var Commands = make(map[string]Commander)

func AddCommand(lCommand, sCommand string, command Commander) {
	Commands[lCommand] = command
	Commands[sCommand] = command
}

func (c *Command) Run() error {
	if command, ok := Commands[c.Command]; !ok {
		return NewError("Command Run()", "command does not exist in map")
	} else {
		return command.Init(c).Run()
	}
}




