package CommanD_Bot

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

type Commands map[string]CommandAction

type BotCommands struct {
	commands Commands
}

func (r *Root) Run() error {
	return r.commands[r.CommandRoot()].RunCommand(r)
}

func (r *Root) Args() []string {
	return ToLower(strings.Fields(r.Content))
}

func (r *Root) CommandArgs() []string{
	// Parce input in to arguments and set them to lower case //
	return r.Args()[2:]
}

func (r *Root) CommandRoot()string {
	return r.Args()[0]
}

func (r *Root) CommandType()string {
	return r.Args()[1]
}

func (r *Root) MessageSend(content string) error {
	_, err := r.ChannelMessageSend(r.ChannelID, content)
	return err
}

