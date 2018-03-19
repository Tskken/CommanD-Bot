package commands

import (
	"github.com/bwmarrin/discordgo"
)

// Bot Command Dictionary //
var BotCommands = make(map[string]map[string]func(*discordgo.Session, *discordgo.Message) error)

// Sub Commands //
var messageCommands = make(map[string]func(*discordgo.Session, *discordgo.Message) error)
var playerCommands = make(map[string]func(*discordgo.Session, *discordgo.Message) error)
var channelCommands = make(map[string]func(*discordgo.Session, *discordgo.Message) error)
var guildCommands = make(map[string]func(*discordgo.Session, *discordgo.Message) error)
var utilityCommands = make(map[string]func(*discordgo.Session, *discordgo.Message) error)

// TODO - Fix helpMap
// var helpMap = make(map[string]cmdInfo)

// Utility Maps //
var banTime = make(map[string]int)

// Load command maps with bot commands //
func Load() {
	// Load all commands in to botCommands map //
	loadMaps()
	//loadHelp()
}

/*
func Save() {
	//saveBotMaps()
}*/

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

// Loads all maps with in commands.go //
func loadMaps() {

	/*
		encdec := EncDec{}
		encdec.OpenFile()
		encdec.NewDecGob()
		err := encdec.DecGob(&BanTime)
		if err != nil {
			log.Println(err)
		}
		encdec.CloseFile()*/

	// Bot Commands map //
	BotCommands["!message"] = messageCommands
	BotCommands["!ms"] = messageCommands
	BotCommands["!player"] = playerCommands
	BotCommands["!pl"] = playerCommands
	BotCommands["!channel"] = channelCommands
	BotCommands["!ch"] = channelCommands
	BotCommands["!guild"] = guildCommands
	BotCommands["!gl"] = guildCommands
	BotCommands["!utility"] = utilityCommands
	BotCommands["!util"] = utilityCommands
	//BotCommands["!help"] = Help

	// Message Commands map //
	messageCommands["-delete"] = deleteMessage
	messageCommands["-del"] = deleteMessage
	messageCommands["-clear"] = clearMessages
	messageCommands["-cl"] = clearMessages

	// Player Commands map //
	playerCommands["-kick"] = kickMember
	playerCommands["-k"] = kickMember
	playerCommands["-ban"] = banMember
	playerCommands["-b"] = banMember
	playerCommands["-bantimer"] = newBanTimer
	playerCommands["-bt"] = newBanTimer

	// Channel Commands map //
	channelCommands["-create"] = createChannel
	channelCommands["-c"] = createChannel
	channelCommands["-delete"] = deleteChannel
	channelCommands["-del"] = deleteChannel

	// Utility Commands map //
	utilityCommands["-dice"] = diceRole
	utilityCommands["-d"] = diceRole
	/*utilityCommands["-ign"] = ign
	utilityCommands["-trinity"] = trinity
	utilityCommands["-t"] = trinity*/

}
