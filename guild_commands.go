package CommanD_Bot

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"strconv"
)

// Set guild command structure //
func loadGuildCommand() *commands {
	// Create guild command structure //
	g := commands{}

	// Load guild command help info //
	g.commandInfo = loadGuildCommandInfo()

	// Create sub command map //
	g.subCommands = make(map[string]func(*discordgo.Session, *discordgo.Message, []string) error)

	// Add word filter command to map //
	g.subCommands["-wordfilter"] = wordFilter
	g.subCommands["-wf"] = wordFilter

	// Add ban time command to map //
	g.subCommands["-bantime"] = banTime
	g.subCommands["-bt"] = banTime

	// Add word filter threshold command to map //
	g.subCommands["-wordfilterthreshold"] = wordFilterThreshold
	g.subCommands["-wft"] = wordFilterThreshold

	// Add spam filter threshold command to map //
	g.subCommands["-spamfilterthreshold"] = spamFilterThreshold
	g.subCommands["-sft"] = spamFilterThreshold

	// Return a reference to the new guild command structure //
	return &g
}

// Set help info for guild commands //
func loadGuildCommandInfo() *commandInfo {
	// Create help info structure //
	g := commandInfo{}

	// Set guild command default info //
	g.detail = "**!guild** or **!g** : All commands that pertain to manipulating guild info."

	// Create sub command help info map //
	g.commands = make(map[string]string)

	// Set word filter command info //
	g.commands["-wordfilter"] = "**-wordfilter** or **-wf**.\n" +
		"**Info**: Adds or removes messages from the word filter.\n" +
		"**Arguments:**\n" +
		"    **add <word>:** Adds the given word to the filter\n" +
		"    **remove <word>:** Removes the givne word from the filter"

	// Set ban time command info //
	g.commands["-bantime"] = "**-bantime** or **-bt**.\n" +
		"**Info**: changes the ban length when baning some one from the server.\n" +
		"**Arguments:**\n" +
		"    **<number grater then 0>:** Changes the ban time to the time entered."

	// Set word filter threshold command info //
	g.commands["-wordfilterthresh"] = "**-wordfilterthresh** or **-wft**.\n" +
		"**Info**: Change the threshold for deleting messages with bad words.\n" +
		"**Arguments:**\n" +
		"    **<number between 0.0 and 1.0>:** Changes the threshold to the given number."

	// Set spam filter threshold command info //
	g.commands["-spamfilterthresh"] = "**-spamfilterthresh** or **-sft**.\n" +
		"**Info**: Changes the threshold for deleting messages through the spam filter.\n" +
		"**Arguments:**\n" +
		"    *<number between 0.0 and 1.0>:** Changes the threshold to the given number."

	// Return a reference to the new guild info structure //
	return &g
}

// Word filter modification function //
// - returns an error (nil if non)
func wordFilter(session *discordgo.Session, message *discordgo.Message, args []string) error {
	// Check if user is an admin //
	// - returns an error if err is not nil
	// - returns if user is not an admin
	if admin, err := isAdmin(session, message); err != nil {
		return err
	} else if !admin {
		_, err := session.ChannelMessageSend(message.ChannelID, "You do not have permission to change the ban length.")
		return err
	}

	// Check to make sure all arguments were given //
	// - returns an error if they were not
	if len(args) < 4 {
		return errors.New("the number of messages with in the command was to short")
	}

	// Get guild to edit the world filter for //
	// - returns an error if err is not nil
	if g, err := getGuild(session, message); err != nil {
		return err
	} else {
		// Get server info for guild //
		// - returns an error if guild does not exist in list
		if server, ok := serverList[g.ID]; !ok {
			return errors.New("no server with in serverList")
		} else {
			// Change word for given argument //
			// - add: add a word to the list
			// - remove: remove the word from the list
			// - returns an error if it was not add or remove
			switch args[2] {
			case "add":
				// Add word to list //
				// - returns an error (nil if non)
				return server.editWordFilter(args[3], true)
			case "remove":
				// Remove word from list //
				// - returns an error (nil if non)
				return server.editWordFilter(args[3], false)
			default:
				// Return an error as argument was not add or remove
				return errors.New("argument given could not be understood")
			}
		}
	}
}

