package CommanD_Bot

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"strings"
)

// Set player command structure //
func loadPlayerCommand() *commands {
	// Create player command structure //
	p := commands{}

	// Load player command info structure //
	p.commandInfo = loadPlayerCommandInfo()

	// Create sub command map //
	p.subCommands = make(map[string]func(*discordgo.Session, *discordgo.Message, []string) error)

	// Set kick sub command //
	p.subCommands["-kick"] = kickMember
	p.subCommands["-k"] = kickMember

	// Set ban sub command //
	p.subCommands["-ban"] = banMember
	p.subCommands["-b"] = banMember

	// Return a reference to command structure //
	return &p
}

// Create CommandInfo struct data //
func loadPlayerCommandInfo() *commandInfo {
	// Create player command info structure //
	p := commandInfo{}

	// Set default info for player commands //
	p.detail = "**!player** or **!pl** : All commands that pertain to manipulating players with in a server."

	// Create sub command info map //
	p.commands = make(map[string]string)

	// Set kick sub command info //
	p.commands["-kick"] = "**-kick** or **-k**.\n*" +
		"*Info**: Kicks a given member from the server.\n" +
		"**Arguments:**\n" +
		"    **<Member name>**: The member given will be kicked from the server. (only admins can use this command)."

	// Set ban sub command info //
	p.commands["-ban"] = "**-ban** or **-b**\n" +
		"**Info**: Bans a given user from the server for a preset amount of time. (default is 30 days).\n" +
		"**Arguments:**\n" +
		"    **<Member name>**: The member given will be baned from the server. (only admins can use this)."

	// Return reference to command info structure //
	return &p
}

// User kick function //
// - returns an error (nil if non)
// TODO - Fix kick with reason functionality
func kickMember(session *discordgo.Session, message *discordgo.Message, args []string) error {
	// Check if user is admin //
	// - return an error if err is not nil
	if admin, err := isAdmin(session, message); err != nil {
		return err
	} else if !admin {
		// User was not an admin //
		// - return an error (nil if non)
		_, err := session.ChannelMessageSend(message.ChannelID, "You do not have the permission to kick someone.")
		return err
	}

	// Gets the guild the messages was created in //
	// - returns an error if err is not nil
	if guild, err := getGuild(session, message); err != nil {
		return err
	} else {
		// Find the user with in the guild //
		for _, member := range guild.Members {
			// Check if mention of user is the same as passed user mention //
			if args[2] == member.User.Mention() {
				if len(args) == 3 {
					// Kick user with out reason //
					return session.GuildMemberDelete(guild.ID, member.User.ID)
				} else if len(args) >= 4 {
					// Kick user with reason //
					return session.GuildMemberDeleteWithReason(guild.ID, member.User.ID, strings.Join(args[3:], " "))
				} else {
					// Could not find user with in guild //
					// - returns an error (nil if non)
					_, err := session.ChannelMessageSend(message.ChannelID, "You did not enter a user to kick.  Type !help -kick for more info.")
					return err
				}
			}
		}
		// User was not found in server //
		// - return an error for no user found
		return errors.New("given user mention was not found in server")
	}
}

// User ban function //
// - returns an error (nil if non)
// TODO - Fix Ban with reason functionality
func banMember(session *discordgo.Session, message *discordgo.Message, args []string) error {
	// Check if user is an admin //
	// - returns an error if err is not nil
	if admin, err := isAdmin(session, message); err != nil {
		return err
	} else if !admin {
		// User was not an admin //
		// - return an error (nil if non)
		_, err := session.ChannelMessageSend(message.ChannelID, "You do not have the permission to kick someone.")
		return err
	}

	// Gets the guild the messages was created in //
	// - returns an error if err is not nil
	if guild, err := getGuild(session, message); err != nil {
		return err
	} else {
		// Get the ban time for a server //
		// - returns an error if guild is not in server list
		server, ok := serverList[guild.ID]
		if !ok {
			return errors.New("guild did not exist in serverList")
		}

		if len(args) == 3 {
			// Find the user to ban with in the guild //
			for _, member := range guild.Members {

				// Check if mention of user is equal to argument //
				if args[2] == member.User.Mention() {
					// Ban the user for set server ban time //
					// - returns an error (nil if non)
					return session.GuildBanCreate(guild.ID, member.User.ID, int(server.BanTime))
				}
			}

			// given user was found in guild //
			// - return an error
			return errors.New("no user was found")
		} else if len(args) >= 4 {
			// Find the user to ban with in the guild //
			for _, member := range guild.Members {

				// Check if mention of user is qual to argument //
				if member.Nick == args[2] || member.User.Username == args[2] {
					// Ban user with given reason for set server ban time //
					// - returns an error (nil if non)
					return session.GuildBanCreateWithReason(guild.ID, member.User.ID, strings.Join(args[3:], " "), int(server.BanTime))
				}
			}

			// given user was found in guild //
			// - return an error
			return errors.New("no user was found")
		} else {
			// Given arguments were not correct //
			// - return an error (nil if non)
			_, err := session.ChannelMessageSend(message.ChannelID, "Could not understand given arguments.")
			return err
		}
	}
}
