package CommanD_Bot

/*
import (
	"github.com/bwmarrin/discordgo"
	"github.com/tsukinai/CommanD-Bot/utility"
)

func setBad(s *discordgo.Session, m *discordgo.Message) error {
	ms := utility.ParceInput(m.Content)

	if len(ms) > 3 {
		_, err := s.ChannelMessageSend(m.ChannelID, "Did not fallow proper format.  Type !help -filter for more info.")
		return err
	}

	if v, ok := KeyWordMap[ms[2]]; !ok {
		KeyWordMap[ms[2]] = true
	} else {
		if v == true {
			_, err := s.ChannelMessageSend(m.ChannelID, "That word is already flagged as 'bad'.")
			return err
		} else {
			KeyWordMap[ms[2]] = true
		}
	}
	return nil
}

func setGood(s *discordgo.Session, m *discordgo.Message) error {
	ms := utility.ParceInput(m.Content)

	if len(ms) > 3 {
		_, err := s.ChannelMessageSend(m.ChannelID, "Did not fallow proper format.  Type !help -filter for more info.")
		return err
	}

	if v, ok := KeyWordMap[ms[2]]; !ok {
		KeyWordMap[ms[2]] = false
	} else {
		if v == false {
			_, err := s.ChannelMessageSend(m.ChannelID, "That word is already flagged as 'good'.")
			return err
		} else {
			KeyWordMap[ms[2]] = false
		}
	}
	return nil
}
*/
