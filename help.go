package CommanD_Bot

// Command info struct //
type commandInfo struct {
	detail   string            // Details about command
	commands map[string]string // Map of sub commands for command
}

// Get help for a command //
// - Returns the detail information and list of sub commands for command
func (mc *commandInfo) help() string {
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
func (mc *commandInfo) helpCommand(arg string) string {
	return mc.commands[arg]
}
