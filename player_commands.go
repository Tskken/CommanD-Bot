package CommanD_Bot

import (
	"github.com/bwmarrin/discordgo"
)

type playerCommand commands

func (pc *playerCommand) command(s *discordgo.Session, m *discordgo.Message) error {
	args := ParceInput(m.Content)
	if len(args) < 2 {
		_, err := s.ChannelMessageSend(m.ChannelID, pc.commandInfo.Help())
		return err
	} else {
		return pc.subCommands[args[1]](s, m)
	}
}

func LoadPlayerCommand() *playerCommand {
	p := playerCommand{}
	p.commandInfo = loadPlayerCommandInfo()
	p.subCommands = make(map[string]func(*discordgo.Session, *discordgo.Message) error)
	p.subCommands["-kick"] = kickMember
	p.subCommands["-k"] = kickMember
	p.subCommands["-ban"] = banMember
	p.subCommands["-b"] = banMember

	return &p
}

// Create CommandInfo struct data //
func loadPlayerCommandInfo() *CommandInfo {

	p := &CommandInfo{}
	p.detail = "**!player** or **!pl** : All commands that pertain to manipulating players with in a server."
	p.commands = make(map[string]string)
	p.commands["-kick"] = "**-kick** or **-k**.\n**Info**: Kicks a given member from the server.\n" +
		"**Arguments:**\n		**<Member name>**: The member given will be kicked from the server. (only admins can use this command)."
	p.commands["-ban"] = "**-ban** or **-b**\n**Info**: Bans a given user from the server for a preset amount of time. (default is 30 days)." +
		"**Arguments:**\n	**<Member name>**: The member given will be baned from the server. (only admins can use this)."

	 return p
}


/*
// CommandInfo for PlayerCommands //
var PlayerCommandInfo *CommandInfo

// Set PlayerCommand info to struct //
func setPCInfo() {
	PlayerCommandInfo = &CommandInfo{}
	PlayerCommandInfo.detail = "**-player** or **-pl** : All commands that pertain to manipulating players with in a server."
	PlayerCommandInfo.commands = make(map[string]string)
	PlayerCommandInfo.commands["-kick"] = "**-kick** or **-k**.\n**Info**: Kicks a given member from the server.\n" +
		"**Arguments:**\n		**<Member name>**: The member given will be kicked from the server. (only admins can use this command)."
	PlayerCommandInfo.commands["-ban"] = "**-ban** or **-b**\n**Info**: Bans a given user from the server for a preset amount of time. (default is 30 days)." +
		"**Arguments:**\n	**<Member name>**: The member given will be baned from the server. (only admins can use this)."
}

// Get help for Player Commands //
// - Returns an error (nil if non)
func playerHelp(s *discordgo.Session, m *discordgo.Message) error {
	// Parce message on a space //
	ms := ParceInput(m.Content)

	// Check the number of arguments //
	// To few arguments given //
	if len(ms) < 2 {
		_, err := s.ChannelMessageSend(m.ChannelID, "Please enter a type of command you want help with.")
		return err
		// Get help for command //
	} else if len(ms) == 2 {
		_, err := s.ChannelMessageSend(m.ChannelID, PlayerCommandInfo.Help())
		return err
		// Get help for sub-command //
	} else if len(ms) == 3 {
		_, err := s.ChannelMessageSend(m.ChannelID, PlayerCommandInfo.HelpCommand(ms[2]))
		return err
	} else {
		_, err := s.ChannelMessageSend(m.ChannelID, "You gave to many arguments.")
		return err
	}
}*/

// Function to kick a user from the server //
// TODO - Fix kick with reason functionality
func kickMember(s *discordgo.Session, m *discordgo.Message) error {
	// Check for admin permitions //
	// Return an error if err is not nil
	admin, err := IsAdmin(s, m)
	if err != nil {
		return err
	}

	// Check caller is an admin //
	// Prints an error saying you do not have permission if you are not an admin
	if admin != true {
		_, err := s.ChannelMessageSend(m.ChannelID, "You do not have the permission to kick someone.")
		return err
	}

	// Gets the guild the messages was created in //
	// Returns an error if err is not nil
	if guild, err := GetGuild(s, m); err != nil {
		return err
	} else {
		// Parce the content of the message on a space //
		args := ParceInput(m.Content)

		// Get Username or Nick name //
		uName := ToString(args[2:], " ")

		// Find the user with in the guild //
		for _, member := range guild.Members {
			// If members nick name or username is the same  as the name passed in as the username argument kick them from the server //
			if member.Nick == uName || member.User.Username == uName {
				// Kick the user with out a reason //
				if len(args) == 3 {
					// Kick user //
					return s.GuildMemberDelete(guild.ID, member.User.ID)
					// Kick a user with a reason //
				} else if len(args) >= 4 {
					// Kick user with reason //
					return s.GuildMemberDeleteWithReason(guild.ID, member.User.ID, ToString(args[3:], " "))
					// User did not enter a username or nick name to kick //
					// Print an error to the channel
				} else {
					_, err := s.ChannelMessageSend(m.ChannelID, "You did not enter a user to kick.  Type !help -kick for more info.")
					return err
				}
			}
		}
	}

	// Return an error as the entered username or nick name was not found in the guild //
	return NewError("Username or nick name was not found in guild.", "player_commands.go")
}

