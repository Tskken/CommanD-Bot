package main

/*
CommanD-Bot V0.6
Last Update: 11/20/2017
Author: Dylan Blanchard

main.go

Main program file.  Entry point
*/

import (
	// Golang imports //
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	// External imports //
	"github.com/bwmarrin/discordgo"
	"github.com/theskiier14/CommanD-Bot/CommanD"
)

// Current BotSession //
var botSession *discordgo.Session

// Pre-load and start bot //
// - Loads maps
// - Create new BotSession
func init() {
	// Load command maps //
	CommanD.Load()

	// Create new BotSession //
	botSession = CommanD.New("Bot <REDACTED>")
}

// Program entry point //
// - Waits for CTRL-C or terminate signal to close program
// - On terminate signal close bot session
func main() {
	// Wait until CTRL-C or other terminate signal is received. //
	fmt.Println("CommanD-Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Close bot session //
	err := botSession.Close()
	if err != nil {
		log.Println(err)
	}

	// End of main //
	log.Println("End of main.")
	return
}
