package CommanD_Bot

import (
	"bufio"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
)

var token            string // Bot Token


// Create Bot info //
// - Returns error (nil if non)
func NewBot()(*discordgo.Session, error) {

	r := bufio.NewReader(os.Stdin)
	log.Println("Enter Token: ")
	t, err := r.ReadString('\n')
	if err != nil {
		return nil, err
	}

	token = "Bot " + t[:len(t)-1]
	// Create new sessin with bot token //
	// - Returns an error if err is not nil
	if b, err := New(token); err != nil {
		return nil, err
	} else {
		return b, nil
	}
}

// GuildCreate event for when the bot joins a guild //
func GuildCreate(s *discordgo.Session, g *discordgo.GuildCreate) {

	// Checks if the Bot role exist in the guild //
	// Creates the role if it does not exist
	// Logs error if err is not nil
	if err := CheckBotRole(s, g.Guild); err != nil {
		PrintError(err)
	}

	// Checks if the bot role exist in the guild //
	// Creates the role if it does not exist
	// Logs error if err is not nil
	if _, err := RoleCheck(s, g.Guild); err != nil {
		PrintError(err)
	}
}

func GuildDelete(s *discordgo.Session, g *discordgo.GuildDelete) {
	if err := RemoveBanTimer(g.Guild); err != nil {
		PrintError(err)
	}
}

// MessageCreate event for when a message is sent with in a channel //
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignores all messages from the bot it self //
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Temp function call to teach algorithm //
	if err := MScan(s, m.Message); err != nil {
		PrintError(err)
	}

	// Ignores all messages with out the ! key letter for commands //
	if m.Content[0] != '!' {
		return
	}
	//str := utility.ParceInput(m.Content)
	args := ParceInput(m.Content)
	arg := StrToLower(args[0])

	switch v := botCommands[arg].(type) {
	case *messageCommand:
		RunCommand(s, m.Message, v)
	case *channelCommand:
		RunCommand(s, m.Message, v)
	case *playerCommand:
		RunCommand(s, m.Message, v)
	case *utilityCommand:
		RunCommand(s, m.Message, v)
	default:
		log.Println("Error in new Command setup")

	}
}

// Creates a new discordgo.Session and loads all maps //
func New(token string) (*discordgo.Session, error)  {
	// Creates new bot session with given token //
	// Logs an error if err is not nil
	if session, err := discordgo.New(token); err != nil {
		return nil, err
	} else {
		// Set event handlers //
		session.AddHandler(GuildCreate)
		//session.AddHandler(GuildDelete)
		session.AddHandler(MessageCreate)

		// Load all bot maps //
		botCommands = LoadCommands()

		// Create classifiers //
		err := NewFilter()
		if err != nil {
			return nil, err
		}

		// Load classifier data from file //
		if err := LoadFilter(); err != nil {
			return nil, err
		}

		// Open session //
		// Log error if err is not nil
		if err = session.Open(); err != nil {
			return  nil, err
		}
		return session, nil
	}
}
