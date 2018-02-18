package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/tsukinai/CommanD-Bot/botErrors"
	"github.com/tsukinai/CommanD-Bot/servers"
	"github.com/tsukinai/CommanD-Bot/utility"
	"time"
	"log"
)

/*// Wrapper to run all message commands through the messageCommands map //
func MessageCommands(s *discordgo.Session, m *discordgo.Message, admin bool) error {
	// Parces message content and gives the argument passed to !message set to lowercase //
	// Prints an error to the server if there was no argument passed to !message
	if arg, err := utility.ToLower(utility.ParceInput(m.Content), 1); err != nil {
		_, err := s.ChannelMessageSend(m.ChannelID, "There was no argument passed with !message.  Type !help -message to see supported options.")
		return err
	} else {
		// Get command from messageCommands map //
		// Prints an error to the server if the argument is not in the map
		if cmd, ok := messageCommands[*arg]; !ok {
			_, err := s.ChannelMessageSend(m.ChannelID, *arg+" is not an understood argument.  Type !help messages to get a list of commands.")
			return err
		} else {
			// Run command //
			return cmd(s, m, admin)
		}
	}
}*/

// Main message delete function //
func deleteMessage(s *discordgo.Session, m *discordgo.Message) error {
	admin, err := servers.IsAdmin(s, m)
	if err != nil {
		return err
	}

	// Parce message content on a space //
	args := utility.ParceInput(m.Content)

	// Check the length of args for possible passed arguments //
	switch len(args) {
	// -del was called with no arguments //
	// Deletes the last message sent in server or by user if user is not an admin
	case 2:
		// Get messages to delete //
		// Returns an error if err is not nil
		if messages, err := toDelete(s, m, "", 1, admin); err != nil {
			return err
		} else {
			// Delete messages //
			return s.ChannelMessagesBulkDelete(m.ChannelID, messages)
		}
	// -del was called with either a number of messages or a user name //
	// If its a number it deletes that number of messages in that channel or by that user if the user is not an admin
	// If its a name it deletes the last message by that given user name
	case 3:
		// Try's to convert second argument to an integer //
		if i, err := utility.StrToInt(args[2]); err != nil {
			// Argument two was a username //
			// Get last message sent by given user to delete //
			// Returns an error if err is not nil
			if messages, err := toDelete(s, m, args[2], 1, admin); err != nil {
				return err
			} else {
				// Delete given messages //
				return s.ChannelMessagesBulkDelete(m.ChannelID, messages)
			}
		} else {
			// Argument two was a number //
			// Checks to make sure the given number is not less then or equal to 0 and not greater then 99 //
			// Prints an error if not true
			if i >= 100 || i <= 0 {
				_, err := s.ChannelMessageSend(m.ChannelID, "You entered a number that I don't understand. Please enter a number between 1-99.")
				return err
			}

			// Gets last given number of messages to delete.  Gets messages by user if not admin //
			// Returns an error if err is not nil
			if messages, err := toDelete(s, m, "", i, admin); err != nil {
				return err
			} else {
				// Delete given messages //
				return s.ChannelMessagesBulkDelete(m.ChannelID, messages)
			}
		}
	// -del was called with both a number to delete and a username //
	// If you are an admin you can not use this command
	default:
		// Check if the user is an admin //
		// if not print an error and return
		if admin == false {
			_, err := s.ChannelMessageSend(m.ChannelID, "You can not use this command.  If you want to multipal of your own messagse just type !ms -del <number of messages>.")
			return err
		}

		// Try and convert the second argument to an integer //
		// Returns an error is err is not nil
		if i, err := utility.StrToInt(args[2]); err != nil {
			return err
		} else {
			// Checks to make sure the given number is not less then or equal to 0 and not greater then 99 //
			// Prints an error if not true
			if i >= 100 || i <= 0 {
				_, err := s.ChannelMessageSend(m.ChannelID, "You entered a number that I don't understand. Please enter a number between 1-99.")
				return err
			}
			// Gets messages by user to be deleted //
			// Returns an error if err is not nil
			if messages, err := toDelete(s, m, args[3], i, true); err != nil {
				return err
			} else {
				// Deletes messages //
				return s.ChannelMessagesBulkDelete(m.ChannelID, messages)
			}
		}
	}
}

