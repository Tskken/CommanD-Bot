package CommanD_Bot

import (
	"github.com/bwmarrin/discordgo"
	"strconv"
)

// TODO - Comment
func MessageCommands(s *discordgo.Session, m *discordgo.Message, admin bool) error{
	arg, err := ToLower(ParceInput(m.Content), 1)
	if err != nil {
		return err
	} else if arg == nil {
		_, err := s.ChannelMessageSend(m.ChannelID, "There was no argument passed with !message.  Type !help -message to see supported options.")
		return err
	}

	switch *arg{
	case "-delete":
		return deleteMessage(s, m, admin)
	case "-del":
		return deleteMessage(s, m, admin)
	case "-clear":
		return clearMessages(s, m, admin)
	case "-cl":
		return clearMessages(s, m, admin)
	default:
		_, err := s.ChannelMessageSend(m.ChannelID, *arg + " is not an understood argument.  Type !help messages to get a list of commands.")
		return err
	}
	return NewError("Switch case failed", "message_commands.go")
}

// Deletes messages with in a channel //
// - s: discord server info
// - m: original message trigger
// - admin: user permission level
func deleteMessage(s *discordgo.Session, m *discordgo.Message, admin bool) error {
	// Parce message content in to a string array //
	// - Parced on a space
	// - Returns a string array
	args := ParceInput(m.Content)


	// Delete Messages based on args[] //
	// - args[] len 2: delete the last message
	// -- Admin True: delete last message
	// -- Admin False: delete last message by sender of delete call
	// - args[] len 3: delete the last n number of messages were n is args[1] second argument passed
	// -- Admin True: delete last n messages
	// -- Admin False: delete last n messages by sender of delete call
	// - args[] len 4: delete last n number of messages by given user.  n args[1], username args[2].
	// -- Non-Admins have no use for this command.  Functions the same as deleting a given number of messages as they can not delete messages that are not there own.
	if len(args) == 2 {
		// Delete last messages //
		// - args[0] command call
		// - toDelete() returns a string[] of messages to delete
		messages, err := toDelete(s, m, "", 1, admin)
		if err != nil {
			return err
		}
		return s.ChannelMessagesBulkDelete(m.ChannelID, messages)
	} else if len(args) == 3 {
		// args[] len 2 //
		// - args[0]: command
		// - args[1]: int or string
		// -- int: number of messages to delete
		// -- string: username of message to delete

		// Convert delete amount from string to int //
		// - i: amount to delete as int
		i, err := strconv.Atoi(args[2])
		if err != nil {
			// args[1] was not able to convert to an int //
			// - run delete for a given user on just one message
			// - toDelete() returns a string[] of messages to delete
			messages, err := toDelete(s, m, args[2], 1, admin)
			if err != nil {
				return err
			}
			return s.ChannelMessagesBulkDelete(m.ChannelID, messages)
		}
		// TODO - Comment
		if i >= 100 {
			_, err  = s.ChannelMessageSend(m.ChannelID, "You can only delete up to 99 messages at a time.")
			return err
		}
		// args[1] was able to be converted to an int i //
		// - run delete on i number of messages
		// - toDelete() returns a string[] of messages to delete
		messages, err := toDelete(s, m,"", i, admin)
		if err != nil {
			return err
		}
		return s.ChannelMessagesBulkDelete(m.ChannelID, messages)
	}

	// args[] is len 3 //
	// - args[0] command
	// - args[1] number to delete
	// - args[2] name of user to delete messages

	// Convert delete amount from string to int //
	// - i: amount to delete as int
	i, err := strconv.Atoi(args[2])
	if err != nil {
		return err
	}

	// TODO - Comment
	if i >= 100 {
		_, err  = s.ChannelMessageSend(m.ChannelID, "You can only delete up to 99 messages at a time.")
		return err
	}
	// Delete messages //
	// - toDelete() returns a string[] of messages to delete
	messages, err := toDelete(s, m, args[3], i, admin)
	if err != nil {
		return err
	}
	return s.ChannelMessagesBulkDelete(m.ChannelID, messages)
}

// Wrapper for deleteUserMessages //
// - s: discordgo session
// - m: disordgo messges.  Original message delete call
// - uName: username of messages to delete
// - i: number of messges to delete
// - admin: user admin role (true: has adin role | false: does not have admin role)
// - Returns list of message ID's to delete
func toDelete(s *discordgo.Session, m *discordgo.Message, uName string, i int, admin bool)([]string, error) {
	// String array of message ID's to be deleted //
	toDelete := make([]string, 0)

	// get list of messages to delete for non-admin //
	// - returns a []string of i number of messages by m.Auther.Username
	if admin == false {
		return getUserMessages(s, m, m.Author.Username, i)

	}

	// get list of messages to delete for admin user with given Username //
	// - returns a []string of i number of messages by uName username
	if uName != "" {
		return getUserMessages(s, m, uName, i)
	}

	// No uName was given //
	// - run delete on i number of messages

	// get list of i number of messages plus 1 for original delete call //
	messages, err := s.ChannelMessages(m.ChannelID, i + 1, "", "", "")
	if err != nil {
		return nil, err
	}

	// add messages to toDelete []string //
	for _, message := range messages {
		toDelete = append(toDelete, message.ID)
	}

	// Return list of messages to delete //
	return toDelete, nil
}

// Find user messages to delete //
// - s: discordgo session
// - m: disordgo messges.  Original message delete call
// - uName: username of messages to delete
// - i: number of messges to delete
// - Returns list of message ID's to delete
func getUserMessages(s *discordgo.Session, m *discordgo.Message, uName string, i int)([]string, error){
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
			return nil, err
		}
		// If no messages were found //
		// - No more messages to delete in channel
		// - return nil error
		if len(messages) == 0 {
			// If toDelete is empty //
			// - No messages were found in channel
			// - return nil error
			if len(toDelete) == 0 {
				return nil, NewError("There was no messages to delete with in given channel", "message_commands.go")
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
	return toDelete, nil
}

// TODO - Comment
func clearMessages(s *discordgo.Session, m *discordgo.Message, admin bool)error{
	if admin != true {
		_, err := s.ChannelMessageSend(m.ChannelID, "You do not have the permissions to use this command.")
		return err
	}
	ms, err := toDelete(s, m, "", 99, true)
	if err != nil {
		return err
	}
	return s.ChannelMessagesBulkDelete(m.ChannelID, ms)
}