// Function to ban a user from the server for a given amount of time //
// TODO - Fix Ban with reason functionality
func banMember(s *discordgo.Session, m *discordgo.Message) error {
	// Check if user is an admin //
	// - Returns an error if err is not nil
	admin, err := IsAdmin(s, m)
	if err != nil {
		return err
	}

	// Check caller is an admin //
	// Prints an error saying you do not have permission if you are not an admin
	if admin != true {
		_, err := s.ChannelMessageSend(m.ChannelID, "You do not have the permission to kick someone.")
		return err
	}

	// Gets the guild the messages was created in //
	// Returns an error if err is not nil
	if guild, err := GetGuild(s, m); err != nil {
		return err
	} else {
		// Parce the content of the message on a space //
		args := ParceInput(m.Content)

		// Get the ban time for a server //
		// Prints an error if the servers ban time as not been set
		if time, ok := BanTime[guild.Name]; !ok {
			_, err := s.ChannelMessageSend(m.ChannelID, "The ban time for the server has not been set.  Type !help -bantime for more info.")
			return err
		} else {
			// Ban the user with out a reason //
			if len(args) == 3 {
				// Find the user to ban with in the guild //
				for _, member := range guild.Members {
					// If members nick name or username is the same  as the name passed in as the username argument kick them from the server //
					if member.Nick == args[2] || member.User.Username == args[2] {
						// Ban the user for the servers set amount of time //
						return s.GuildBanCreate(guild.ID, member.User.ID, time)
					}
				}
				// Ban the user with a reason //
			} else if len(args) >= 4 {
				// Find the user to ban with in the guild //
				for _, member := range guild.Members {
					// If members nick name or username is the same  as the name passed in as the username argument kick them from the server //
					if member.Nick == args[2] || member.User.Username == args[2] {
						// Ban the user for the servers set amount of time with a reason //
						return s.GuildBanCreateWithReason(guild.ID, member.User.ID, ToString(args[3:], " "), time)
					}
				}
				// User did not enter a username or nick name to kick //
				// Print an error to the channel
			} else {
				_, err := s.ChannelMessageSend(m.ChannelID, "You did not enter a user to ban.  Type !help -ban for more info.")
				return err
			}
		}
	}
	return NewError("If statement failed", "player_commands.go")
}

// Sets the servers ban time duration //
func newBanTimer(s *discordgo.Session, m *discordgo.Message) error {
	// Checks if user is an admin //
	// - Returns an error if err is not nil
	admin, err := IsAdmin(s, m)
	if err != nil {
		return err
	}

	// Check caller is an admin //
	// Prints an error saying you do not have permission if you are not an admin
	if admin != true {
		_, err := s.ChannelMessageSend(m.ChannelID, "You do not have the permission to kick someone.")
		return err
	}

	// Gets the guild the messages was created in //
	// Returns an error if err is not nil
	if guild, err := GetGuild(s, m); err != nil {
		return err
	} else {
		// Check if the server already as a time set //
		// If it does print it to the channel
		if time, ok := BanTime[guild.Name]; ok {
			if _, err := s.ChannelMessageSend(m.ChannelID, "Ban time was "+IntToStr(time)); err != nil {
				return err
			}
		}

		// Parce the contents of the messages on a space //
		args := ParceInput(m.Content)

		// If the there was no time given print an error to the server //
		if len(args) < 3 {
			_, err := s.ChannelMessageSend(m.ChannelID, "You need to give an amount of time to change the ban time to.")
			return err
			// The time was passed as an argument //
		} else {
			// Convert the amount of time to an int //
			// Returns an error if err is not nil
			if time, err := StrToInt(args[2]); err != nil {
				return err
			} else {
				// Set the ban time for the new time given //
				BanTime[guild.Name] = time

				// Print the new ban time to the server //
				_, err = s.ChannelMessageSend(m.ChannelID, "Ban time has been set to "+IntToStr(BanTime[guild.Name]))
				return err
			}
		}
	}

}
