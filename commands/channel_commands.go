package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tsukinai/CommanD-Bot/utility"
)

// TODO - Comment
func ChannelCommands(s *discordgo.Session, m *discordgo.Message, admin bool) error {
	arg, err := utility.ToLower(utility.ParceInput(m.Content),1)
	if err != nil {
		return err
	}
	if cmd, ok := channelCommands[*arg]; !ok {
		_, err := s.ChannelMessageSend(m.ChannelID, *arg + " is not a recognized option with in !channel.  Type !help -channel for a list of supported options.")
		return err
	} else {
		return cmd(s, m, admin)
	}
}

// TODO - Implement CreateChannel
func createChannel(s *discordgo.Session, m *discordgo.Message, admin bool) error {return nil}

// TODO - Implement DeleteChannel
func deleteChannel(s *discordgo.Session, m *discordgo.Message, admin bool)error {return nil}