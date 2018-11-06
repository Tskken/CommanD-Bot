package CommanD_Bot

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

type PlayerCommands struct {
	commands map[string]func(*Root)error
}

func (m *PlayerCommands) RunCommand(root *Root) error {
	return m.commands[root.CommandType()](root)
}

// Set player command structure //
func LoadPlayerCommand() *PlayerCommands {
	// Create player command structure //
	p := PlayerCommands{}

	// Create sub command map //
	p.commands = make(map[string]func(*Root) error)

	// Set kick sub command //
	p.commands["-kick"] = KickMember
	p.commands["-k"] = KickMember

	// Set ban sub command //
	p.commands["-ban"] = BanMember
	p.commands["-b"] = BanMember

	// TODO - Comment
	p.commands["-mute"] = MuteUser
	p.commands["-m"] = MuteUser

	// TODO - Comment
	p.commands["-unmute"] = UnMuteUser
	p.commands["-um"] = UnMuteUser

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
func KickMember(root *Root) error {
	// Check if user is admin //
	// - return an error if err is not nil
	if admin, err := root.IsAdmin(); err != nil {
		return err
	} else if !admin {
		// User was not an admin //
		// - return an error (nil if non)
		if err := root.MessageSend("You do not have the permission to kick someone."); err != nil {
			return err
		}
		return root.DeleteMessage(root.ID)
	}

	// Gets the guild the messages was created in //
	// - returns an error if err is not nil
	if guild, err := root.GetGuild(); err != nil {
		return err
	} else {
		// Find the user with in the guild //
		for _, member := range guild.Members {
			// Check if mention of user is the same as passed user mention //
			if root.CommandArgs()[0] == member.User.Mention() {
				if len(root.CommandArgs()) == 1 {
					// Kick user with out reason //

					if err := root.GuildMemberDelete(guild.ID, member.User.ID); err != nil {
						return err
					}

					if err := root.MessageSend(root.CommandArgs()[0] + " was kicked"); err != nil {
						return err
					}

					return root.DeleteMessage(root.ID)
				} else if len(root.CommandArgs()) >= 2 {
					// Kick user with reason //
					if err := root.GuildMemberDeleteWithReason(guild.ID, member.User.ID, strings.Join(root.CommandArgs()[2:], " ")); err != nil {
						return err
					}

					return root.DeleteMessage(root.ID)
				} else {
					// Could not find user with in guild //
					// - returns an error (nil if non)
					if err := root.MessageSend("You did not enter a user to kick.  Type !help -kick for more info."); err != nil {
						return err
					}
					return root.DeleteMessage(root.ID)
				}
			}
		}

		if err := root.MessageSend("given user mention was not found in server"); err != nil {
			return err
		}
		// User was not found in server //
		// - return an error for no user found
		return root.DeleteMessage(root.ID)
	}
}

// User ban function //
// - returns an error (nil if non)
// TODO - Fix Ban with reason functionality
func BanMember(root *Root) error {
	// Check if user is an admin //
	// - returns an error if err is not nil
	if admin, err := root.IsAdmin(); err != nil {
		return err
	} else if !admin {
		// User was not an admin //
		// - return an error (nil if non)
		if err := root.MessageSend( "You do not have the permission to kick someone."); err != nil {
			return err
		}
		return root.DeleteMessage(root.ID)
	}

	// Gets the guild the messages was created in //
	// - returns an error if err is not nil
	if guild, err := root.GetGuild(); err != nil {
		return err
	} else {
		// Get the ban time for a server //
		// - returns an error if guild is not in server list
		server, ok := serverList[guild.ID]
		if !ok {
			if err := root.DeleteMessage(root.ID); err != nil {
				return err
			}
			return errors.New("guild did not exist in serverList")
		}

		if len(root.CommandArgs()) == 1 {
			// Find the user to ban with in the guild //
			for _, member := range guild.Members {

				// Check if mention of user is equal to argument //
				if root.CommandArgs()[0] == member.User.Mention() {
					if err := root.DeleteMessage(root.ID); err != nil {
						return err
					}

					// Ban the user for set server ban time //
					// - returns an error (nil if non)
					return root.GuildBanCreate(guild.ID, member.User.ID, int(server.BanTime))
				}
			}

			// given user was found in guild //
			// - return an error
			return errors.New("no user was found")
		} else if len(root.CommandArgs()) >= 2 {
			// Find the user to ban with in the guild //
			for _, member := range guild.Members {

				// Check if mention of user is equal to argument //
				if GetMention(member, root.CommandArgs()[0]) {
					if err := root.DeleteMessage(root.ID); err != nil {
						return err
					}

					// Ban user with given reason for set server ban time //
					// - returns an error (nil if non)
					return root.GuildBanCreateWithReason(guild.ID, member.User.ID, strings.Join(root.CommandArgs()[1:], " "), int(server.BanTime))
				}
			}

			if err := root.MessageSend("given user mention was not found in server"); err != nil {
				return err
			}

			// given user was found in guild //
			// - return an error
			return root.DeleteMessage(root.ID)
		} else {
			// Given arguments were not correct //
			// - return an error (nil if non)
			if err := root.MessageSend("Could not understand given arguments."); err != nil {
				return err
			}

			return root.DeleteMessage(root.ID)
		}
	}
}

// TODO - Comment
func MuteUser(root *Root) error {
	if len(root.CommandArgs()) < 2 {
		return errors.New("to few args in muteUser")
	}

	if ok, err := root.IsAdmin(); err != nil {
		return err
	} else if !ok {
		member, err := root.GetMember()
		if err != nil {
			return err
		}
		return root.MessageSend(member.User.Mention()+" You do not have permission to use that root.")
	}

	guild, err := root.GetGuild()
	if err != nil {
		return err
	}

	for _, member := range guild.Members {
		if member.User.Mention() == root.CommandArgs()[0] {

			var key time.Duration

			d := root.CommandArgs()[1]
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

			t, err := strconv.Atoi(root.CommandArgs()[1][:len(d)-1])
			if err != nil {
				return err
			}

			duration := time.Duration(t) * key

			server, ok := serverList[guild.ID]
			if !ok {
				return errors.New("server did not exist in server list")
			}

			if err := server.Mute(member.User.ID, duration); err != nil {
				return err
			}

			if err := root.MessageSend(root.CommandArgs()[0]+" has been muted for "+duration.Truncate(time.Second).String()); err != nil {
				return err
			}

			return root.DeleteMessage(root.ID)
		}
	}

	return errors.New("user was not found")
}

// TODO - Comment
func UnMuteUser(root *Root) error {
	if len(root.CommandArgs()) < 1 {
		return errors.New("to few args in muteUser")
	}

	if ok, err := root.IsAdmin(); err != nil {
		return err
	} else if !ok {
		member, err := root.GetMember()
		if err != nil {
			return err
		}
		return root.MessageSend(member.User.Mention()+" You do not have permission to use that root.")
	}

	guild, err := root.GetGuild()
	if err != nil {
		return err
	}

	server, ok := serverList[guild.ID]
	if !ok {
		return errors.New("guild did not exist in serverList")
	}

	for _, member := range guild.Members {
		if member.User.Mention() == root.CommandArgs()[0] {
			if err := server.UnMute(member.User.ID); err != nil {
				return err
			}

			if err := root.MessageSend(root.CommandArgs()[0]+" has been unmuted"); err != nil {
				return err
			}

			return root.DeleteMessage(root.ID)

		}
	}
	return errors.New("could not find user")
}
