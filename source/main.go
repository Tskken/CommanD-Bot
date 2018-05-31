package main

/*
CommanD-Bot Beta
Author: Dylan Blanchard
*/

import (
	"github.com/tsukinai/CommanD-Bot"

	"log"
	"os"
	"os/signal"
	"syscall"
)

// Entry point //
func main() {
	// Create new Bot session //
	// - logs an error if err is not nil
	if bot, err := CommanD_Bot.NewBot(); err != nil {
		log.Println(err)
	} else {
		// Bot is running //
		log.Println("Bot is now running...")
		// Wait for user input to close program //
		sc := make(chan os.Signal, 1)
		signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
		<-sc

		// Close bot session and save data to files //
		// - logs an error if err is not nil
		if err := CommanD_Bot.CloseBot(bot); err != nil {
			log.Println(err)
		}

		// Bot session ended //
		log.Println("Session ended.")
	}

}
