package commands

import (
	"github.com/bwmarrin/discordgo"
)

/*
 - TODO : Refactor command structure with interfaces
*/

// Bot Command Dictionary //
var BotCommands = make(map[string]map[string]func(*discordgo.Session, *discordgo.Message) error)

// Sub Commands //
var messageCommands = make(map[string]func(*discordgo.Session, *discordgo.Message) error)
var helpCommands = make(map[string]func(*discordgo.Session, *discordgo.Message) error)
var playerCommands = make(map[string]func(*discordgo.Session, *discordgo.Message) error)
var channelCommands = make(map[string]func(*discordgo.Session, *discordgo.Message) error)
var guildCommands = make(map[string]func(*discordgo.Session, *discordgo.Message) error)
var utilityCommands = make(map[string]func(*discordgo.Session, *discordgo.Message) error)
var filterCommands = make(map[string]func(*discordgo.Session, *discordgo.Message) error)

// KeyWord filter map //
var KeyWordMap map[string]bool

// Load command maps with bot commands //
func Load() {
	// Load all commands in to botCommands map //
	loadMaps()
}

// Loads all maps with in commands.go //
func loadMaps() {

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
	BotCommands["!filter"] = filterCommands
	BotCommands["!fl"] = filterCommands
	BotCommands["!help"] = helpCommands

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

	// Filter Commands map //
	filterCommands["-bad"] = setBad
	filterCommands["-b"] = setBad
	filterCommands["-good"] = setGood
	filterCommands["-g"] = setGood

	// Help Commands //
	setMCInfo()
	setPCInfo()
	helpCommands["-messages"] = messageHelp
	helpCommands["-ms"] = messageHelp
	helpCommands["-player"] = playerHelp
	helpCommands["-pl"] = playerHelp
}
