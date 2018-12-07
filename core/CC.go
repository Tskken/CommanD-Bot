package core

import (
	"bufio"
	"github.com/bwmarrin/discordgo"
	"io"
	"log"
	"os"
	"path/filepath"
)

const tokenPath = "../CommanD-Bot/core/src/token"

type CentralCommand struct {
	*discordgo.Session
}

var CC *CentralCommand

func New() error {
	log.Println("getting token file path...")
	fPath, err := filepath.Abs(tokenPath)
	if err != nil {
		return err
	}

	log.Println("getting token file...")
	file, err := os.Open(fPath)
	if err != nil {
		return err
	}

	log.Println("creating io reader...")
	reader := bufio.NewReader(file)

	log.Println("reading token from file...")
	token, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return err
	}

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