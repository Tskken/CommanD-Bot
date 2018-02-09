package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tsukinai/CommanD-Bot/botErrors"
	"github.com/tsukinai/CommanD-Bot/servers"
	"github.com/tsukinai/CommanD-Bot/utility"
	"strconv"
)

// TODO - Comment
func MessageCommands(s *discordgo.Session, m *discordgo.Message, admin bool) error {
	arg, err := utility.ToLower(utility.ParceInput(m.Content), 1)
	if err != nil {
		return err
	} else {
		_, err := s.ChannelMessageSend(m.ChannelID, "There was no argument passed with !message.  Type !help -message to see supported options.")
		return err
	}

	if cmd, ok := messageCommands[*arg]; !ok {
		_, err := s.ChannelMessageSend(m.ChannelID, *arg+" is not an understood argument.  Type !help messages to get a list of commands.")
		return err
	} else {
		return cmd(s, m, admin)
	}
}

// Deletes messages with in a channel //
func deleteMessage(s *discordgo.Session, m *discordgo.Message, admin bool) error {
	// Parce message content in to a string array //
	args := utility.ParceInput(m.Content)

	switch len(args) {
	case 2:
		// Delete last messages //
		if messages, err := toDelete(s, m, "", 1, admin); err != nil {
			return err
		} else {
			return s.ChannelMessagesBulkDelete(m.ChannelID, messages)
		}
	case 3:
		// Convert delete amount from string to int //
		if i, err := strconv.Atoi(args[2]); err != nil {
			if messages, err := toDelete(s, m, args[2], 1, admin); err != nil {
				return err
			} else {
				return s.ChannelMessagesBulkDelete(m.ChannelID, messages)
			}
		} else {
			if i >= 100 || i <= 0 {
				_, err := s.ChannelMessageSend(m.ChannelID, "You entered a number that I don't understand. Please enter a number between 1-99.")
				return err
			}

			if messages, err := toDelete(s, m, "", i, admin); err != nil {
				return err
			} else {
				return s.ChannelMessagesBulkDelete(m.ChannelID, messages)
			}
		}
	default:
		// Convert delete amount from string to int //
		if i, err := strconv.Atoi(args[2]); err != nil {
			return err
		} else {
			if i >= 100 || i <= 0 {
				_, err := s.ChannelMessageSend(m.ChannelID, "You entered a number that I don't understand. Please enter a number between 1-99.")
				return err
			}
			// Delete messages //
			if messages, err := toDelete(s, m, args[3], i, admin); err != nil {
				return err
			} else {
				return s.ChannelMessagesBulkDelete(m.ChannelID, messages)
			}
		}
	}
}

// Wrapper for deleteUserMessages //
func toDelete(s *discordgo.Session, m *discordgo.Message, uName string, i int, admin bool) ([]string, error) {
	// String array of message ID's to be deleted //
	toDelete := make([]string, 0)

	// get list of messages to delete for non-admin //
	if admin == false {
		return getUserMessages(s, m, m.Author.Username, i)

	}

	// get list of messages to delete for admin user with given Username //
	if uName != "" {
		return getUserMessages(s, m, uName, i)
	}

	// No uName was given //

	// get list of i number of messages plus 1 for original delete call //
	if messages, err := s.ChannelMessages(m.ChannelID, i+1, "", "", ""); err != nil {
		return nil, err
	} else {
		// add messages to toDelete []string //
		for _, message := range messages {
			toDelete = append(toDelete, message.ID)
		}

		if len(toDelete) == 0 {
			return nil, botErrors.NewError("There was no messages to delete.", "message_commands.go")
		}

		// Return list of messages to delete //
		return toDelete, nil
	}
}

// Find user messages to delete //
func getUserMessages(s *discordgo.Session, m *discordgo.Message, uName string, i int) ([]string, error) {
	// String array of message ID's to be deleted //
	toDelete := make([]string, 0)

	// Current message ID //
	current := m.ID

	// Loop while toDelete list is less then the number to be deleted //
	for len(toDelete) < i {
		// Get i number of messages from channel starting before current message ID//
		if messages, err := s.ChannelMessages(m.ChannelID, i, current, "", ""); err != nil {
			return nil, err
		} else {
			// If no messages were found //
			if len(messages) == 0 && len(toDelete) == 0 {
				return nil, botErrors.NewError("There was no messages to delete with in given channel", "message_commands.go")
			}

			// Move current message pointer to last message in messages //
			current = messages[len(messages)-1].ID

			// Loop through messages and add each message ID to toDelete list if uName matches message author username //
			for _, message := range messages {
				if user, err := servers.GetMember(s, message); err != nil {
					return nil, err
				} else if user.Nick == uName {
					toDelete = append(toDelete, message.ID)
				} else if message.Author.Username == uName {
					toDelete = append(toDelete, message.ID)
				}
			}
		}
	}

	// Add delete message call to toDelete list //
	toDelete = append(toDelete, m.ID)

	// Return list of messages to delete //
	return toDelete, nil
}

// TODO - Comment
func clearMessages(s *discordgo.Session, m *discordgo.Message, admin bool) error {
	if admin != true {
		_, err := s.ChannelMessageSend(m.ChannelID, "You do not have the permissions to use this command.")
		return err
	}

	for ms, err := toDelete(s, m, "", 99, true); err == nil && len(ms) != 0; ms, err = toDelete(s, m, "", 99, true) {
		if err = s.ChannelMessagesBulkDelete(m.ChannelID, ms); err != nil {
			return err
		}
	}

	return nil
}
