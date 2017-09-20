package main

/*
CommanD-Bot V0.2
Last Update: 9/16/2017
Auther: Dylan Blanchard


*/

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

/*
func init() {

}*/

func getBotToken() string {
	return "Bot " + "MzU3OTUwMTc3OTQ1OTc2ODM5.DJxWsw.Mcv6ET5Hj73WYY7BsycYNQ10li0"
}

func callError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	_, err := s.ChannelMessageSend(m.ChannelID, m.Content)
	callError(err)
}

func main() {

	session, err := discordgo.New(getBotToken())
	callError(err)

	session.AddHandler(messageCreate)

	err = session.Open()
	callError(err)

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	session.Close()

	return
}
