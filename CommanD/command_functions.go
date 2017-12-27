package CommanD

/*
Last Updated: 11/20/27
Author: Dylan Blanchard

commands_functions.go

Command functions
*/

import (
	// Golang imports //
	"log"
	"strconv"

	// External imports //
	"github.com/bwmarrin/discordgo"
)

// Returns all commands a user can use
// - s: discord server info
// - m: original discord message trigger
// TODO - FIX !help
// TODO - Refactor function layout
func Help(s *discordgo.Session, m *discordgo.Message, admin bool) error {
	args := ParceInput(m.Content)

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
}

// Deletes messages with in a channel //
// - s: discord server info
// - m: original message trigger
// - admin: user permission level
func DeleteMessage(s *discordgo.Session, m *discordgo.Message, admin bool) error {
	// Parce message content in to a string array //
	// - Parced on a space
	// - Returns a string array
	args := ParceInput(m.Content)


	// Delete Messages based on args[] //
	// - args[] len 1: delete the last message
	// -- Admin True: delete last message
	// -- Admin False: delete last message by sender of delete call
	// - args[] len 2: delete the last n number of messages were n is args[1] second argument passed
	// -- Admin True: delete last n messages
	// -- Admin False: delete last n messages by sender of delete call
	// - args[] len 3: delete last n number of messages by given user.  n args[1], username args[2].
	// -- Non-Admins have no use for this command.  Functions the same as deleting a given number of messages as they can not delete messages that are not there own.
	if len(args) == 1 {
		// Delete last messages //
		// - args[0] command call
		// - toDelete() returns a string[] of messages to delete
		return s.ChannelMessagesBulkDelete(m.ChannelID, toDelete(s, m, "", 1, admin))
	} else if len(args) == 2 {
		// args[] len 2 //
		// - args[0]: command
		// - args[1]: int or string
		// -- int: number of messages to delete
		// -- string: username of message to delete

		// Convert delete amount from string to int //
		// - i: amount to delete as int
		i, err := strconv.Atoi(args[1])
		if err != nil {
			// args[1] was not able to convert to an int //
			// - run delete for a given user on just one message
			// - toDelete() returns a string[] of messages to delete
			return s.ChannelMessagesBulkDelete(m.ChannelID, toDelete(s, m, args[1], 1, admin))
		}
		// args[1] was able to be converted to an int i //
		// - run delete on i number of messages
		// - toDelete() returns a string[] of messages to delete
		return s.ChannelMessagesBulkDelete(m.ChannelID, toDelete(s, m,"", i, admin))
	}

	// args[] is len 3 //
	// - args[0] command
	// - args[1] number to delete
	// - args[2] name of user to delete messages

	// Convert delete amount from string to int //
	// - i: amount to delete as int
	i, err := strconv.Atoi(args[1])
	if err != nil {
		log.Println(err)
	}
	// Delete messages //
	// - toDelete() returns a string[] of messages to delete
	return s.ChannelMessagesBulkDelete(m.ChannelID, toDelete(s, m, args[2], i, admin))
}

// Wrapper for deleteUserMessages //
// - s: discordgo session
// - m: disordgo messges.  Original message delete call
// - uName: username of messages to delete
// - i: number of messges to delete
// - admin: user admin role (true: has adin role | false: does not have admin role)
// - Returns list of message ID's to delete
func toDelete(s *discordgo.Session, m *discordgo.Message, uName string, i int, admin bool)[]string {
	// String array of message ID's to be deleted //
	toDelete := make([]string, 0)

	// get list of messages to delete for non-admin //
	// - returns a []string of i number of messages by m.Auther.Username
	if admin == false {
		return deleteUserMessages(s, m, m.Author.Username, i)
	}

	// get list of messages to delete for admin user with given Username //
	// - returns a []string of i number of messages by uName username
	if uName != "" {
		return deleteUserMessages(s, m, uName, i)
	}

	// No uName was given //
	// - run delete on i number of messages

	// get list of i number of messages plus 1 for original delete call //
	messages, err := s.ChannelMessages(m.ChannelID, i + 1, "", "", "")
	if err != nil {
		log.Println(err)
		return nil
	}

	// add messages to toDelete []string //
	for _, message := range messages {
		toDelete = append(toDelete, message.ID)
	}

	// Return list of messages to delete //
	return toDelete
}

// Find user messages to delete //
// - s: discordgo session
// - m: disordgo messges.  Original message delete call
// - uName: username of messages to delete
// - i: number of messges to delete
// - Returns list of message ID's to delete
func deleteUserMessages(s *discordgo.Session, m *discordgo.Message, uName string, i int)[]string{
	// String array of message ID's to be deleted //
	toDelete := make([]string, 0)

	// Current message ID //
	// - Set to first message
	current := m.ID

	// Loop while toDelete list is less then the number to be deleted //
	for len(toDelete) < i {
		// Get i number of messages from channel starting before current message ID//
		messages, err := s.ChannelMessages(m.ChannelID, i, current, "", "")
		if err != nil {
			log.Println(err)
			return nil
		}
		// If no messages were found //
		// - No more messages to delete in channel
		// - return nil error
		if len(messages) == 0 {
			log.Println("No more messages to delete")
			// If toDelete is empty //
			// - No messages were found in channel
			// - return nil error
			if len(toDelete) == 0 {
				log.Println("No messges in channel to delete")
				return nil
			}
		}

		// Move current message pointer to last message in messages //
		current = messages[len(messages)-1].ID

		// Loop through messages and add each message ID to toDelete list if uName matches message author username //
		for _, message := range messages {
			if message.Author.Username == uName {
				toDelete = append(toDelete, message.ID)
			}
		}
	}

	// Add delete message call to toDelete list //
	toDelete = append(toDelete, m.ID)

	// Return list of messages to delete //
	return toDelete
}

// TODO - Implement CreateChannel
func CreateChannel(s *discordgo.Session, m *discordgo.Message, admin bool) error {
	return nil
}

// TODO - Implement DeleteChannel
func DeleteChannel(s *discordgo.Session, m *discordgo.Message, admin bool)error {
	return nil
}

// TODO - Implement KickMember
func KickMember(s *discordgo.Session, m *discordgo.Message, admin bool)error {
	if admin == false {
		_, err := s.ChannelMessageSend(m.ChannelID, "You do not have the permission to kick someone.")
		return err
	}

	return nil
}

// TODO - Implement BanMember
func BanMember(s *discordgo.Session, m *discordgo.Message, admin bool)error {
	if admin == false {
		_, err := s.ChannelMessageSend(m.ChannelID, "You do not have the permission to kick someone.")
		return err
	}

	return nil
}