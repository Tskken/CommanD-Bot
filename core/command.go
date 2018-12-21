package core

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

// TODO: Over hall command handler... round 2
type Commander interface {
	Init(command *Command) map[string]HandlerFunction
}

type Command struct {
	*Root
	*ParsedCommand
}

type HandlerFunction func() error

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

func AddCommand(command Commander, keys ...string) {
	for _, k := range keys {
		Commands[k] = command
	}
}

func (c *Command) Run() error {
	defer func() {
		err := c.DeleteMessages(c.ID)
		if err != nil {
			log.Println(err)
		}
	}()

	if command, ok := Commands[c.Command]; !ok {
		return NewError("Command Run()", "command does not exist in map")
	} else {
		if fnc, ok := command.Init(c)[c.Option]; !ok {
			return NewError("Run()", "unknown option given")
		} else {
			return fnc()
		}
	}
}
