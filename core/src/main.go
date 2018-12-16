package main

/*
Project CC 1.0
Author: Dylan Blanchard
Github: https://github.com/Tskana
*/

import (
	"github.com/Tskana/CommanD-Bot/core"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	err := core.New()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("adding message commands...")
	err = LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
}

// Entry point //
func main() {
	log.Println("Project CC 1.0")
	log.Println("running...")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	err := core.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("session closed")
}
