package CommanD_Bot

/*
// Set help command structure //
func loadHelpCommand() *Commands {
	// Create help command structure //
	h := &Commands{}

	// Load help command info structure //
	h.commandInfo = loadHelpCommandInfo()

	// Create help sub command map //
	h.subCommands = make(map[string]func(*discordgo.Session, *discordgo.Message, []string) error)

	// Set message commands in help message //
	h.subCommands["!messages"] = helpMessages
	h.subCommands["!ms"] = helpMessages

	// Set player commands in help message //
	h.subCommands["!player"] = helpMessages
	h.subCommands["!pl"] = helpMessages

	// Set channel commands in help message //
	h.subCommands["!channel"] = helpMessages
	h.subCommands["!ch"] = helpMessages

	// Set utility commands in help message //
	h.subCommands["!utility"] = helpMessages
	h.subCommands["!util"] = helpMessages

	// Set guild commands in help message //
	h.subCommands["!guild"] = helpMessages
	h.subCommands["!g"] = helpMessages

	// Returns a reference to help command structure //
	return h
}

// Set help command info structure //
func loadHelpCommandInfo() *commandInfo {
	// Create help info structure //
	h := commandInfo{}

	// Set help default info //
	h.detail = "type !help and a command to get more info on each command.\n" +
		"**Commands:**\n" +
		"    **!message:** Commands for messages\n" +
		"    **!player:** Commands for Players\n" +
		"    **!channel:** Commands for channels\n" +
		"    **!utility:** Utility commands\n" +
		"    **!guild:** commands for a guild"

	// Return a reference to the help info structure //
	return &h
}

// Help command //
// - returns an error (nil if non)
func helpMessages(command *RootCommand) error {
	// Check for help info for a main command //
	if len(args) == 2 {
		// Get main command structure //
		c := botCommands[args[1]]

		// Send help to channel for command //
		// - returns an error (nil if non)
		if _, err := session.ChannelMessageSend(message.ChannelID, c.commandInfo.help()); err != nil {
			return err
		}
		return deleteMessage(session, message.ChannelID, message.ID)
	} else if len(args) == 3 {
		// Get main command structure //
		c := botCommands[args[1]]

		// Send help info for sub command in command structure //
		// - returns an error (nil if non)
		if _, err := session.ChannelMessageSend(message.ChannelID, c.commandInfo.helpCommand(args[2])); err != nil {
			return err
		}
		return deleteMessage(session, message.ChannelID, message.ID)
	} else {
		// Return an error if given incorrect arguments //
		return errors.New("not the right number of arguments in help command call")
	}
}
*/