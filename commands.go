package CommanD_Bot

/*
Last Updated: 11/20/27
Author: Dylan Blanchard

commands.go

Command function implementation
*/

import (
	// Golang imports //

	// External imports //
	"github.com/bwmarrin/discordgo"
)

/*
type commandInfo struct {
	name string
	info []string
	args []string
	perm byte
}*/

// Bot Command Dictionary //
var BotCommands = make(map[string]func(*discordgo.Session, *discordgo.Message, bool) error)

//var HelpCommands = make(map[string][]commandInfo)

// Load command maps with bot commands //
// - BotCommands loaded with all commands
// - commandHelp loaded with all info on commands
func Load() {
	// Load all commands in to botCommands map //
	//loadKeyMap()
	loadBotCommands()
	//loadHelpCommands()
}

// Loads commands in to botCommands map //
// TODO - Add commands to BotCommands map
func loadBotCommands() {
	// Bot Commands dictionary //
	// - Runs all command functions given string key as command name
	BotCommands["!message"] = MessageCommands
	BotCommands["!ms"] = MessageCommands
	BotCommands["!player"] = PlayerCommands
	BotCommands["!pl"] = PlayerCommands
	BotCommands["!channel"] = ChannelCommands
	BotCommands["!ch"] = ChannelCommands
	BotCommands["!utility"] = UtilityCommands
	BotCommands["!util"] = UtilityCommands
	/*
	BotCommands[1] = Help
	BotCommands[2] = DeleteMessage
	BotCommands[3] = CreateChannel
	BotCommands[4] = DeleteChannel
	BotCommands[5] = KickMember
	BotCommands[6] = BanMember
	*/
}

/*
// TODO - Refactor Help structure
func loadHelpCommands(){
	info := make([]commandInfo, 0)

	cmd := commandInfo{}
	cmd.name = "!delete"
	cmd.info = append(cmd.info, "Delete last message sent in channel")
	cmd.info = append(cmd.info, "Delete last message you sent in channel")
	cmd.perm = 2

	info = append(info, cmd)

	cmd.info = nil
	cmd.info = append(cmd.info, "Delete the last given number of messages")
	cmd.info = append(cmd.info, "Delete the last given number of messages sent by you")
	cmd.args = append(cmd.args, "<number to delete>")

	info = append(info, cmd)


	HelpCommands["!delete"] = info
}*/
