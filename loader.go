package main

import (
	"bufio"
	"encoding/json"
	"github.com/Tskana/CommanD-Bot/bc"
	"github.com/Tskana/CommanD-Bot/core"
	"github.com/Tskana/CommanD-Bot/mc"
	"io"
	"log"
	"os"
	"path/filepath"
)

const (
	configPath = "../CommanD-Bot/config.json"
	permPath   = "../CommanD-Bot/permissions.json"
	tokenPath  = "../CommanD-Bot/token"

	Message = "message"
	Bot     = "bot"
)

func LoadConfig() error {
	path, err := filepath.Abs(configPath)
	if err != nil {
		return err
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	dec := json.NewDecoder(file)

	type data struct {
		Command string   `json:"command"`
		Keys    []string `json:"keys"`
	}

	dat := make([]data, 0)

	err = dec.Decode(&dat)
	if err != nil {
		return err
	}

	for _, d := range dat {
		switch d.Command {
		case Message:
			core.AddCommand(new(mc.MessageCommand), d.Keys...)
		case Bot:
			core.AddCommand(new(bc.BotCommand), d.Keys...)
		}
	}
	return nil
}

func LoadPermissions() error {
	path, err := filepath.Abs(permPath)
	if err != nil {
		return err
	}
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	dec := json.NewDecoder(file)
	jsonData := make([]core.Permissions, 0)
	err = dec.Decode(&jsonData)
	if err != nil {
		return err
	}

	for _, data := range jsonData {
		core.BotPermissions[data.Name] = data
	}

	return nil
}

func LoadToken() (string, error) {
	fPath, err := filepath.Abs(tokenPath)
	if err != nil {
		return "", err
	}

	file, err := os.Open(fPath)
	if err != nil {
		return "", err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	reader := bufio.NewReader(file)

	token, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}

	return token, nil
}
