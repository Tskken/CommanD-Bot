package main

/*
CommanD-Bot Beta
Author: Dylan Blanchard
*/

import (
	"github.com/tsukinai/CommanD-Bot"
	//"github.com/tsukinai/CommanD-Bot/filter"
	//"github.com/tsukinai/CommanD-Bot/storage"
	//"github.com/tsukinai/CommanD-Bot/commands"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var bot *discordgo.Session

// On start //
func init() {
	//commands.Test()

	// Create new Bot session //
	if b, err := CommanD_Bot.NewBot(); err != nil {
		log.Println(err)
	} else {
		bot = b
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

	CommanD_Bot.CloseBot(bot)

	log.Println("Session ended.")
}
