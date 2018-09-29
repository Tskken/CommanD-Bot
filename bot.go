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
	// Bot token //
	var token string

	// Creates a standard input reader //
	r := bufio.NewReader(os.Stdin)
	log.Println("Enter Token: ")
	// Read from stdin  till a new line char //
	// - returns an error if err is not nil
	if t, err := r.ReadString('\n'); err != nil {
		return nil, err
	} else {
		// Set token as given input //
		// - removes the new line char at the end of the input
		token = "Bot " + strings.TrimRight(t, "\n")
	}

	// Create discord session with token //
	// - returns an error if err is not nil
	if session, err := discordgo.New(token); err != nil {
		return nil, err
	} else {
		// Set event handlers //
		session.AddHandler(guildCreate)
		session.AddHandler(messageCreate)

		// Load commands //
		StartUp()

		/*// Load classifier data from file //
		// - returns an error if err is not nil
		if err := loadFilter(); err != nil {
			return nil, err
		}*/

		// Loads server data from file //
		// - return an error if err is not nil
		if err := loadServer(); err != nil {
			log.Println(err)
		}

		// Open session //
		// - returns error if err is not nil
		if err = session.Open(); err != nil {
			return nil, err
		}

		// Return bot session //
		return session, nil
	}
}

// Close bot session and save data to file //
// - Returns error (nil if non)
func CloseBot(session *discordgo.Session) error {
	// Save filter data to file //
	// - returns an error if err is not nil
	/*if err := saveFilter(); err != nil {
		return err
	}*/

	// Save server data to file //
	// - returns an error if err is not nil
	if err := saveServer(); err != nil {
		return err
	}

	// Close bot session //
	// - returns an error if err is not nil
	if err := session.Close(); err != nil {
		return err
	}

	return nil
}

// GuildCreate event handling when the bot joins a server //
func guildCreate(session *discordgo.Session, create *discordgo.GuildCreate) {
	// Checks if admin role exists in server //
	// - if it does not exist create it
	// - logs an error if err is not nil
	if err := roleCheck(session, create.Guild); err != nil {
		log.Println(err)
	}

	// Check if server exists in server list //
	// - if it does not exist create default data for the server
	if _, ok := serverList[create.Guild.ID]; !ok {
		log.Println("Creating new server data...")
		// Create new server data //
		// - logs an error if err is not nil
		if s, err := newServer(); err != nil {
			log.Println(err)
		} else {
			// Add new server data to server list //
			serverList[create.Guild.ID] = s
			log.Println(create.Guild.Name + " server data create...")
		}
	}
}

// MessageCreate event handling when a message is sent in a text channel //
func messageCreate(session *discordgo.Session, create *discordgo.MessageCreate) {
	// Ignores all messages from the bot it self //
	if create.Author.ID == session.State.User.ID {
		return
	}

	/*
		// Temp function call to teach algorithm //
		if err := scan(session, create.Message); err != nil {
			log.Println(err)
		}*/

	// TODO - Comment
	if muted, err := isMuted(session, create.Message); err != nil {
		log.Println(err)
	} else if muted {
		log.Println("user is muted")
		return
	}

	// Ignores all messages with out the ! char for commands //
	if !strings.HasPrefix(create.Content, "!") {
		return
	}

	// Parce input in to arguments and set them to lower case //
	args := toLower(strings.Fields(create.Content))

	var keys []CommandKey
	for key := range args[:2] {
		keys = append(keys, CommandKey(key))
	}

	root := RootCommand{session, create.Message, keys, args[2:]}

	// Run command struct //
	// - logs an error if err is not nil
	if err := Run(root); err != nil {
		log.Println(err)
	}

}
