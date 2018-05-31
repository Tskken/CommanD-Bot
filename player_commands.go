package CommanD_Bot

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"strings"
	"log"
	"time"
	"strconv"
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

	// TODO - Comment
	p.subCommands["-mute"] = muteUser
	p.subCommands["-m"] = muteUser

	// TODO - Comment
	p.subCommands["-unmute"] = unMuteUser
	p.subCommands["-um"] = unMuteUser

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
		if _, err := session.ChannelMessageSend(message.ChannelID, "You do not have the permission to kick someone."); err != nil {
			return err
		}
		return deleteMessage(session, message.ChannelID, message.ID)
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

					if err := session.GuildMemberDelete(guild.ID, member.User.ID); err != nil {
						return err
					}

					if _, err := session.ChannelMessageSend(message.ChannelID, args[2] + " was kicked"); err != nil {
						return err
					}

					return deleteMessage(session, message.ChannelID, message.ID)
				} else if len(args) >= 4 {
					// Kick user with reason //
					if err := session.GuildMemberDeleteWithReason(guild.ID, member.User.ID, strings.Join(args[3:], " ")); err != nil {
						return err
					}

					return deleteMessage(session, message.ChannelID, message.ID)
				} else {
					// Could not find user with in guild //
					// - returns an error (nil if non)
					if _, err := session.ChannelMessageSend(message.ChannelID, "You did not enter a user to kick.  Type !help -kick for more info."); err != nil {
						return err
					}
					return deleteMessage(session, message.ChannelID, message.ID)
				}
			}
		}

		if _, err := session.ChannelMessageSend(message.ChannelID, "given user mention was not found in server"); err != nil {
			return err
		}
		// User was not found in server //
		// - return an error for no user found
		return deleteMessage(session, message.ChannelID, message.ID)
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
		if _, err := session.ChannelMessageSend(message.ChannelID, "You do not have the permission to kick someone."); err != nil {
			return err
		}

		if _, err := session.ChannelMessageSend(message.ChannelID, args[2] + " was banned"); err != nil {
			return err
		}

		return deleteMessage(session, message.ChannelID, message.ID)
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
			if err := deleteMessage(session, message.ChannelID, message.ID); err != nil {
				return err
			}
			return errors.New("guild did not exist in serverList")
		}

		if len(args) == 3 {
			// Find the user to ban with in the guild //
			for _, member := range guild.Members {

				// Check if mention of user is equal to argument //
				if args[2] == member.User.Mention() {
					if err := deleteMessage(session, message.ChannelID, message.ID); err != nil {
						return err
					}

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
					if err := deleteMessage(session, message.ChannelID, message.ID); err != nil {
						return err
					}

					// Ban user with given reason for set server ban time //
					// - returns an error (nil if non)
					return session.GuildBanCreateWithReason(guild.ID, member.User.ID, strings.Join(args[3:], " "), int(server.BanTime))
				}
			}

			if _, err := session.ChannelMessageSend(message.ChannelID, "given user mention was not found in server"); err != nil {
				return err
			}

			// given user was found in guild //
			// - return an error
			return deleteMessage(session, message.ChannelID, message.ID)
		} else {
			// Given arguments were not correct //
			// - return an error (nil if non)
			if _, err := session.ChannelMessageSend(message.ChannelID, "Could not understand given arguments."); err != nil {
				return err
			}

			return deleteMessage(session, message.ChannelID, message.ID)
		}
	}
}

// TODO - Comment
func isMuted(session *discordgo.Session, message *discordgo.Message) (bool, error) {
	guild, err := getGuild(session, message)
	if err != nil {
		return false, err
	}
	server := serverList[guild.ID]

	member, err := getMember(session, message)
	if err != nil {
		return false, err
	}

	if muteTime, muted := server.isMuted(member.User.ID); muted {
		log.Println("is muted till " + time.Until(muteTime).Truncate(time.Second).String())
		if err := deleteMessage(session, message.ChannelID, message.ID); err != nil {
			return false, err
		}

		_, err := session.ChannelMessageSend(message.ChannelID, member.User.Mention()+" you are muted for "+time.Until(muteTime).Truncate(time.Second).String())
		return true, err
	}

	return false, nil
}

// TODO - Comment
func muteUser(session *discordgo.Session, message *discordgo.Message, args []string) error {
	if len(args) < 4 {
		return errors.New("to few args in muteUser")
	}

	if ok, err := isAdmin(session, message); err != nil {
		return err
	} else if !ok {
		member, err := getMember(session, message)
		if err != nil {
			return err
		}
		_, err = session.ChannelMessageSend(message.ChannelID, member.User.Mention()+" You do not have permission to use that command.")
		return err
	}

	guild, err := getGuild(session, message)
	if err != nil {
		return err
	}

	for _, member := range guild.Members {
		if member.User.Mention() == args[2] {

			var key time.Duration

			d := args[3]
			k := d[len(d)-1]
			switch k {
			case 's':
				key = time.Second
			case 'm':
				key = time.Minute
			case 'h':
				key = time.Hour
			default:
				return errors.New("unknown time key " + string(d))
			}

			t, err := strconv.Atoi(args[3][:len(d)-1])
			if err != nil {
				return err
			}

			duration := time.Duration(t) * key

			server, ok := serverList[guild.ID]
			if !ok {
				return errors.New("server did not exist in server list")
			}

			if err := server.mute(member.User.ID, duration); err != nil {
				return err
			}

			if _, err := session.ChannelMessageSend(message.ChannelID, args[2]+" has been muted for "+duration.Truncate(time.Second).String()); err != nil {
				return err
			}

			return deleteMessage(session, message.ChannelID, message.ID)
		}
	}

	return errors.New("user was not found")
}

// TODO - Comment
func unMuteUser(session *discordgo.Session, message *discordgo.Message, args []string) error {
	if len(args) < 3 {
		return errors.New("to few args in muteUser")
	}

	if ok, err := isAdmin(session, message); err != nil {
		return err
	} else if !ok {
		member, err := getMember(session, message)
		if err != nil {
			return err
		}
		_, err = session.ChannelMessageSend(message.ChannelID, member.User.Mention()+" You do not have permission to use that command.")
		return err
	}

	guild, err := getGuild(session, message)
	if err != nil {
		return err
	}

	server, ok := serverList[guild.ID]
	if !ok {
		return errors.New("guild did not exist in serverList")
	}

	for _, member := range guild.Members {
		if member.User.Mention() == args[2] {
			if err := server.unMute(member.User.ID); err != nil {
				return err
			}

			if _, err := session.ChannelMessageSend(message.ChannelID, args[2]+" has been unmuted"); err != nil {
				return err
			}

			return deleteMessage(session, message.ChannelID, message.ID)

		}
	}
	return errors.New("could not find user")
}
