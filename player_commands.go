package CommanD_Bot

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func loadPlayerCommand() *commands {
	p := commands{}
	p.commandInfo = loadPlayerCommandInfo()
	p.subCommands = make(map[string]func(*discordgo.Session, *discordgo.Message) error)
	p.subCommands["-kick"] = kickMember
	p.subCommands["-k"] = kickMember
	p.subCommands["-ban"] = banMember
	p.subCommands["-b"] = banMember

	return &p
}

// Create CommandInfo struct data //
func loadPlayerCommandInfo() *commandInfo {

	p := commandInfo{}
	p.detail = "**!player** or **!pl** : All commands that pertain to manipulating players with in a server."
	p.commands = make(map[string]string)
	p.commands["-kick"] = "**-kick** or **-k**.\n*" +
		"*Info**: Kicks a given member from the server.\n" +
		"**Arguments:**\n" +
		"    **<Member name>**: The member given will be kicked from the server. (only admins can use this command)."
	p.commands["-ban"] = "**-ban** or **-b**\n" +
		"**Info**: Bans a given user from the server for a preset amount of time. (default is 30 days).\n" +
		"**Arguments:**\n" +
		"    **<Member name>**: The member given will be baned from the server. (only admins can use this)."
	p.commands["-bantimer"] = "**-bantimer** or **-bt**\n" +
		"**Info**: Sets the current time for banning a user.  Default is 30 days.\n" +
		"**Arguments:**\n" +
		"    **<Number to set time to>**: The number of days the ban time will be set to."
	return &p
}

// Function to kick a user from the server //
// TODO - Fix kick with reason functionality
func kickMember(session *discordgo.Session, message *discordgo.Message) error {
	// Check for admin permitions //
	// Return an error if err is not nil
	admin, err := isAdmin(session, message)
	if err != nil {
		return err
	}

	// Check caller is an admin //
	// Prints an error saying you do not have permission if you are not an admin
	if !admin {
		_, err := session.ChannelMessageSend(message.ChannelID, "You do not have the permission to kick someone.")
		return err
	}

	// Gets the guild the messages was created in //
	// Returns an error if err is not nil
	if guild, err := getGuild(session, message); err != nil {
		return err
	} else {
		// Parce the content of the message on a space //
		args := strings.Fields(message.Content)

		// Find the user with in the guild //
		for _, member := range guild.Members {
			// If members nick name or username is the same  as the name passed in as the username argument kick them from the server //
			if args[2] == member.User.Mention() {
				// Kick the user with out a reason //
				if len(args) == 3 {
					// Kick user //
					return session.GuildMemberDelete(guild.ID, member.User.ID)
					// Kick a user with a reason //
				} else if len(args) >= 4 {
					// Kick user with reason //
					return session.GuildMemberDeleteWithReason(guild.ID, member.User.ID, strings.Join(args[3:], " "))
					// User did not enter a username or nick name to kick //
					// Print an error to the channel
				} else {
					_, err := session.ChannelMessageSend(message.ChannelID, "You did not enter a user to kick.  Type !help -kick for more info.")
					return err
				}
			}
		}
	}

	// Return an error as the entered username or nick name was not found in the guild //
	return errors.New("username or nick name was not found in guild")
}

// Function to ban a user from the server for a given amount of time //
// TODO - Fix Ban with reason functionality
func banMember(session *discordgo.Session, message *discordgo.Message) error {
	// Check if user is an admin //
	// - Returns an error if err is not nil
	admin, err := isAdmin(session, message)
	if err != nil {
		return err
	}

	// Check caller is an admin //
	// Prints an error saying you do not have permission if you are not an admin
	if !admin {
		_, err := session.ChannelMessageSend(message.ChannelID, "You do not have the permission to kick someone.")
		return err
	}

	// Gets the guild the messages was created in //
	// Returns an error if err is not nil
	if guild, err := getGuild(session, message); err != nil {
		return err
	} else {
		// Parce the content of the message on a space //
		args := strings.Fields(message.Content)

		// Get the ban time for a server //
		// Prints an error if the servers ban time as not been set
		server, ok := serverList[guild.ID]
		if !ok {
			return errors.New("guild did not exist in serverList")
		}

		// Ban the user with out a reason //
		if len(args) == 3 {
			// Find the user to ban with in the guild //
			for _, member := range guild.Members {
				// If members nick name or username is the same  as the name passed in as the username argument kick them from the server //
				if args[2] == member.User.Mention() {
					// Ban the user for the servers set amount of time //
					return session.GuildBanCreate(guild.ID, member.User.ID, int(server.BanTime))
				}
			}
			// Ban the user with a reason //
		} else if len(args) >= 4 {
			// Find the user to ban with in the guild //
			for _, member := range guild.Members {
				// If members nick name or username is the same  as the name passed in as the username argument kick them from the server //
				if member.Nick == args[2] || member.User.Username == args[2] {
					// Ban the user for the servers set amount of time with a reason //
					return session.GuildBanCreateWithReason(guild.ID, member.User.ID, strings.Join(args[3:], " "), int(server.BanTime))
				}
			}
			// User did not enter a username or nick name to kick //
			// Print an error to the channel
		} else {
			_, err := session.ChannelMessageSend(message.ChannelID, "You did not enter a user to ban.  Type !help -ban for more info.")
			return err
		}

	}
	return errors.New("banMember() failed")
}
