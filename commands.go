package CommanD_Bot

import (
	"github.com/bwmarrin/discordgo"
)

var botCommands BotCommands

type Commands map[CommandKey]CommandAction

type RootCommand struct {
	session *discordgo.Session
	message *discordgo.Message
	keys []CommandKey
	args []string
}

type BotCommands struct {
	commands Commands
}

func Run(command RootCommand) error {
	return botCommands.commands[command.keys[0]].RunCommand(command)
}

func StartUp(){
	botCommands = BotCommands{}

	mc := loadMessageCommand()
	botCommands.commands["!messages"] = mc
	botCommands.commands["!ms"] = mc

	cc := loadChannelCommand()
	botCommands.commands["!channel"] = cc
	botCommands.commands["!ch"] = cc

	gc := loadGuildCommand()
	botCommands.commands["!guild"] = gc
	botCommands.commands["!gl"] = gc

	pc := loadPlayerCommand()
	botCommands.commands["!player"] = pc
	botCommands.commands["!pl"] = pc

	uc := loadUtilityCommand()
	botCommands.commands["!utility"] = uc
	botCommands.commands["!ut"] = uc
}