package main

/*
CommanD-Bot V0.2
Last Update: 10/22/2017
Auther: Dylan Blanchard


*/

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/theskiier14/CommanD-Bot/CommanD"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	/*_, err := s.ChannelMessageSend(m.ChannelID,  m.Content)
		callError(err)*/
}

func main() {

	// Set Token ID //
	token := "Bot " + "MzU3OTUwMTc3OTQ1OTc2ODM5.DJxWsw.Mcv6ET5Hj73WYY7BsycYNQ10li0"

	// Create new discordgo instince //
	session, err := discordgo.New(token)
	if err != nil {
		log.Println(err)
	}
	// Set bot handlers //
	session.AddHandler(messageCreate)

	// Start Bot //
	err = session.Open()
	if err != nil {
		log.Println(err)
	}
	// Close Bot //
	defer session.Close()

	// Main Code //



	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Main ends //
	return
}
