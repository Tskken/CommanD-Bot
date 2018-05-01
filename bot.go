package CommanD_Bot

import (
	"bufio"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
)

// Create Bot info //
// - Returns error (nil if non)
func NewBot() (*discordgo.Session, error) {

	// Read in for bot token //
	r := bufio.NewReader(os.Stdin)
	log.Println("Enter Token: ")
	t, err := r.ReadString('\n')
	if err != nil {
		return nil, err
	}

	// Set token as given input //
	// - Ignores \n char ar the end of readline
	token := "Bot " + t[:len(t)-1]
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

	// Checks if the admin role exist in the guild //
	// Creates the role if it does not exist
	// Logs error if err is not nil
	if _, err := RoleCheck(s, g.Guild); err != nil {
		PrintError(err)
	}

	if _, ok := serverList[g.Guild.ID]; !ok {
		log.Println("running serverList add")
		s, err := newServer(g.Guild)
		if err != nil {
			PrintError(err)
		} else {
			serverList[g.Guild.ID] = s
		}
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

	args := ParceInput(m.Content)
	arg := StrToLower(args[0])

	if cmd, ok := botCommands[arg]; !ok {
		PrintError(NewError("Could not understand given command", "bot.go"))
	} else {
		err := RunCommand(s, m.Message, cmd)
		if err != nil {
			log.Println(err)
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

		// Load commands //
		loadCommands()

		// Load classifier data from file //
		if err := loadFilter(); err != nil {
			return nil, err
		}

		if err := LoadServer(); err != nil {
			return nil, err
		}

		// Open session //
		// Log error if err is not nil
		if err = session.Open(); err != nil {
			return nil, err
		}
		return session, nil
	}
}
