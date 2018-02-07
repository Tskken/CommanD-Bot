package servers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tsukinai/CommanD-Bot/botErrors"
)

// - If it does not create it
// -- Set bot role to Bot-Bot
// - Returns an error (nil if non)
func CheckBotRole(s *discordgo.Session, g *discordgo.Guild) error {
	// Check if Bot role exist in guild roles //
	for _, role := range g.Roles {
		// Check role for "Bot" //
		// - If true, return nil
		if role.Name == "Bot" {
			return nil
		}
	}

	// Role was not found.  Create new role //
	role, err := s.GuildRoleCreate(g.ID)
	if err != nil {
		return err
	}
	// Set new roles values //
	_, err = s.GuildRoleEdit(g.ID, role.ID, "Bot", 15705102, true, 8, true)
	if err != nil {
		return err
	}

	// Set Bot role to Bot-Bot //
	// - Returns an error if role add errors (nil if non)
	for _, member := range g.Members {
		if member.User.Username == "Bot-Bot" {
			err = s.GuildMemberRoleAdd(g.ID, member.User.ID, role.ID)
			return err
		}
	}

	// Bot-Bot was not found with in guild //
	return botErrors.NewError("Bot was not found with in server","bot.go")
}

// Check if terminal text channel exists //
// - if it does not create it
// - Returns error (nil of non)
func ChannelCheck(s *discordgo.Session, g *discordgo.Guild) error {
	// For each channel with in guild check if command-prompt exists //
	// - channel: current channel struct in range
	for _, channel := range g.Channels {
		// Check if channel name is command-prompt //
		// - If true break.  command-prompt channel existed
		if channel.Name == "terminal" {
			return nil
		}
	}

	// channel was not found.  Create command-prompt channel //
	_, err := s.GuildChannelCreate(g.ID, "terminal", "text")
	return err
}

// Check if Admin role exists //
// - if it does not create it
// - Returns:
// -- RoleID as string
// -- error (nil if non)
func RoleCheck(s *discordgo.Session, g *discordgo.Guild)(string, error){
	// For each role with in guild check the role //
	// - role: Role struct of current role to check
	for _, role := range g.Roles {
		// Check if role name is Bot Master //
		if role.Name == "Admin" {
			return role.ID, nil
		}
	}

	// Role was not found.  Create new role //
	role, err := s.GuildRoleCreate(g.ID)
	if err != nil {
		return "", err
	}
	// Set new roles values //
	_, err = s.GuildRoleEdit(g.ID, role.ID, "Admin", 16724736, true, 8, true)
	return role.ID, err
}

// Wrapper to get guild channel //
// - s: Current session
// - m: Origin message for GetChannel request
// Returns:
// - discordgo Channel
// - error (nil if non)
func GetChannel(s *discordgo.Session, m *discordgo.Message)(*discordgo.Channel, error) {
	// Get Channel messege was sent in //
	return s.State.Channel(m.ChannelID)
}

// Wrapper to get guild //
// - s: Current session
// - c: Origin channel request was sent in
// Returns:
// - discordgo Guild
// - error (nil if non)
func GetGuild(s *discordgo.Session, m *discordgo.Message)(*discordgo.Guild, error) {
	// Get Guild the channel exist in //
	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		return nil,  err
	}
	return s.State.Guild(c.GuildID)
}

// Wrapper to get guild Member //
// - s: Current session
// - g: Guild member exist in
// - m: Origin message for GetMember request
// Returns:
// - discordgo Member
// - error (nil if non)
func GetMember(s *discordgo.Session, m *discordgo.Message)(*discordgo.Member, error) {
	// Get member who entered the message with in the guild channel //
	g, err := GetGuild(s, m)
	if err != nil {
		return nil, err
	}
	return s.GuildMember(g.ID, m.Author.ID)
}