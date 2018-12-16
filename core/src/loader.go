package main

import (
	"encoding/json"
	"github.com/Tskana/CommanD-Bot/bc"
	"github.com/Tskana/CommanD-Bot/core"
	"github.com/Tskana/CommanD-Bot/mc"
	"os"
	"path/filepath"
)

const (
	configPath = "../CommanD-Bot/core/src/config.json"

	Message = "message"
	Bot = "bot"
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

	dec := json.NewDecoder(file)

	type data struct {
		Command string `json:"command"`
		Keys []string `json:"keys"`
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
