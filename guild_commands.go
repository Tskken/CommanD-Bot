package CommanD_Bot

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"strings"
	"strconv"
)

func loadGuildCommand() *commands {
	g := commands{}

	g.commandInfo = loadGuildCommandInfo()
	g.subCommands = make(map[string]func(*discordgo.Session, *discordgo.Message)error)
	g.subCommands["-wordfilter"] = wordFilter
	g.subCommands["-wf"] = wordFilter
	g.subCommands["-bantime"] = banTime
	g.subCommands["-bt"] = banTime
	g.subCommands["-wordfilterthreshold"] = wordFilterThreashold
	g.subCommands["-wft"] = wordFilterThreashold
	g.subCommands["-spamfilterthreshold"] = spamFilterThreshold
	g.subCommands["-sft"] = spamFilterThreshold

	return &g
}

func loadGuildCommandInfo() *commandInfo {
	g := commandInfo{}
	g.detail = "**!guild** or **!g** : All commands that pertain to manipulating guild info."
	g.commands = make(map[string]string)
	g.commands["-wordfilter"] = "**-wordfilter** or **-wf**.\n" +
		"**Info**: Adds or removes messages from the word filter.\n" +
		"**Arguments:**\n" +
			"    **add <word>:** Adds the given word to the filter\n" +
			"    **remove <word>:** Removes the givne word from the filter"
	g.commands["-bantime"] = "**-bantime** or **-bt**.\n" +
		"**Info**: changes the ban length when baning some one from the server.\n" +
		"**Arguments:**\n" +
		"    **<number grater then 0>:** Changes the ban time to the time entered."
	g.commands["-wordfilterthresh"] = "**-wordfilterthresh** or **-wft**.\n" +
		"**Info**: Change the threshold for deleting messages with bad words.\n" +
		"**Arguments:**\n" +
		"    **<number between 0.0 and 1.0>:** Changes the threshold to the given number."
	g.commands["-spamfilterthresh"] = "**-spamfilterthresh** or **-sft**.\n" +
		"**Info**: Changes the threshold for deleting messages through the spam filter.\n" +
		"**Arguments:**\n" +
		"    *<number between 0.0 and 1.0>:** Changes the threshold to the given number."
	return &g
}

func wordFilter(session *discordgo.Session, message *discordgo.Message) error {
	admin, err := isAdmin(session, message)
	if err != nil {
		return err
	}
	if !admin {
		_, err := session.ChannelMessageSend(message.ChannelID, "You do not have permission to change the ban length.")
		return err
	}

	args := strings.Fields(message.Content)

	if len(args) < 4 {
		return errors.New("the number of messages with in the command was to short")
	}

	g, err := getGuild(session, message)
	if err != nil {
		return err
	}
	server, ok := serverList[g.ID]
	if !ok {
		return errors.New("no server with in serverList")
	}

	switch args[2] {
	case "add":
		return server.editWordFilter(args[3], true)
	case "remove":
		return server.editWordFilter(args[3], false)
	default:
		return errors.New("argument given could not be understood")
	}
}

func banTime(session *discordgo.Session, message *discordgo.Message) error {
	admin, err := isAdmin(session, message)
	if err != nil {
		return err
	}
	if !admin {
		_, err := session.ChannelMessageSend(message.ChannelID, "You do not have permission to change the ban length.")
		return err
	}

	// Gets the guild the messages was created in //
	// Returns an error if err is not nil
	if guild, err := getGuild(session, message); err != nil {
		return err
	} else {
		server := serverList[guild.ID]
		// Check if the server already as a time set //
		// If it does print it to the channel
		if _, err := session.ChannelMessageSend(message.ChannelID, "Ban time was "+strconv.Itoa(int(server.BanTime))); err != nil {
			return err
		}

		// Parce the contents of the messages on a space //
		args := strings.Fields(message.Content)

		// If the there was no time given print an error to the server //
		if len(args) < 3 {
			_, err := session.ChannelMessageSend(message.ChannelID, "You need to give an amount of time to change the ban time to.")
			return err
			// The time was passed as an argument //
		} else {
			// Convert the amount of time to an int //
			// Returns an error if err is not nil
			if time, err := strconv.Atoi(args[2]); err != nil {
				return err
			} else {
				// Set the ban time for the new time given //
				if time <= 0 {
					return errors.New("time was under 0 days")
				}

				server := serverList[guild.ID]
				server.setBanTimer(uint(time))

				// Print the new ban time to the server //
				_, err = session.ChannelMessageSend(message.ChannelID, "Ban time has been set to "+strconv.Itoa(int(server.BanTime)))
				return err
			}
		}
	}
}

func wordFilterThreashold(session *discordgo.Session, message *discordgo.Message) error {
	admin, err := isAdmin(session, message)
	if err != nil {
		return err
	}
	if !admin {
		_, err := session.ChannelMessageSend(message.ChannelID, "You do not have permition to change word filter settings.")
		return err
	}

	args := strings.Fields(message.Content)

	if len(args) < 3 {
		return errors.New("number of arguments was to few")
	}

	guild, err := getGuild(session, message)
	if err != nil {
		return err
	}
	server, ok := serverList[guild.ID]
	if !ok {
		return errors.New("guild was not in the serverList")
	}

	if thresh, err := strconv.Atoi(args[2]); err != nil {
		return err
	} else {
		server.setWordFilterThreshold(float64(thresh))
	}

	return nil
}

func spamFilterThreshold(session *discordgo.Session, message *discordgo.Message) error {
	admin, err := isAdmin(session, message)
	if err != nil {
		return err
	}
	if !admin {
		_, err := session.ChannelMessageSend(message.ChannelID, "You do not have permition to change word filter settings.")
		return err
	}

	args := strings.Fields(message.Content)

	if len(args) < 3 {
		return errors.New("number of arguments was to few")
	}

	guild, err := getGuild(session, message)
	if err != nil {
		return err
	}
	server, ok := serverList[guild.ID]
	if !ok {
		return errors.New("guild was not in the serverList")
	}

	if thresh, err := strconv.Atoi(args[2]); err != nil {
		return err
	} else {
		server.setSpamFilterThreshold(float64(thresh))
	}

	return nil
}