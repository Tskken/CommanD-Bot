package CommanD_Bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jbrukh/bayesian"
	"github.com/tsukinai/CommanD-Bot/botErrors"
	"github.com/tsukinai/CommanD-Bot/commands"
	"github.com/tsukinai/CommanD-Bot/filter"
	"github.com/tsukinai/CommanD-Bot/servers"
	"github.com/tsukinai/CommanD-Bot/utility"
)

var Bot *botInfo

type botInfo struct {
	botSession       *discordgo.Session
	filterClassifier *bayesian.Classifier
	token            string
}

func NewBot() (*botInfo, error) {
	b := botInfo{nil, nil, "Bot MzU3OTUwMTc3OTQ1OTc2ODM5.DOYtIQ.oa9Fqrl8RlhyunioLrmfItnpBkE"}
	if s, c, err := New(b.token); err != nil {
		return nil, err
	} else {
		b.botSession = s
		b.filterClassifier = c
		return &b, nil
	}
}

func (b *botInfo) GetSession() *discordgo.Session {
	return b.botSession
}

func (b *botInfo) GetClassifiers() *bayesian.Classifier {
	return b.filterClassifier
}

// GuildCreate event for when the bot joins a guild //
func GuildCreate(s *discordgo.Session, g *discordgo.GuildCreate) {
	/*
		// Checks if the terminal channel exist //
		// Creates the text channel if it does not exist
		// Logs error if err is not nil
		if err := servers.ChannelCheck(s, g.Guild); err != nil {
			botErrors.PrintError(err)
		}*/

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

	// TODO - Comment
	if err := commands.SetBanTimer(g.Guild); err != nil {
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

	/*
		if !Bot.filterClassifier.DidConvertTfIdf {
			if err := filter.CvTFIDF(Bot.filterClassifier); err != nil {
				botErrors.PrintError(err)
			}
		}*/

	// Ignores all messages with out the ! key letter for commands //
	if m.Content[0] != '!' {
		return
	}
	//str := utility.ParceInput(m.Content)
	arg := utility.StrToLower(utility.ParceInput(m.Content)[0])

	// Get the command from the command list for the first argument given //
	// Logs an error if the command does not exist
	if cmd, ok := commands.BotCommands[arg]; !ok {
		// Given command did not exist in map //
		info := "Command does not exist: " + arg
		botErrors.PrintError(botErrors.NewError(info, "bot.go"))
		return
	} else {
		//str := utility.ParceInput(m.Content)
		arg := utility.StrToLower(utility.ParceInput(m.Content)[1])

		if cmd, ok := cmd[arg]; !ok {
			// Given command did not exist in map //
			info := "Command does not exist: " + arg
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
		session.AddHandler(MessageCreate)

		// Load all bot maps //
		commands.Load()
		classifier, err := filter.NewFilter()
		if err != nil {
			return nil, nil, err
		}
		filter.Load(classifier)

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
