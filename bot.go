package CommanD_Bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tsukinai/CommanD-Bot/botErrors"
	"github.com/tsukinai/CommanD-Bot/commands"
	"github.com/tsukinai/CommanD-Bot/servers"
	"github.com/tsukinai/CommanD-Bot/utility"
)

// GuildCreate event for when the bot joins a guild //
func GuildCreate(s *discordgo.Session, g *discordgo.GuildCreate) {
	// Checks if the terminal channel exist //
	// Creates the text channel if it does not exist
	// Logs error if err is not nil
	if err := servers.ChannelCheck(s, g.Guild); err != nil {
		botErrors.PrintError(err)
	}

	// Checks if the Bot role exist in the guild //
	// Creates the role if it does not exist
	// Logs error if err is not nil
	if err := servers.CheckBotRole(s, g.Guild); err != nil {
		botErrors.PrintError(err)
	}

	// Checks if the bot role exist in the guild //
	// Creates the role if it does not exist
	// Logs error if err is not nil
	if _, err := servers.RoleCheck(s, g.Guild); err != nil {
		botErrors.PrintError(err)
	}

	// Check if ban time is set for the server //
	// If not set ban time to 30 days by default //
	if _, ok := commands.BanTime[g.Name]; ok != true {
		commands.BanTime[g.Name] = 30
	}
}

// MessageCreate event for when a message is sent with in a channel //
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignores all messages from the bot it self //
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Ignores all messages with out the ! key letter for commands //
	if m.Content[0] != '!' {
		return
	}


	// Parse the messages on a space and set the first argument to lowercase //
	if arg, err := utility.ToLower(utility.ParceInput(m.Content), 0); err != nil {
		botErrors.PrintError(err)
		return
	} else {
		// Get the command from the command list for the first argument given //
		// Logs an error if the command does not exist
		if cmd, ok := commands.BotCommands[*arg]; !ok {
			// Given command did not exist in map //
			info := "Command does not exist: " + *arg
			botErrors.PrintError(botErrors.NewError(info, "bot.go"))
			return
		} else {
			if arg, err := utility.ToLower(utility.ParceInput(m.Content), 1); err != nil {
				botErrors.PrintError(err)
				return
			} else {
				if cmd, ok := cmd[*arg]; !ok {
					// Given command did not exist in map //
					info := "Command does not exist: " + *arg
					botErrors.PrintError(botErrors.NewError(info, "bot.go"))
					return
				} else {
					// Run command //
					// Logs error if err is not nil
					if err := cmd(s, m.Message); err != nil {
						botErrors.PrintError(err)
					}
				}
			}

		}
	}
}

// Creates a new discordgo.Session and loads all maps //
func New(token string) (*discordgo.Session, error) {
	// Creates new bot session with given token //
	// Logs an error if err is not nil
	if session, err := discordgo.New(token); err != nil {
		return nil, err
	} else {
		// Set event handlers //
		session.AddHandler(GuildCreate)
		session.AddHandler(MessageCreate)

		// Load all bot maps //
		commands.Load()

		// Open session //
		// Log error if err is not nil
		if err = session.Open(); err != nil {
			return nil, err
		} else {
			// Return a reference to the open session //
			return session, nil
		}
	}
}
