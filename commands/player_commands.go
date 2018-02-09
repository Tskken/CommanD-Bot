package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tsukinai/CommanD-Bot/botErrors"
	"github.com/tsukinai/CommanD-Bot/servers"
	"github.com/tsukinai/CommanD-Bot/utility"
	"strconv"
)

// Wrapper function to run all !player commands //
// Can only be run by admin users
func PlayerCommands(s *discordgo.Session, m *discordgo.Message, admin bool) error {
	// Check caller is an admin //
	// Prints an error saying you do not have permission if you are not an admin
	if admin != true {
		_, err := s.ChannelMessageSend(m.ChannelID, "You do not have the permission to kick someone.")
		return err
	}

	// Gets the guild the messages was created in //
	// Returns an error if err is not nil
	if guild, err := servers.GetGuild(s, m); err != nil {
		return err
	} else {
		// gets the argument passed to !player and makes sure its lowercase //
		// Returns an error if err is nil
		if arg, err := utility.ToLower(utility.ParceInput(m.Content), 1); err != nil {
			return err
		} else {
			// Runs the !player argument //
			// Prints an error if the argument does not exist with in the list of supported arguments //
			if cmd, ok := playerCommands[*arg]; !ok {
				_, err := s.ChannelMessageSend(m.ChannelID, *arg+" is not a recognized option with in !player.  Type !help -player for a list of supported options.")
				return err
			} else {
				// Runs command //
				return cmd(s, m, guild)
			}
		}
	}
}

// Function to kick a user from the server //
// TODO - Fix kick with reason functionality
func kickMember(s *discordgo.Session, m *discordgo.Message, g *discordgo.Guild) error {
	// Parce the content of the message on a space //
	args := utility.ParceInput(m.Content)

	// Find the user with in the guild //
	for _, member := range g.Members {
		// If members nick name or username is the same  as the name passed in as the username argument kick them from the server //
		if member.Nick == args[2] || member.User.Username == args[2] {
			// Kick the user with out a reason //
			if len(args) == 3 {
				// Kick user //
				return s.GuildMemberDelete(g.ID, member.User.ID)
				// Kick a user with a reason //
			} else if len(args) >= 4 {
				// Kick user with reason //
				return s.GuildMemberDeleteWithReason(g.ID, member.User.ID, utility.ToString(args[3:]))
				// User did not enter a username or nick name to kick //
				// Print an error to the channel
			} else {
				_, err := s.ChannelMessageSend(m.ChannelID, "You did not enter a user to kick.  Type !help -kick for more info.")
				return err
			}
		}
	}
	// Return an error as the entered username or nick name was not found in the guild //
	return botErrors.NewError("Username or nick name was not found in guild.", "player_commands.go")
}

// Function to ban a user from the server for a given amount of time //
// TODO - Fix Ban with reason functionality
func banMember(s *discordgo.Session, m *discordgo.Message, guild *discordgo.Guild) error {
	// Parce the content of the message on a space //
	args := utility.ParceInput(m.Content)

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
					return s.GuildBanCreateWithReason(guild.ID, member.User.ID, utility.ToString(args[3:]), time)
				}
			}
			// User did not enter a username or nick name to kick //
			// Print an error to the channel
		} else {
			_, err := s.ChannelMessageSend(m.ChannelID, "You did not enter a user to ban.  Type !help -ban for more info.")
			return err
		}
	}
	return botErrors.NewError("If statement failed", "player_commands.go")
}

// Sets the servers ban time duration //
func newBanTimer(s *discordgo.Session, m *discordgo.Message, guild *discordgo.Guild) error {
	// Check if the server already as a time set //
	// If it does print it to the channel
	if time, ok := BanTime[guild.Name]; ok {
		if _, err := s.ChannelMessageSend(m.ChannelID, "Ban time was "+strconv.Itoa(time)); err != nil {
			return err
		}
	}

	// Parce the contents of the messages on a space //
	args := utility.ParceInput(m.Content)

	// If the there was no time given print an error to the server //
	if len(args) < 3 {
		_, err := s.ChannelMessageSend(m.ChannelID, "You need to give an amount of time to change the ban time to.")
		return err
		// The time was passed as an argument //
	} else {
		// Convert the amount of time to an int //
		// Returns an error if err is not nil
		if time, err := strconv.Atoi(args[2]); err != nil {
			return err
		} else {
			// Set the ban time for the new time given //
			BanTime[guild.Name] = time

			// Print the new ban time to the server //
			_, err = s.ChannelMessageSend(m.ChannelID, "Ban time has been set to "+strconv.Itoa(BanTime[guild.Name]))
			return err
		}
	}
}
