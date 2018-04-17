package CommanD_Bot

import (
	"github.com/bwmarrin/discordgo"
	"time"
)

// Command Info struct for MessageCommands //
//var MessageCommandInfo *CommandInfo

type messageCommand commands

func (mc *messageCommand) command(s *discordgo.Session, m *discordgo.Message) error {
	args := ParceInput(m.Content)
	if len(args) < 2 {
		_, err := s.ChannelMessageSend(m.ChannelID, mc.commandInfo.Help())
		return err
	} else {
		return mc.subCommands[args[1]](s, m)
	}
}

func LoadMessageCommand() *messageCommand {
	m := messageCommand{}
	m.commandInfo = loadMessageCommandInfo()
	m.subCommands = make(map[string]func(*discordgo.Session, *discordgo.Message) error)
	m.subCommands["-delete"] = deleteMessage
	m.subCommands["-del"] = deleteMessage
	m.subCommands["-clear"] = clearMessages
	m.subCommands["-cl"] = clearMessages

	return &m
}

// Create CommandInfo struct data //
func loadMessageCommandInfo() *CommandInfo {
	m := &CommandInfo{}
	m.detail = "**!message** or **!ms** : All commands that pertain to manipulating messages with in a server."
	m.commands = make(map[string]string)
	m.commands["-delete"] = "**-delete** or **-del**.\n**Info**: Deletes a given messages with in a server.\n" +
		"**Arguments:**\n		**none**: Deletes the last messages in the chanel by you.\n		**<Number between 1 - 99>:**" +
		" Deletes the last given number of messages by you. \n		**<User Name>:** Deletes the last message by that user (only can be used by admin)." +
		"\n		**<Number between 1 - 99> <User Name>:** Deletes the last given number of messages by the given user (only can be used by admin)."
	return m
}

/*
// Get help info for given message //
// Returns an error (nil if non)
func messageHelp(s *discordgo.Session, m *discordgo.Message) error {
	// parce message on a space //
	ms := utility.ParceInput(m.Content)

	// Check the number of arguments //
	// To few arguments passed //
	if len(ms) < 2 {
		_, err := s.ChannelMessageSend(m.ChannelID, "Please enter a type of command you want help with.")
		return err
	// Get help for given command //
	} else if len(ms) == 2 {
		_, err := s.ChannelMessageSend(m.ChannelID, MessageCommandInfo.Help())
		return err
	// Get help for given sub-command //
	} else if len(ms) == 3 {
		_, err := s.ChannelMessageSend(m.ChannelID, MessageCommandInfo.HelpCommand(ms[2]))
		return err
	// Number of arguments was to large //
	} else {
		_, err := s.ChannelMessageSend(m.ChannelID, "You gave to many arguments.")
		return err
	}
}*/

// Main message delete function //
func deleteMessage(s *discordgo.Session, m *discordgo.Message) error {
	admin, err := IsAdmin(s, m)
	if err != nil {
		return err
	}

	// Parce message content on a space //
	args := ParceInput(m.Content)

	// Check the length of args for possible passed arguments //
	switch len(args) {
	// -del was called with no arguments //
	// Deletes the last message sent in server or by user if user is not an admin
	case 2:
		// Get messages to delete //
		// Returns an error if err is not nil
		if messages, err := ToDelete(s, m, "", 1, admin); err != nil {
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
		if i, err := StrToInt(args[2]); err != nil {
			// Argument two was a username //
			// Get last message sent by given user to delete //
			// Returns an error if err is not nil
			if messages, err := ToDelete(s, m, args[2], 1, admin); err != nil {
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
				_, err := s.ChannelMessageSend(m.ChannelID, "You entered a number that I don't understand. "+
					"Please enter a number between 1-99.")
				return err
			}

			// Gets last given number of messages to delete.  Gets messages by user if not admin //
			// Returns an error if err is not nil
			if messages, err := ToDelete(s, m, "", i, admin); err != nil {
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
			_, err := s.ChannelMessageSend(m.ChannelID, "You can not use this command.  "+
				"If you want to multiple of your own messages just type !ms -del <number of messages>.")
			return err
		}

		// Try and convert the second argument to an integer //
		// Returns an error is err is not nil
		if i, err := StrToInt(args[2]); err != nil {
			return err
		} else {
			// Checks to make sure the given number is not less then or equal to 0 and not greater then 99 //
			// Prints an error if not true
			if i >= 100 || i <= 0 {
				_, err := s.ChannelMessageSend(m.ChannelID, "You entered a number that I don't understand. "+
					"Please enter a number between 1-99.")
				return err
			}
			// Gets messages by user to be deleted //
			// Returns an error if err is not nil
			if messages, err := ToDelete(s, m, args[3], i, true); err != nil {
				return err
			} else {
				// Deletes messages //
				return s.ChannelMessagesBulkDelete(m.ChannelID, messages)
			}
		}
	}
}

// Get messages from channel to delete //
func ToDelete(s *discordgo.Session, m *discordgo.Message, uName string, i int, admin bool) ([]string, error) {
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
			if msTime, err := message.Timestamp.Parse(); err != nil {
				return nil, err
			} else {
				if IsTime(msTime, time.Now()) {
					return toDelete, nil
				}
			}
			toDelete = append(toDelete, message.ID)
		}

		// Check if the list is zero //
		// If it is zero no messages were added to list
		// Return an error
		if len(toDelete) == 0 {
			return nil, NewError("There was no messages to delete.", "message_commands.go")
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
				return nil, NewError("There was no messages to delete with in given channel", "message_commands.go")
			}

			// Move current message pointer to the last message ID with in messages //
			current = messages[len(messages)-1].ID

			// Loop through messages and add each message ID to toDelete list if uName matches the username or nick name of the messages creator //
			for _, message := range messages {
				if msTime, err := message.Timestamp.Parse(); err != nil {
					return nil, err
				} else {
					if IsTime(msTime, time.Now()) {
						return toDelete, nil
					}
				}
				// Gets the user info of the message //
				// Returns an error if err is nil
				if user, err := GetMember(s, message); err != nil {
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
	if admin, err := IsAdmin(s, m); err != nil {
		return err
	} else if !admin {
		// This command can only be used by admin users //
		// If a non-admin calls the function it prints an error stating you do not have the permissions to use this command
		_, err := s.ChannelMessageSend(m.ChannelID, "You do not have the permissions to use this command.")
		return err
	} else {
		// Clear message Loop //
		// Runs toDelete to get the list of messages to delete in batches of 99
		// Loop ends if the list of messages returned from toDelete is 0 or it returns an error
		for ms, err := ToDelete(s, m, "", 99, true); len(ms) != 0; ms, err = ToDelete(s, m, "", 99, true) {
			if err != nil {
				return err
			}
			// Deletes the current batch of 99 messages //
			// Returns an error if err is not nil
			if err := s.ChannelMessagesBulkDelete(m.ChannelID, ms); err != nil {
				return err
			}
		}

		// Clear was successful //
		return nil
	}
}
