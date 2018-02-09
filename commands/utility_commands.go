package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tsukinai/CommanD-Bot/utility"
)

// TODO - Comment
func UtilityCommands(s *discordgo.Session, m *discordgo.Message, admin bool) error {
	arg, err := utility.ToLower(utility.ParceInput(m.ChannelID), 1)
	if err != nil {
		return err
	}
	if cmd, ok := utilityCommands[*arg]; !ok {
		_, err := s.ChannelMessageSend(m.ChannelID, *arg+" is not a recognized option with in !utility.  Type !help -utility for a list of supported options.")
		return err
	} else {
		return cmd(s, m)
	}
}

// TODO - Implement DiceRole
func diceRole(s *discordgo.Session, m *discordgo.Message) error {
	return nil
}

// TODO - Implement IGN
func ign(s *discordgo.Session, m *discordgo.Message) error {
	return nil
}

// TODO - Implement Trinity
func trinity(s *discordgo.Session, m *discordgo.Message) error {
	return nil
}
