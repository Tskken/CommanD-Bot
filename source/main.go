package main

/*
Bot-Bot V0.7
Author: Dylan Blanchard
*/

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tsukinai/CommanD-Bot"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Bot session //
var botSession *discordgo.Session

// On start //
func init() {
	// Create and start a new bot session. //
	// Log error if err is not nil
	if newSession, err := CommanD_Bot.New("Bot MzU3OTUwMTc3OTQ1OTc2ODM5.DOYtIQ.oa9Fqrl8RlhyunioLrmfItnpBkE"); err != nil {
		panic(err)
	} else {
		botSession = newSession
	}
}

// Entry point //
func main() {
	// Bot is running //
	log.Println("Bot is now running.  Press CTRL-C to exit.")
	// Wait for user input to close program //
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Close bot session //
	if err := botSession.Close(); err != nil {
		log.Println(err)
	}

	log.Println("Session ended.")
	return
}
