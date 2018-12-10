package core

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)
func AddHandlers() {
	CC.AddHandler(MessageCreate)
	CC.AddHandler(GuildCreate)
	CC.AddHandler(GuildDelete)
}

// TODO: Implement GuildCreate
func GuildCreate(session *discordgo.Session, create *discordgo.GuildCreate) {}

// TODO: Implement GuildDelete
func GuildDelete(session *discordgo.Session, delete *discordgo.GuildDelete) {}

func MessageCreate(session *discordgo.Session, create *discordgo.MessageCreate) {
	// Ignores all messages from the bot it self //
	if create.Author.ID == CC.State.User.ID {
		return
	} else if !strings.HasPrefix(create.Content, "!") {
		return
	}

	parsedInput, err := ParseMessage(create.Content)
	if err != nil {
		log.Println(err)
		return
	}

	err = NewCommand(session, create.Message, parsedInput).Run()
	if err != nil {
		log.Println(err)
	}
}
