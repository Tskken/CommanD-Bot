package main

/*
CommanD-Bot V0.8
Author: Dylan Blanchard
*/

import (
	"github.com/tsukinai/CommanD-Bot"
	"github.com/tsukinai/CommanD-Bot/filter"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// On start //
func init() {
	// Create new Bot session //
	if b, err := CommanD_Bot.NewBot(); err != nil {
		log.Panic(err)
	} else {
		CommanD_Bot.Bot = b
	}
}

// Entry point //
func main() {
	// Bot is running //
	log.Println("Bot is now running...")
	// Wait for user input to close program //
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Save filter classifiers when closing //
	err := filter.Save(CommanD_Bot.Bot.GetClassifiers())
	if err != nil {
		log.Panic(err)
	}

	// Close bot session //
	if err := CommanD_Bot.Bot.GetSession().Close(); err != nil {
		log.Panic(err)
	}

	log.Println("Session ended.")
}
