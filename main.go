package main
/*
Project CC 1.0
Author: Dylan Blanchard
Github: https://github.com/Tskana
*/

import (
	"fmt"
	"github.com/Tskana/CommanD-Bot/core"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	log.Println("loading token from file...")
	token, err := LoadToken()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("loading bot session...")
	err = core.New(token)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("loading message commands from config...")
	err = LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("loading permissions from config...")
	err = LoadPermissions()
	if err != nil {
		log.Fatal(err)
	}
}

// Entry point //
func main() {
	fmt.Println("Project CC 1.0")
	fmt.Println("running...")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	err := core.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("session closed")
}
