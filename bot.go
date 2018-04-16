package CommanD_Bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jbrukh/bayesian"
	"github.com/tsukinai/CommanD-Bot/botErrors"
	"github.com/tsukinai/CommanD-Bot/commands"
	"github.com/tsukinai/CommanD-Bot/filter"
	"github.com/tsukinai/CommanD-Bot/servers"
	"github.com/tsukinai/CommanD-Bot/utility"
	"bufio"
	"os"
	"log"
)

var Bot *botInfo

// Bot info struct //
type botInfo struct {
	botSession       *discordgo.Session // DiscordGo Session
	filterClassifier *bayesian.Classifier // Bayesian Classifiers
	token            string // Bot Token
}

// Create Bot info //
// - Returns error (nil if non)
func NewBot() (*botInfo, error) {

	r := bufio.NewReader(os.Stdin)
	log.Println("Enter Token: ")
	token, err := r.ReadString('\n')
	if err != nil {
		return nil, err
	}

	// Create botInfo struct with given token //
	b := botInfo{nil, nil, "Bot "+token[:len(token)-1]}
	// Create new sessin with bot token //
	// - Returns an error if err is not nil
	if s, c, err := New(b.token); err != nil {
		return nil, err
	} else {
		// Set returned session to botSession //
		b.botSession = s
		// Set returned classifiers to filterClassifier //
		b.filterClassifier = c
		// Return botInfo reference //
		return &b, nil
	}
}

// Get discord session //
// - Returns a reference to the discordgo session
func (b *botInfo) GetSession() *discordgo.Session {
	return b.botSession
}

// Get filter classifiers ///
// - Return a reference to the bayesian classifiers
func (b *botInfo) GetClassifiers() *bayesian.Classifier {
	return b.filterClassifier
}

// GuildCreate event for when the bot joins a guild //
func GuildCreate(s *discordgo.Session, g *discordgo.GuildCreate) {

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
}

func GuildDelete(s *discordgo.Session, g *discordgo.GuildDelete) {
	if err := servers.RemoveBanTimer(g.Guild); err != nil {
		botErrors.PrintError(err)
	}
}

// MessageCreate event for when a message is sent with in a channel //
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignores all messages from the bot it self //
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Temp function call to teach algorithm //
	if err := filter.MScan(s, m.Message, Bot.GetClassifiers()); err != nil {
		botErrors.PrintError(err)
	}

	// Ignores all messages with out the ! key letter for commands //
	if m.Content[0] != '!' {
		return
	}
	//str := utility.ParceInput(m.Content)
	args := utility.ParceInput(m.Content)
	arg := utility.StrToLower(args[0])

	// Get the command from the command list for the first argument given //
	// Logs an error if the command does not exist
	if cmd, ok := commands.BotCommands[arg]; !ok {
		// Given command did not exist in map //
		info := "Command does not exist: " + arg
		botErrors.PrintError(botErrors.NewError(info, "bot.go"))
		return
	} else {
		//str := utility.ParceInput(m.Content)

		if len(args) < 2 {
			botErrors.PrintError(botErrors.NewError("To few arguments in command call", "bot.go"))
			return
		}

		sArg := utility.StrToLower(args[1])

		if cmd, ok := cmd[sArg]; !ok {
			// Given command did not exist in map //
			info := "Command does not exist: " + sArg
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

// Creates a new discordgo.Session and loads all maps //
func New(token string) (*discordgo.Session, *bayesian.Classifier, error) {
	// Creates new bot session with given token //
	// Logs an error if err is not nil
	if session, err := discordgo.New(token); err != nil {
		return nil, nil, err
	} else {
		// Set event handlers //
		session.AddHandler(GuildCreate)
		//session.AddHandler(GuildDelete)
		session.AddHandler(MessageCreate)

		// Load all bot maps //
		commands.Load()

		// Create classifiers //
		classifier, err := filter.NewFilter()
		if err != nil {
			return nil, nil, err
		}

		// Load classifier data from file //
		if err := filter.Load(classifier); err != nil {
			return nil, nil, err
		}

		// Open session //
		// Log error if err is not nil
		if err = session.Open(); err != nil {
			return nil, classifier, err
		} else {
			// Return a reference to the open session //
			return session, classifier, nil
		}
	}
}
