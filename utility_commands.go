package CommanD_Bot

import (
	"math/rand"
	"strconv"
	"time"
)

type UtilityCommands struct {
	commands map[string]func(*Root)error
}

func (m *UtilityCommands) RunCommand(root *Root) error {
	return m.commands[root.CommandType()](root)
}

// Set utility command structure //
func LoadUtilityCommand() *UtilityCommands {
	// Create utility command structure //
	u := UtilityCommands{}

	// Create utility sub command map //
	u.commands = make(map[string]func(*Root) error)

	// Set dice role function in map //
	u.commands["-dice"] = DiceRole
	u.commands["-d"] = DiceRole

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
func DiceRole(root *Root) error {
	// Convert the third argument to an int //
	// - returns an error if err is not nil
	if val, err := strconv.Atoi(root.CommandArgs()[0]); err != nil {
		if err := root.DeleteMessage(root.ID); err != nil {
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
		if err := root.MessageSend(root.Author.Mention()+" got "+strconv.Itoa(val)); err != nil {
			return err
		}

		return root.DeleteMessage(root.ID)
	}
}
