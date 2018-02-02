package CommanD_Bot

import (
	"log"
	"github.com/bwmarrin/discordgo"
)

var BanTime = make(map[string]int)

// Load command maps with bot commands //
// - BotCommands loaded with all commands
func Load() {
	// Load all commands in to botCommands map //
	loadBotCommands()
	LoadHelp()
}

func Save() {
	//saveBotMaps()
}

func saveBotMaps()error{
	encdec := NewEncDec()
	err := encdec.OpenFile()
	if err != nil {
		return err
	}
	encdec.NewEncGob()
	err = encdec.EncGob(BanTime)
	if err != nil {
		return err
	}
	err = encdec.CloseFile()
	return err
}

// Checks if bot role exist //
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
	return NewError("Bot was not found with in server","bot.go")
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
func GetGuild(s *discordgo.Session, c *discordgo.Channel)(*discordgo.Guild, error) {
	// Get Guild the channel exist in //
	return s.State.Guild(c.GuildID)
}

// Wrapper to get guild Member //
// - s: Current session
// - g: Guild member exist in
// - m: Origin message for GetMember request
// Returns:
// - discordgo Member
// - error (nil if non)
func GetMember(s *discordgo.Session, g *discordgo.Guild, m *discordgo.Message)(*discordgo.Member, error) {
	// Get member who entered the message with in the guild channel //
	return s.GuildMember(g.ID, m.Author.ID)
}

// Bot handler for GuildCreate Events //
// - Called when bot joins a guild or bot starts
// - Runs check for bot channel and admin role
func GuildCreate(s *discordgo.Session, g *discordgo.GuildCreate) {
	// Check for "terminal" text channel with in guild //
	// - If channel does not exist create it
	// - err: Error if channel create errors (nil if non)
	if err := ChannelCheck(s, g.Guild); err != nil {
		PrintError(err)
	}

	// Check for the Bot role with in the server //
	// - If role does not exist create it
	// - Give role to bot
	// - err: Error if role create errors (nil if non)
	if err := CheckBotRole(s, g.Guild); err != nil {
		PrintError(err)
	}

	// Check if "Admin" role exist in guild //
	// - If role does not exist create it
	// - err: Error if role create errors (nil if non)
	if _, err := RoleCheck(s, g.Guild); err != nil {
		PrintError(err)
	}

	if _, ok := BanTime[g.Name]; ok != true {
		BanTime[g.Name] = 30
	}
}

// Bot handler for MessageCreate Events //
// - Called when a message in the session server is made
// - Runs Commands based of message
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Check if its a bot message //
	// - Ignores messages created by itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Checks if the command key was typed //
	// - Ignores messages without an !
	if m.Content[0] != '!' {
		return
	}

	// Check if the user is an admin //
	// - Set admin to false by default
	admin := false

	// Get channel message was sent in //
	// - channel: discordgo channel
	// - err: error check (nil if non)
	channel, err := GetChannel(s, m.Message)
	if err != nil {
		PrintError(err)
	}

	// Get guild channel is in //
	// - guild: discordgo guild
	// - err: error check (nil if non)
	guild, err := GetGuild(s, channel)
	if err != nil {
		PrintError(err)
	}

	// Get member of guild who sent original message //
	// - member: discordgo member
	// - err: error check (nil if non)
	member, err := GetMember(s, guild, m.Message)
	if err != nil {
		PrintError(err)
	}

	// Get admin role ID from guild //
	// - If role does not exist, create it and return ID
	// - Error if there was an issue in creation of the role
	roleID, err := RoleCheck(s, guild)
	if err != nil {
		PrintError(err)
	}
	// Check members roles //
	for _, memRole := range member.Roles {
		// Check if members role is admin //
		if memRole == roleID {
			admin = true
			break
		}
	}

	// TODO - Fix comments
	// Parce message //
	// - Parces on a space
	// - Returns []string
	//args := ParceInput(m.Content)
	arg, err := ToLower(ParceInput(m.Content), 0)
	if err != nil {
		PrintError(err)
		return
	}

	// Check if command exist with in command map //
	// - ok: true or false if command exists
	// -- true: command exists
	// -- false: command does not exist
	if _, ok := BotCommands[*arg]; ok != true {
		// Given command did not exist in map //
		info := "Command does not exist: " + *arg
		PrintError(NewError(info,"bot.go"))
		return
	}

	// Run given command //
	// - args[0]: given command
	// - s: discordgo session
	// - m.Message: original message
	// - admin: user permission level
	err = BotCommands[*arg](s, m.Message, admin)
	if err != nil {
		PrintError(err)
	}
	return
}

// Create a new discordgo session //
func New(token string) *discordgo.Session {
	// Creates a new dicordgo session with token //
	session, err := discordgo.New(token)
	if err != nil {
		log.Println(err)
	}

	// Sets bot handlers //
	session.AddHandler(GuildCreate)
	session.AddHandler(MessageCreate)

	// Opens session connection //
	err = session.Open()
	if err != nil {
		PrintError(err)
	}

	// Returns session //
	return session
}
