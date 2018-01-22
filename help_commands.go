package CommanD_Bot

/*
Last Updated: 11/20/27
Author: Dylan Blanchard

commands_functions.go


Command functions

TODO - REDO !help with SQLight database

TODO - Fix comments with in file
*/

import (
	// Golang imports //
	//"log"

	// External imports //
	//"github.com/bwmarrin/discordgo"
	//"github.com/tsukinai/Bot-Bot/Bot"
)

/*
// Returns all commands a user can use
// - s: discord server info
// - m: original discord message trigger
func Help(s *discordgo.Session, m *discordgo.Message, admin bool) error {
	args := Bot.ParceInput(m.Content)

	if len(args) == 1 {
		_, err := s.ChannelMessageSend(m.ChannelID, "fix help!!!")

		return err
	}

	if info, ok := HelpCommands[args[1]]; ok != true {
		log.Println("Command did not exist in HelpComamnds")
		return nil
	} else {
		var output string
		for _, help := range info {
			output += help.name + " "
			if len(help.args) != 0{
				for _, arg := range help.args {
					output += arg + " "
				}
			}

			output += " - "

			switch help.perm {
			case 0:
				if admin == true {
					break
				}
				output += help.info[0]
				break
			case 1:
				if admin != true {
					break
				}
				output += help.info[0]
				break
			case 2:
				if admin == true {
					output += help.info[0]
				} else {
					output += help.info[1]
				}
				break
			default:
				log.Println("error in help perm")
			}

			output += "\n"
		}
		_, err := s.ChannelMessageSend(m.ChannelID, output)
		return err

	}
	return nil
}*/

