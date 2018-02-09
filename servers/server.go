package servers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tsukinai/CommanD-Bot/botErrors"
)

// Checks the server to make sure it has the Bot role with in it and sets it to the bot //
func CheckBotRole(s *discordgo.Session, g *discordgo.Guild) error {
	// Check all roles in the server //
	for _, role := range g.Roles {
		// If Bot role is found return //
		if role.Name == "Bot" {
			return nil
		}
	}

	// Role was not found with in the server //
	// Create a new role //
	// Returns an error if err is not nil
	if role, err := s.GuildRoleCreate(g.ID); err != nil {
		return err
	} else {
		// Set the roles name to Bot, color, permissions to admin, and let users mention the role //
		// Returns an error if err is not nil
		if _, err = s.GuildRoleEdit(g.ID, role.ID, "Bot", 15705102, true, 8, true); err != nil {
			return err
		} else {
			// Looks for the bot with in the server //
			for _, member := range g.Members {
				//  Set the Bot role to CommanD-Bot //
				// Returns an error if err is not nil
				if member.User.Username == "CommanD-Bot" {
					err = s.GuildMemberRoleAdd(g.ID, member.User.ID, role.ID)
					return err
				}
			}
			// Returns an error as the bot was not found with in the server ... some how ... //
			return botErrors.NewError("Bot was not found with in server", "bot.go")
		}
	}
}

// Check to make sure the terminal channel exist with in the guild //
func ChannelCheck(s *discordgo.Session, g *discordgo.Guild) error {
	// Look though guilds channels //
	for _, channel := range g.Channels {
		// If the terminal exist return //
		if channel.Name == "terminal" {
			return nil
		}
	}

	// terminal did not exist.  Create the terminal channel as a text channel //
	// Returns an error if err is not nil
	_, err := s.GuildChannelCreate(g.ID, "terminal", "text")
	return err
}

// Check if the Admin role is with in the guild and create it if not //
func RoleCheck(s *discordgo.Session, g *discordgo.Guild) (*string, error) {
	// Look though guild roles //
	for _, role := range g.Roles {
		// Admin role exist and return the role ID //
		if role.Name == "Admin" {
			return &role.ID, nil
		}
	}

	// Create a new role with in the guild //
	// Returns an error if err is not nil
	if role, err := s.GuildRoleCreate(g.ID); err != nil {
		return nil, err
	} else {
		// Set the new roles name to Admin, permissions to admin //
		// Return the new roles ID
		// Returns an error if err is not nil
		_, err = s.GuildRoleEdit(g.ID, role.ID, "Admin", 16724736, true, 8, true)
		return &role.ID, err
	}
}

// Gets a channel from with in a guild //
func GetChannel(s *discordgo.Session, m *discordgo.Message) (*discordgo.Channel, error) {
	// Returns a channel //
	// Returns error (nil if non | not nil if error) //
	return s.State.Channel(m.ChannelID)
}

// Gets guild info for the given guild //
func GetGuild(s *discordgo.Session, m *discordgo.Message) (*discordgo.Guild, error) {
	// Get the channel the messages is from with in the guild //
	// Returns an error if err is not nil
	if c, err := s.State.Channel(m.ChannelID); err != nil {
		return nil, err
	} else {
		// Returns the guild info //
		// Returns an error (nil if non | not nil if error) //
		return s.State.Guild(c.GuildID)
	}
}

// Gets a member from with in a guild //
func GetMember(s *discordgo.Session, m *discordgo.Message) (*discordgo.Member, error) {
	// Gets the guild that the member should be in //
	// Returns an error if err is not nil
	if g, err := GetGuild(s, m); err != nil {
		return nil, err
	} else {
		// Returns a member //
		// Returns an error (nil if non | not nil if error)
		return s.GuildMember(g.ID, m.Author.ID)
	}
}
