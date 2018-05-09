package CommanD_Bot

import (
	"bufio"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"strings"
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

	if session, err := discordgo.New(token); err != nil {
		return nil, err
	} else {
		// Set event handlers //
		session.AddHandler(guildCreate)
		session.AddHandler(messageCreate)

		// Load commands //
		loadCommands()

		// Load classifier data from file //
		if err := loadFilter(); err != nil {
			return nil, err
		}

		if err := loadServer(); err != nil {
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

func CloseBot(session *discordgo.Session) error {
	// Save filter classifiers when closing //
	err := saveFilter()
	if err != nil {
		log.Panic(err)
	}

	err = saveServer()
	if err != nil {
		log.Panic(err)
	}

	// Close bot session //
	if err := session.Close(); err != nil {
		log.Panic(err)
	}
	return nil
}

// GuildCreate event for when the bot joins a guild //
func guildCreate(session *discordgo.Session, create *discordgo.GuildCreate) {
	// Checks if the admin role exist in the guild //
	// Creates the role if it does not exist
	// Logs error if err is not nil
	if err := roleCheck(session, create.Guild); err != nil {
		log.Println(err)
	}

	if _, ok := serverList[create.Guild.ID]; !ok {
		log.Println("running serverList add")
		s, err := newServer()
		if err != nil {
			log.Println(err)
		} else {
			serverList[create.Guild.ID] = s
		}
	}
}

// MessageCreate event for when a message is sent with in a channel //
func messageCreate(session *discordgo.Session, create *discordgo.MessageCreate) {
	// Ignores all messages from the bot it self //
	if create.Author.ID == session.State.User.ID {
		return
	}

	/*
	TODO - Fix spam filter
	// Temp function call to teach algorithm //
	if err := scan(session, create.Message); err != nil {
		log.Println(err)
	}*/

	// Ignores all messages with out the ! key letter for commands //
	if create.Content[0] != '!' {
		return
	}

	args := strings.Fields(create.Content)
	arg := strings.ToLower(args[0])

	if cmd, ok := botCommands[arg]; !ok {
		log.Println("Could not understand given command")
	} else {
		err := runCommand(session, create.Message, cmd)
		if err != nil {
			log.Println(err)
		}
	}
}
