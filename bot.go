package CommanD_Bot

import (
	"log"
	"github.com/bwmarrin/discordgo"
	"github.com/tsukinai/CommanD-Bot/botErrors"
	"github.com/tsukinai/CommanD-Bot/commands"
	"github.com/tsukinai/CommanD-Bot/servers"
	"github.com/tsukinai/CommanD-Bot/utility"
)

// Bot handler for GuildCreate Events //
func GuildCreate(s *discordgo.Session, g *discordgo.GuildCreate) {
	// Check for "terminal" text channel with in guild //
	if err := servers.ChannelCheck(s, g.Guild); err != nil {
		botErrors.PrintError(err)
	}

	// Check for the Bot role with in the server //
	if err := servers.CheckBotRole(s, g.Guild); err != nil {
		botErrors.PrintError(err)
	}

	// Check if "Admin" role exist in guild //
	if _, err := servers.RoleCheck(s, g.Guild); err != nil {
		botErrors.PrintError(err)
	}

	if _, ok := commands.BanTime[g.Name]; ok != true {
		commands.BanTime[g.Name] = 30
	}
}

// Bot handler for MessageCreate Events //
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Check if its a bot message //
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Checks if the command key was typed //
	if m.Content[0] != '!' {
		return
	}

	// Check if the user is an admin //
	admin := false


	// Get guild channel is in //
	guild, err := servers.GetGuild(s, m.Message)
	if err != nil {
		botErrors.PrintError(err)
	}

	// Get member of guild who sent original message //
	member, err := servers.GetMember(s, m.Message)
	if err != nil {
		botErrors.PrintError(err)
	}

	// Get admin role ID from guild //
	roleID, err := servers.RoleCheck(s, guild)
	if err != nil {
		botErrors.PrintError(err)
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
	arg, err := utility.ToLower(utility.ParceInput(m.Content), 0)
	if err != nil {
		botErrors.PrintError(err)
		return
	}

	// Check if command exist with in command map //
	if _, ok := commands.BotCommands[*arg]; ok != true {
		// Given command did not exist in map //
		info := "Command does not exist: " + *arg
		botErrors.PrintError(botErrors.NewError(info,"bot.go"))
		return
	}

	// Run given command //
	err = commands.BotCommands[*arg](s, m.Message, admin)
	if err != nil {
		botErrors.PrintError(err)
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

	// Load commands //
	commands.Load()

	// Opens session connection //
	err = session.Open()
	if err != nil {
		botErrors.PrintError(err)
	}

	// Returns session //
	return session
}
