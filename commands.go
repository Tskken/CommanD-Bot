package CommanD_Bot

import (
	"github.com/bwmarrin/discordgo"
)

// Bot Command Dictionary //
var BotCommands = make(map[string]func(*discordgo.Session, *discordgo.Message, bool) error)


// Loads commands in to botCommands map //
// TODO - Fix comments
func loadBotCommands() {


	/*
	encdec := EncDec{}
	encdec.OpenFile()
	encdec.NewDecGob()
	err := encdec.DecGob(&BanTime)
	if err != nil {
		log.Println(err)
	}
	encdec.CloseFile()*/

	// Bot Commands dictionary //
	// - Runs all command functions given string key as command name
	BotCommands["!message"] = MessageCommands
	BotCommands["!ms"] = MessageCommands
	BotCommands["!player"] = PlayerCommands
	BotCommands["!pl"] = PlayerCommands
	BotCommands["!channel"] = ChannelCommands
	BotCommands["!ch"] = ChannelCommands
	BotCommands["!guild"] = GuildCommands
	BotCommands["!gl"] = GuildCommands
	BotCommands["!utility"] = UtilityCommands
	BotCommands["!util"] = UtilityCommands
	BotCommands["!help"] = Help

}


