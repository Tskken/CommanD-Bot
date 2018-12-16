package core

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// TODO: Figure out permission shit...

const filePath = "../CommanD-Bot/core/src/permissions.json"

type BotPermission map[string]Permissions

var BotPermissions = make(BotPermission)

type Permissions struct {
	Name string `json:"name"`
	Permissions []string `json:"permissions"`
}

func NewPermissions() error {
	path, err := filepath.Abs(filePath)
	if err != nil {
		return err
	}
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	dec := json.NewDecoder(file)
	jsonData := make([]Permissions, 0)
	err = dec.Decode(&jsonData)
	if err != nil {
		return err
	}

	for _, data := range jsonData {
		BotPermissions[data.Name] = data
	}

	return nil
}