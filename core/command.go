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

type HandlerFunction func()error

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

func AddCommand(command Commander, keys... string) {
	for _, k := range keys {
		Commands[k] = command
	}
}

func (c *Command) Run() error {
	if command, ok := Commands[c.Command]; !ok {
		return NewError("Command Run()", "command does not exist in map")
	} else {
		return command.Init(c).Run()
	}
}




