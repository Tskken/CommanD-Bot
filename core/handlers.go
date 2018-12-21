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

func GuildCreate(session *discordgo.Session, create *discordgo.GuildCreate) {
	log.Println("checking guild permissions...")
	for key := range BotPermissions {
		exists := false
		for _, r := range create.Roles {
			if r.Name == key {
				exists = true
			}
		}

		if !exists {
			log.Println("role did not exist. creating role...")
			if role, err := session.GuildRoleCreate(create.ID); err != nil {
				log.Fatal(err)
			} else {
				// Set the new roles name to Admin and permissions to admin //
				// - returns an error if err is not nil
				_, err = session.GuildRoleEdit(create.ID, role.ID, key, 0, false, 0, false)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}

// TODO: Implement GuildDelete
func GuildDelete(_ *discordgo.Session, _ *discordgo.GuildDelete) {}

func MessageCreate(session *discordgo.Session, create *discordgo.MessageCreate) {
	// Ignores all messages from the bc it self //
	if create.Author.ID == CC.State.User.ID {
		return
	} else if !strings.HasPrefix(create.Content, "!") {
		return
	}

	parsedInput := ParseMessage(create.Content)

	err := NewCommand(session, create.Message, parsedInput).Run()
	if err != nil {
		log.Println(err)
	}
}
