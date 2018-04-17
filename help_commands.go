package CommanD_Bot

//import "github.com/bwmarrin/discordgo"

// Command info struct //
type CommandInfo struct {
	detail   string            // Details about command
	commands map[string]string // Map of sub commands for command
}

// Get help for a command //
// - Returns the detail information and list of sub commands for command
func (mc *CommandInfo) Help() string {
	// Add detail info to output //
	output := "**Command**: " + mc.detail + "\n"
	output += "**Sub-Commands**: \n"
	// Add each sub command to output //
	for _, ms := range mc.commands {
		output += "	" + ms + "\n"
	}

	// Return output //
	return output
}

// Get help for a sub-command //
// - Returns the info about the sub command
func (mc *CommandInfo) HelpCommand(arg string) string {
	return mc.commands[arg]
}

type HelpCommands struct {

}

/*
type HelpCommand Commands

func (mc *HelpCommand) Command(s *discordgo.Session, m *discordgo.Message) error {
	args := ParceInput(m.Content)
	if len(args) < 2 {
		_, err := s.ChannelMessageSend(m.ChannelID, mc.commandInfo.Help())
		return err
	} else {
		return mc.subCommands[args[1]](s, m)
	}
}

func LoadHelpCommand() *MessageCommand {
	m := MessageCommand{}
	m.commandInfo = loadHelpCommandInfo()
	m.subCommands = make(map[string]func(*discordgo.Session, *discordgo.Message) error)
	/*m.subCommands["-delete"] = deleteMessage
	m.subCommands["-del"] = deleteMessage
	m.subCommands["-clear"] = clearMessages
	m.subCommands["-cl"] = clearMessages

	return &m
}

// Create CommandInfo struct data //
func loadHelpCommandInfo() *CommandInfo {
	m := &CommandInfo{}
	m.detail =
	m.commands = make(map[string]string)
	m.commands["-delete"] = "**-delete** or **-del**.\n**Info**: Deletes a given messages with in a server.\n" +
		"**Arguments:**\n		**none**: Deletes the last messages in the chanel by you.\n		**<Number between 1 - 99>:**" +
		" Deletes the last given number of messages by you. \n		**<User Name>:** Deletes the last message by that user (only can be used by admin)." +
		"\n		**<Number between 1 - 99> <User Name>:** Deletes the last given number of messages by the given user (only can be used by admin)."
	return m
}
*/