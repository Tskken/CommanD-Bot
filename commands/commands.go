package commands

import (
	"github.com/bwmarrin/discordgo"
)

// Bot Command Dictionary //
var BotCommands = make(map[string]func(*discordgo.Session, *discordgo.Message, bool) error)
var BanTime = make(map[string]int)

// Load command maps with bot commands //
func Load() {
	// Load all commands in to botCommands map //
	loadBotCommands()
	loadHelp()
}

func Save() {
	//saveBotMaps()
}

/*
func saveBotMaps()error{
	encdec := NewEncDec()
	err := encdec.OpenFile()
	if err != nil {
		return err
	}
	encdec.NewEncGob()
	err = encdec.EncGob(BanTime)
	if err != nil {
		return err
	}
	err = encdec.CloseFile()
	return err
}*/

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


