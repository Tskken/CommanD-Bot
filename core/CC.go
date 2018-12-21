package core

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

type CentralCommand struct {
	*discordgo.Session
}

var CC *CentralCommand

func New(token string) error {
	log.Println("creating discordgo session...")
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return err
	}

	log.Println("binding session to CC...")
	CC = &CentralCommand{
		session,
	}

	log.Println("adding session handlers...")
	AddHandlers()

	log.Println("starting session...")
	return CC.Open()
}

func Close() error {
	log.Println("closing session...")
	return CC.Close()
}
