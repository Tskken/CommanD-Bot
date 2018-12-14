package core

import (
	"encoding/json"
	"github.com/bwmarrin/discordgo"
	"os"
	"path/filepath"
)

// TODO: Figure out permission shit...

const filePath = "../CommanD-Bot/core/src/permissions.json"

type BotPermission []*Permissions

var BotPermissions *BotPermission

type Permissions struct {
	Name string `json:"name"`
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
	BotPermissions = new(BotPermission)
	return dec.Decode(BotPermissions)
}

func (p *Permissions) CheckUserPermission(member *discordgo.Member) bool {
	for _, r := range member.Roles {
		if p.Name == r {
			return true
		}
	}
	return false
}