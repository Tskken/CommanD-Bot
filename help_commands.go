package CommanD_Bot

import "github.com/bwmarrin/discordgo"

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

func loadHelpCommandInfo() *CommandInfo {
	h := CommandInfo{}
	h.detail = "type !help and a command to get more info on each command.\n" +
		"**Commands:**\n" +
		"	**!message:** Commands for messages\n" +
		"	**!player:** Commands for Players\n" +
		"	**!channel:** Commands for channels\n" +
		"	**!utility:** Utility commands"
	return &h
}

func helpMessages(s *discordgo.Session, m *discordgo.Message) error {
	args := ToLower(ParceInput(m.Content))

	if len(args) == 2 {
		c := botCommands[args[1]]
		s.ChannelMessageSend(m.ChannelID, c.commandInfo.Help())
	} else if len(args) == 3 {
		c := botCommands[args[1]]
		s.ChannelMessageSend(m.ChannelID, c.commandInfo.HelpCommand(args[2]))
	} else {
		return NewError("Not the right number of arguments in help commnad call", "help_commands.go")
	}
	return nil
}
