package core

import (
	"github.com/Tskana/CommanD-Bot/messagecommands"
	"github.com/bwmarrin/discordgo"
	"reflect"
)

type Commander interface {
	Init(command *Command)Commander
	Run()error
}

type Command struct {
	*discordgo.Session
	*discordgo.Message
	*ParsedCommand
}

func NewCommand(session *discordgo.Session, message *discordgo.Message, command *ParsedCommand) *Command {
	return &Command{
		session,
		message,
		command,
	}
}

type Commands struct {
	MessageCommands Commander `!message:"MessageCommands" !ms:"MessageCommands"`
}

var RootCommands *Commands

func InitCommands() {
	RootCommands = &Commands{
		MessageCommands:new(messagecommands.MessageCommand),
	}
}

func (c *Command) Run() {
	t := reflect.TypeOf(RootCommands)
	for i := 0; i < t.NumField(); i++ {
		if value, ok := t.Field(i).Tag.Lookup(c.Command); ok {
			switch value {
			case "MessageCommands":
				RootCommands.MessageCommands.Init(c).Run()
			default:
			}
		}
	}
}




