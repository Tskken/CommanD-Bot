package CommanD_Bot

import (
	"github.com/bwmarrin/discordgo"
)

type Root struct {
	*discordgo.Session
	*discordgo.Message
}
