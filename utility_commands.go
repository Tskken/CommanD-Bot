package CommanD_Bot

import (
	"github.com/bwmarrin/discordgo"
)

// TODO - Comment
func UtilityCommands(s *discordgo.Session, m *discordgo.Message, admin bool) error {
	arg, err := ToLower(ParceInput(m.ChannelID), 1)
	if err != nil {
		return err
	}

	switch *arg{
	case "-dice":
		return DiceRole(s, m)
	case "-role":
		return DiceRole(s, m)
	case "-d":
		return DiceRole(s, m)
	case "-ign":
		return IGN(s, m)
	case "-ingamename":
		return IGN(s, m)
	case "-trinity":
		return Trinity(s, m)
	case "-t":
		return Trinity(s, m)
	default:
		_, err := s.ChannelMessageSend(m.ChannelID, *arg + "Is not a recognized option with in !util.  !help -util for a list of options that are supported.")
		return err
	}
	return nil
}

// TODO - Implement DiceRole
func DiceRole(s *discordgo.Session, m *discordgo.Message)error {
	return nil
}

// TODO - Implement IGN
func IGN(s *discordgo.Session, m *discordgo.Message)error {
	return nil
}

// TODO - Implement Trinity
func Trinity(s *discordgo.Session, m *discordgo.Message)error {
	return nil
}