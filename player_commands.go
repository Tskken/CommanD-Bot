package CommanD_Bot

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"time"
)

// Set player command structure //
func loadPlayerCommand() *PlayerCommands {
	// Create player command structure //
	p := PlayerCommands{}

	// Create sub command map //
	p.commands = make(map[CommandKey]func(RootCommand) error)

	// Set kick sub command //
	p.commands["-kick"] = kickMember
	p.commands["-k"] = kickMember

	// Set ban sub command //
	p.commands["-ban"] = banMember
	p.commands["-b"] = banMember

	// TODO - Comment
	p.commands["-mute"] = muteUser
	p.commands["-m"] = muteUser

	// TODO - Comment
	p.commands["-unmute"] = unMuteUser
	p.commands["-um"] = unMuteUser

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
func kickMember(command RootCommand) error {
	// Check if user is admin //
	// - return an error if err is not nil
	if admin, err := command.isAdmin(); err != nil {
		return err
	} else if !admin {
		// User was not an admin //
		// - return an error (nil if non)
		if _, err := command.ChannelMessageSend(command.ChannelID, "You do not have the permission to kick someone."); err != nil {
			return err
		}
		return command.deleteMessage(command.ID)
	}

	// Gets the guild the messages was created in //
	// - returns an error if err is not nil
	if guild, err := command.getGuild(); err != nil {
		return err
	} else {
		// Find the user with in the guild //
		for _, member := range guild.Members {
			// Check if mention of user is the same as passed user mention //
			if command.args[0] == member.User.Mention() {
				if len(command.args) == 1 {
					// Kick user with out reason //

					if err := command.GuildMemberDelete(guild.ID, member.User.ID); err != nil {
						return err
					}

					if _, err := command.ChannelMessageSend(command.ChannelID, command.args[0] + " was kicked"); err != nil {
						return err
					}

					return command.deleteMessage(command.ID)
				} else if len(command.args) >= 2 {
					// Kick user with reason //
					if err := command.GuildMemberDeleteWithReason(guild.ID, member.User.ID, strings.Join(command.args[2:], " ")); err != nil {
						return err
					}

					return command.deleteMessage(command.ID)
				} else {
					// Could not find user with in guild //
					// - returns an error (nil if non)
					if _, err := command.ChannelMessageSend(command.ChannelID, "You did not enter a user to kick.  Type !help -kick for more info."); err != nil {
						return err
					}
					return command.deleteMessage(command.ID)
				}
			}
		}

		if _, err := command.ChannelMessageSend(command.ChannelID, "given user mention was not found in server"); err != nil {
			return err
		}
		// User was not found in server //
		// - return an error for no user found
		return command.deleteMessage(command.ID)
	}
}

// User ban function //
// - returns an error (nil if non)
// TODO - Fix Ban with reason functionality
func banMember(command RootCommand) error {
	// Check if user is an admin //
	// - returns an error if err is not nil
	if admin, err := command.isAdmin(); err != nil {
		return err
	} else if !admin {
		// User was not an admin //
		// - return an error (nil if non)
		if _, err := command.ChannelMessageSend(command.ChannelID, "You do not have the permission to kick someone."); err != nil {
			return err
		}
		return command.deleteMessage(command.ID)
	}

	// Gets the guild the messages was created in //
	// - returns an error if err is not nil
	if guild, err := command.getGuild(); err != nil {
		return err
	} else {
		// Get the ban time for a server //
		// - returns an error if guild is not in server list
		server, ok := serverList[guild.ID]
		if !ok {
			if err := command.deleteMessage(command.ID); err != nil {
				return err
			}
			return errors.New("guild did not exist in serverList")
		}

		if len(command.args) == 1 {
			// Find the user to ban with in the guild //
			for _, member := range guild.Members {

				// Check if mention of user is equal to argument //
				if command.args[0] == member.User.Mention() {
					if err := command.deleteMessage(command.ID); err != nil {
						return err
					}

					// Ban the user for set server ban time //
					// - returns an error (nil if non)
					return command.GuildBanCreate(guild.ID, member.User.ID, int(server.BanTime))
				}
			}

			// given user was found in guild //
			// - return an error
			return errors.New("no user was found")
		} else if len(command.args) >= 2 {
			// Find the user to ban with in the guild //
			for _, member := range guild.Members {

				// Check if mention of user is equal to argument //
				if member.Nick == command.args[0] || member.User.Username == command.args[0] {
					if err := command.deleteMessage(command.ID); err != nil {
						return err
					}

					// Ban user with given reason for set server ban time //
					// - returns an error (nil if non)
					return command.GuildBanCreateWithReason(guild.ID, member.User.ID, strings.Join(command.args[1:], " "), int(server.BanTime))
				}
			}

			if _, err := command.ChannelMessageSend(command.ChannelID, "given user mention was not found in server"); err != nil {
				return err
			}

			// given user was found in guild //
			// - return an error
			return command.deleteMessage(command.ID)
		} else {
			// Given arguments were not correct //
			// - return an error (nil if non)
			if _, err := command.ChannelMessageSend(command.ChannelID, "Could not understand given arguments."); err != nil {
				return err
			}

			return command.deleteMessage(command.ID)
		}
	}
}

// TODO - Comment
func (r *Root) isMuted() (bool, error) {
	guild, err := r.getGuild()
	if err != nil {
		return false, err
	}
	server := serverList[guild.ID]

	member, err := r.getMember()
	if err != nil {
		return false, err
	}

	if muteTime, muted := server.isMuted(member.User.ID); muted {
		log.Println("is muted till " + time.Until(muteTime).Truncate(time.Second).String())
		if err := r.deleteMessage(r.ID); err != nil {
			return false, err
		}

		_, err := r.ChannelMessageSend(r.ChannelID, member.User.Mention()+" you are muted for "+time.Until(muteTime).Truncate(time.Second).String())
		return true, err
	}

	return false, nil
}

// TODO - Comment
func muteUser(command RootCommand) error {
	if len(command.args) < 2 {
		return errors.New("to few args in muteUser")
	}

	if ok, err := command.isAdmin(); err != nil {
		return err
	} else if !ok {
		member, err := command.getMember()
		if err != nil {
			return err
		}
		_, err = command.ChannelMessageSend(command.ChannelID, member.User.Mention()+" You do not have permission to use that command.")
		return err
	}

	guild, err := command.getGuild()
	if err != nil {
		return err
	}

	for _, member := range guild.Members {
		if member.User.Mention() == command.args[0] {

			var key time.Duration

			d := command.args[1]
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

			t, err := strconv.Atoi(command.args[1][:len(d)-1])
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

			if _, err := command.ChannelMessageSend(command.ChannelID, command.args[0]+" has been muted for "+duration.Truncate(time.Second).String()); err != nil {
				return err
			}

			return command.deleteMessage(command.ID)
		}
	}

	return errors.New("user was not found")
}

// TODO - Comment
func unMuteUser(command RootCommand) error {
	if len(command.args) < 1 {
		return errors.New("to few args in muteUser")
	}

	if ok, err := command.isAdmin(); err != nil {
		return err
	} else if !ok {
		member, err := command.getMember()
		if err != nil {
			return err
		}
		_, err = command.ChannelMessageSend(command.ChannelID, member.User.Mention()+" You do not have permission to use that command.")
		return err
	}

	guild, err := command.getGuild()
	if err != nil {
		return err
	}

	server, ok := serverList[guild.ID]
	if !ok {
		return errors.New("guild did not exist in serverList")
	}

	for _, member := range guild.Members {
		if member.User.Mention() == command.args[0] {
			if err := server.unMute(member.User.ID); err != nil {
				return err
			}

			if _, err := command.ChannelMessageSend(command.ChannelID, command.args[0]+" has been unmuted"); err != nil {
				return err
			}

			return command.deleteMessage(command.ID)

		}
	}
	return errors.New("could not find user")
}
