package CommanD_Bot

import (
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"strconv"
	"time"
)

// Set utility command structure //
func loadUtilityCommand() *commands {
	// Create utility command structure //
	u := commands{}

	// Load utility command info structure //
	u.commandInfo = loadUtilityCommandInfo()

	// Create utility sub command map //
	u.subCommands = make(map[string]func(*discordgo.Session, *discordgo.Message, []string) error)

	// Set dice role function in map //
	u.subCommands["-dice"] = diceRole
	u.subCommands["-d"] = diceRole

	// Return reference to utility command structure //
	return &u
}

// Create CommandInfo structure //
func loadUtilityCommandInfo() *commandInfo {
	// Create utility command info structure //
	u := commandInfo{}

	// Set utility default info //
	u.detail = "**!utility** or **!util** : Extra fun commands for utility with in your server."

	// Create sub command info map //
	u.commands = make(map[string]string)

	// Set dice sub command info //
	u.commands["-dice"] = "**-dice** or **-d**.\n" +
		"**Info**: Roles a dice.\n" +
		"**Arguments:**\n" +
		"    **<Number of sides>**: Roles a dice for the given number of sides (number will be between 1 and the number you gave)."

	// Return reference to utility command info structure //
	return &u
}

// Roles a dice //
// - returns an error (nil if non)
func diceRole(session *discordgo.Session, message *discordgo.Message, args []string) error {
	// Convert the third argument to an int //
	// - returns an error if err is not nil
	if val, err := strconv.Atoi(args[2]); err != nil {
		if err := deleteMessage(session, message.ChannelID, message.ID); err != nil {
			return err
		}

		return err
	} else {
		// Create a new rand instance with time now seed //
		rng := rand.New(rand.NewSource(time.Now().UnixNano()))

		// Get a random number from 0 to the given value //
		val := rng.Intn(val)

		// Shift by 1 for based 1 instead of base 0 //
		val++

		// Print random number to channel //
		// - returns an error if err is not nil
		if _, err := session.ChannelMessageSend(message.ChannelID, message.Author.Mention()+" got "+strconv.Itoa(val)); err != nil {
			return err
		}

		return deleteMessage(session, message.ChannelID, message.ID)
	}
}
