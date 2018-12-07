package core

import (
	"github.com/bwmarrin/discordgo"
	"log"
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

// TODO: Implement MessageCreate
func MessageCreate(session *discordgo.Session, create *discordgo.MessageCreate) {
	// Ignores all messages from the bot it self //
	if create.Author.ID == CC.State.User.ID {
		return
	}

	command, err := ParseMessage(create.Content)
	if err != nil {
		log.Println(err)
		return
	} else if command == nil {
		return
	}
}
