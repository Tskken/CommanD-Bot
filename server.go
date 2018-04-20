package CommanD_Bot

import (
	"github.com/bwmarrin/discordgo"
)

// Utility Maps //
var banTime = make(map[string]int)

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
			return NewError("Bot was not found with in server", "bot.go")
		}
	}
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

// Check if a user is an admin //
func IsAdmin(s *discordgo.Session, m *discordgo.Message) (bool, error) {
	// Sets admin to false by default //
	admin := false

	// Get the guild the message was sent in //
	// Logs an error if err is not nil
	guild, err := GetGuild(s, m)
	if err != nil {
		return admin, err
	}

	// Gets the member that created the message from the guild //
	// Logs an error if err is not nil
	member, err := GetMember(s, m)
	if err != nil {
		return admin, err
	}

	// Gets the admin role ID from the guild //
	// Creates the role and returns the new ID if it does not exist
	// Logs an error if err is not nil
	if roleID, err := RoleCheck(s, guild); err != nil {
		return admin, err
	} else {
		// Check members roles //
		for _, memRole := range member.Roles {
			// Member has admin role give them admin permissions //
			if memRole == *roleID {
				admin = true
				return admin, err
			}
		}
		return admin, err
	}
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