// Ban time change function //
// - returns an error (nil if non)
func banTime(session *discordgo.Session, message *discordgo.Message, args []string) error {
	// Check if user is an admin //
	// - return an error if err is not nil
	// - return if user is not an admin
	if admin, err := isAdmin(session, message); err != nil {
		return err
	} else if !admin {
		_, err := session.ChannelMessageSend(message.ChannelID, "You do not have permission to change the ban length.")
		return err
	}

	// Gets the guild the messages was created in //
	// - returns an error if err is not nil
	if guild, err := getGuild(session, message); err != nil {
		return err
	} else {
		// Get server info for guild //
		// - returns an error if guild does not exist in list
		server, ok := serverList[guild.ID]
		if !ok {
			return errors.New("guild did not exist in server list")
		}

		// Send the current ban time to the channel //
		// - returns an error if err is not nil
		if _, err := session.ChannelMessageSend(message.ChannelID, "Ban time was "+strconv.Itoa(int(server.BanTime))); err != nil {
			return err
		}

		// Check length of args //
		// - return if no new time was given
		// - change time if a new time was given
		if len(args) < 3 {
			return nil
		} else {
			// Convert the given time to an integer //
			// - return an error if err is not nil
			if time, err := strconv.Atoi(args[2]); err != nil {
				return err
			} else {
				// Check if given time is less then or equal to zero //
				// - return an error if true
				if time <= 0 {
					return errors.New("time was under 0 days")
				}
				// Set new ban time //
				server.BanTime = uint(time)

				// Print the new ban time to the server //
				// - return an error if err is not nil
				_, err = session.ChannelMessageSend(message.ChannelID, "Ban time has been set to "+strconv.Itoa(int(server.BanTime)))
				return err
			}
		}
	}
}

// World filter threshold function //
// - return an error (nil if non)
func wordFilterThreshold(session *discordgo.Session, message *discordgo.Message, args []string) error {
	// Check if user is an admin //
	// - returns an error if err is not nil
	// - returns if user is not an admin
	if admin, err := isAdmin(session, message); err != nil {
		return err
	} else if !admin {
		_, err := session.ChannelMessageSend(message.ChannelID, "You do not have permition to change word filter settings.")
		return err
	}

	// Get guild the message was sent in //
	// - returns an error if err is not nil
	if guild, err := getGuild(session, message); err != nil {
		return err
	} else {
		// Get server info for guild //
		// - returns an error if guild is not in server list
		server, ok := serverList[guild.ID]
		if !ok {
			return errors.New("guild was not in the serverList")
		}

		// Send message for current word filter threshold //
		// - returns an error if err is not nil
		if _, err := session.ChannelMessageSend(message.ChannelID, "Current word filter threshold is "+strconv.FormatFloat(server.wordFilterThresh, 'f', 2, 64)); err != nil {
			return err
		}

		// Check if new threshold was given //
		// - returns if false
		if len(args) < 3 {
			return nil
		} else {
			// Convert argument to float for new threshold //
			// - return an error if err is not nil
			if thresh, err := strconv.ParseFloat(args[2], 64); err != nil {
				return err
			} else {
				// Check if threshold given is with in range //
				// - returns an error if its not
				if thresh < 0.0 || thresh > 1.0 {
					return errors.New("given threshold was ether less then 0.0 or greater then 1.0")
				}

				// Set new threshold to server info //
				server.wordFilterThresh = thresh

				// Send message displaying changed word filter threshold //
				// - returns an error (nil if non)
				_, err := session.ChannelMessageSend(message.ChannelID, "New word filter threshold is "+strconv.FormatFloat(server.wordFilterThresh, 'f', 2, 64))
				return err
			}
		}
	}
}

// Spam filter threshold function //
// - returns an error (nil if non)
func spamFilterThreshold(session *discordgo.Session, message *discordgo.Message, args []string) error {
	// Check if user is an admin //
	// - returns an error if err is not nil
	// - returns if user is not an admin
	if admin, err := isAdmin(session, message); err != nil {
		return err
	} else if !admin {
		_, err := session.ChannelMessageSend(message.ChannelID, "You do not have permition to change word filter settings.")
		return err
	}

	// Get guild the message was sent in //
	// - returns an error if err is not nil
	if guild, err := getGuild(session, message); err != nil {
		return err
	} else {
		// Get server info for guild //
		// - returns an error if guild is not in server list
		server, ok := serverList[guild.ID]
		if !ok {
			return errors.New("guild was not in the serverList")
		}

		// Send message for current spam filter threshold //
		// - returns an error if err is not nil
		if _, err := session.ChannelMessageSend(message.ChannelID, "Current spam filter threshold is "+strconv.FormatFloat(server.spamFilterThresh, 'f', 2, 64)); err != nil {
			return err
		}

		// Check if new threshold was given //
		// - returns if false
		if len(args) < 3 {
			return nil
		} else {
			// Convert argument to float for new threshold //
			// - return an error if err is not nil
			if thresh, err := strconv.ParseFloat(args[2], 64); err != nil {
				return err
			} else {
				// Check if threshold given is with in range //
				// - returns an error if its not
				if thresh < 0.0 || thresh > 1.0 {
					return errors.New("given threshold was ether less then 0.0 or greater then 1.0")
				}

				// Set new threshold to server info //
				server.spamFilterThresh = thresh

				// Send message displaying changed spam filter threshold //
				// - returns an error (nil if non)
				_, err := session.ChannelMessageSend(message.ChannelID, "New word filter threshold is "+strconv.FormatFloat(server.spamFilterThresh, 'f', 2, 64))
				return err
			}
		}
	}
}
