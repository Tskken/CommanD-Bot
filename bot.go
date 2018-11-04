package CommanD_Bot

import (
	"bufio"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"strings"
)

type Bot struct {
	*discordgo.Session
	*BotCommands
}

func (b *Bot) SetCommands() *Bot {
	botCommands := &BotCommands{}

	mc := LoadMessageCommand()
	botCommands.commands["!messages"] = mc
	botCommands.commands["!ms"] = mc

	cc := LoadChannelCommand()
	botCommands.commands["!channel"] = cc
	botCommands.commands["!ch"] = cc

	gc := LoadGuildCommand()
	botCommands.commands["!guild"] = gc
	botCommands.commands["!gl"] = gc

	pc := LoadPlayerCommand()
	botCommands.commands["!player"] = pc
	botCommands.commands["!pl"] = pc

	uc := LoadUtilityCommand()
	botCommands.commands["!utility"] = uc
	botCommands.commands["!ut"] = uc

	b.BotCommands = botCommands

	return b
}

// Create Bot info //
// - Returns error (nil if non)
func NewBot() (*Bot, error) {
	b := &Bot{}

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
		b.Session = session

		b.AddHandlers()

		// Load commands //
		b.SetCommands()

		/*// Load classifier data from file //
		// - returns an error if err is not nil
		if err := loadFilter(); err != nil {
			return nil, err
		}*/

		// Loads server data from file //
		// - return an error if err is not nil
		if err := LoadServer(); err != nil {
			log.Println(err)
		}

		// Open session //
		// - returns error if err is not nil
		if err = b.Open(); err != nil {
			return nil, err
		}

		// Return bot session //
		return b, nil
	}
}

// Close bot session and save data to file //
// - Returns error (nil if non)
func (b *Bot) CloseBot() error {
	// Save filter data to file //
	// - returns an error if err is not nil
	/*if err := saveFilter(); err != nil {
		return err
	}*/

	// Save server data to file //
	// - returns an error if err is not nil
	if err := SaveServer(); err != nil {
		return err
	}

	// Close bot session //
	// - returns an error if err is not nil
	if err := b.Close(); err != nil {
		return err
	}

	return nil
}

func (b *Bot) AddHandlers() {
	b.AddHandler(GuildCreate)
	b.AddHandler(MessageCreate)
}

// GuildCreate event handling when the bot joins a server //
func GuildCreate(session *discordgo.Session, create *discordgo.GuildCreate) {
	// Checks if admin role exists in server //
	// - if it does not exist create it
	// - logs an error if err is not nil
	if err := RoleCheck(session, create.Guild); err != nil {
		log.Println(err)
	}

	// Check if server exists in server list //
	// - if it does not exist create default data for the server
	if _, ok := serverList[create.Guild.ID]; !ok {
		log.Println("Creating new server data...")
		// Create new server data //
		// - logs an error if err is not nil
		if s, err := NewServer(); err != nil {
			log.Println(err)
		} else {
			// Add new server data to server list //
			serverList[create.Guild.ID] = s
			log.Println(create.Guild.Name + " server data create...")
		}
	}
}

// MessageCreate event handling when a message is sent in a text channel //
func MessageCreate(session *discordgo.Session, create *discordgo.MessageCreate) {
	// Ignores all messages from the bot it self //
	if create.Author.ID == session.State.User.ID {
		return
	}

	r := &Root{
		session,
		create.Message,
	}

	// TODO - Comment
	if muted, err := r.IsMuted(); err != nil {
		log.Println(err)
	} else if muted {
		log.Println("user is muted")
		return
	}

	/*
		// Temp function call to teach algorithm //
		if err := scan(session, create.Message); err != nil {
			log.Println(err)
		}*/

	// Ignores all messages with out the ! char for commands //
	if !strings.HasPrefix(create.Content, "!") {
		return
	}

	// Run command struct //
	// - logs an error if err is not nil
	if err := r.Run(); err != nil {
		log.Println(err)
	}
}