// Get messages from channel to delete //
func toDelete(s *discordgo.Session, m *discordgo.Message, uName string, i int, admin bool) ([]string, error) {
	// List of message ID's to delete //
	toDelete := make([]string, 0)

	// If user is not an admin only get messages by that given user //
	if admin == false {
		// Get messages from that user and return them //
		return getUserMessages(s, m, m.Author.Username, i)

	}

	// Check if a username is given //
	if uName != "" {
		// Get messages by given user name //
		return getUserMessages(s, m, uName, i)
	}

	// No username was given /
	// Get given number of messages to delete //
	// Returns an error if err is not nil
	if messages, err := s.ChannelMessages(m.ChannelID, i+1, "", "", ""); err != nil {
		return nil, err
	} else {
		// Adds each message ID's to toDelete list //
		for _, message := range messages {
			if t, err := message.Timestamp.Parse(); err != nil {
				return nil, err
			}else {
				log.Println("Entered Time check")
				time := time.Now()
				if t.Year() != time.Year() || (t.YearDay() + 14) <= time.YearDay() {
					log.Println("Entered year day check")
					_, err := s.ChannelMessageSend(m.ChannelID, "No more messages to delete in this channel newer then 14 days.")
					if err != nil {
						return nil, err
					}
					return toDelete, botErrors.NewError("No more messages newer then 14 days.", "messages_commands.go")
				}
			}
			toDelete = append(toDelete, message.ID)
		}

		// Check if the list is zero //
		// If it is zero no messages were added to list
		// Return an error
		if len(toDelete) == 0 {
			return nil, botErrors.NewError("There was no messages to delete.", "message_commands.go")
		}

		// Return list of messages to delete //
		return toDelete, nil
	}
}

// Find user messages to delete //
func getUserMessages(s *discordgo.Session, m *discordgo.Message, uName string, i int) ([]string, error) {
	// List of message ID's to delete //
	toDelete := make([]string, 0)

	// Current message ID //
	current := m.ID

	// Loop while toDelete list is less then the number of messages to be deleted //
	for len(toDelete) < i {
		// Get the given number of messages from the channel //
		// Returns an error if err is not nil
		if messages, err := s.ChannelMessages(m.ChannelID, i, current, "", ""); err != nil {
			return nil, err
		} else {
			// If there was no messages in the channel to begin with, return an error //
			if len(messages) == 0 && len(toDelete) == 0 {
				return nil, botErrors.NewError("There was no messages to delete with in given channel", "message_commands.go")
			}

			// Move current message pointer to the last message ID with in messages //
			current = messages[len(messages)-1].ID

			// Loop through messages and add each message ID to toDelete list if uName matches the username or nick name of the messages creator //
			for _, message := range messages {
				if t, err := message.Timestamp.Parse(); err != nil {
					return nil, err
				}else {
					time := time.Now()
					if t.Year() == time.Year() && (t.YearDay() + 14) <= time.YearDay() {
						ms, err := s.ChannelMessageSend(m.ChannelID, "No more messages to delete in this channel newer then 14 days.")
						if err != nil {
							return nil, err
						}
						toDelete = append(toDelete, ms.ID)
						return toDelete, botErrors.NewError("No more messages newer then 14 days.", "messages_commands.go")
					}
				}
				// Gets the user info of the message //
				// Returns an error if err is nil
				if user, err := servers.GetMember(s, message); err != nil {
					return nil, err
					// If the nick name of the message creator is equal to uName add the message ID to the list toDelete //
				} else if user.Nick == uName {
					toDelete = append(toDelete, message.ID)
					// If the username of the message creator is equal to uName add the message ID to the list toDelete //
				} else if message.Author.Username == uName {
					toDelete = append(toDelete, message.ID)
				}
			}
		}
	}

	// Add the initial !message -delete <args> call to the toDelete list //
	toDelete = append(toDelete, m.ID)

	// Return list of messages to delete //
	return toDelete, nil
}

// Clear all messages newer then 2 weeks from a channel //
func clearMessages(s *discordgo.Session, m *discordgo.Message) error {
	admin, err := servers.IsAdmin(s, m)
	if err != nil {
		return err
	}

	// This command can only be used by admin users //
	// If a non-admin calls the function it prints an error stating you do not have the permissions to use this command
	if admin != true {
		_, err := s.ChannelMessageSend(m.ChannelID, "You do not have the permissions to use this command.")
		return err
	}

	// Clear message Loop //
	// Runs toDelete to get the list of messages to delete in batches of 99
	// Loop ends if the list of messages returned from toDelete is 0 or it returns an error
	for ms, _ := toDelete(s, m, "", 99, true); len(ms) != 0 ; ms, _ = toDelete(s, m, "", 99, true) {
		// Deletes the current batch of 99 messages //
		// Returns an error if err is not nil
		if err := s.ChannelMessagesBulkDelete(m.ChannelID, ms); err != nil {
			return err
		}
	}

	// Clear was successful //
	return nil
}
