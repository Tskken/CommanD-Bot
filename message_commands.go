package CommanD_Bot

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"strconv"
	"time"
)

// Set message command structure //
func loadMessageCommand() *commands {
	// Create message command structure //
	m := commands{}

	// Load message command info structure //
	m.commandInfo = loadMessageCommandInfo()

	// Create sub command map //
	m.subCommands = make(map[string]func(*discordgo.Session, *discordgo.Message, []string) error)

	// Set message delete function //
	m.subCommands["-delete"] = deleteMessage
	m.subCommands["-del"] = deleteMessage

	// Set message clear function //
	m.subCommands["-clear"] = clearMessages
	m.subCommands["-cl"] = clearMessages

	// Return reference to command structure //
	return &m
}

// Create CommandInfo structure //
func loadMessageCommandInfo() *commandInfo {
	// Create message command info structure //
	m := commandInfo{}

	// Set default info //
	m.detail = "**!message** or **!ms** : All commands that pertain to manipulating messages with in a server."

	// Create sub command info map //
	m.commands = make(map[string]string)

	// Set info for message delete function //
	m.commands["-delete"] = "**-delete** or **-del**.\n" +
		"**Info**: Deletes a given messages with in a server.\n" +
		"**Arguments:**\n" +
		"    **none**: Deletes the last messages in the chanel by you.\n" +
		"    **<Number between 1 - 99>:** Deletes the last given number of messages by you.\n" +
		"    **<User Name>:** Deletes the last message by that user (only can be used by admin).\n" +
		"    **<Number between 1 - 99> <User Name>:** Deletes the last given number of messages by the given user (only can be used by admin)."

	// Set info for clear message function //
	m.commands["-clear"] = "**-clear** or **-cl**.\n" +
		"**Info**: Deletes all message with in a channel newer then 14 days."

	// Returns a reference to command info structure //
	return &m
}

// Message delete function //
// - returns an error (nil if non)
func deleteMessage(session *discordgo.Session, message *discordgo.Message, args []string) error {
	// Gets user admin status //
	// - returns error if err is not nil
	// - true: is admin
	// - false: not admin
	admin, err := isAdmin(session, message)
	if err != nil {
		return err
	}

	// Check the length of args for possible passed arguments //
	switch len(args) {
	// -del was called with no arguments //
	// Deletes the last message sent in server or by user if user is not an admin
	case 2:
		// Get messages to delete //
		// - returns an error if err is not nil
		if messages, err := toDelete(session, message, "", 1, admin); err != nil {
			return err
		} else {
			// Delete messages //
			// - returns error (nil if non)
			return session.ChannelMessagesBulkDelete(message.ChannelID, messages)
		}
	// -del was called with either a number of messages or a user name //
	// - Number: deletes that number of messages in that channel or by that user if the user is not an admin
	// - UserName: deletes the last message by that given user name (only can be used by admins)
	case 3:
		// Try's to convert second argument to an integer //
		if i, err := strconv.Atoi(args[2]); err != nil {
			if !admin {
				// - returns an error (nil if non)
				_, err := session.ChannelMessageSend(message.ChannelID, "You are not an admin, you can not delete message by a specific user.")
				return err
			}
			// Argument two was a username //
			// - get last message sent by given user to delete
			// - returns an error if err is not nil
			if messages, err := toDelete(session, message, args[2], 1, admin); err != nil {
				return err
			} else {
				// Delete given messages //
				// - returns an error (nil if non)
				return session.ChannelMessagesBulkDelete(message.ChannelID, messages)
			}
		} else {
			// Argument two was a number //
			// - checks given number is between 1 and 99
			if i >= 100 || i <= 0 {
				// Send message to channel saying you do not have permission for this action //
				// - returns an error (nil if non)
				_, err := session.ChannelMessageSend(message.ChannelID, "You entered a number that I don't understand. Please enter a number between 1-99.")
				return err
			}

			// Gets last given number of messages to delete.  Gets messages by user if not admin //
			// - returns an error if err is not nil
			if messages, err := toDelete(session, message, "", i, admin); err != nil {
				return err
			} else {
				// Delete given messages //
				// - returns an error (nil if non)
				return session.ChannelMessagesBulkDelete(message.ChannelID, messages)
			}
		}
	// -del was called with both a number to delete and a username //
	// - can not be called by non admin users
	default:
		// Check if the user is an admin //
		if !admin {
			// Send message saying user was not an admin //
			// - return an error (nil if non)
			_, err := session.ChannelMessageSend(message.ChannelID, "You can not use this command. If you want to multiple of your own messages just type !ms -del <number of messages>.")
			return err
		}

		// Try and convert the second argument to an integer //
		// - returns an error is err is not nil
		if i, err := strconv.Atoi(args[2]); err != nil {
			return err
		} else {
			// Checks number is with in range of 1 and 99 //
			if i >= 100 || i <= 0 {
				// Send message to channel for value being out of range //
				// - return an error (nil if non)
				_, err := session.ChannelMessageSend(message.ChannelID, "You entered a number that I don't understand. Please enter a number between 1-99.")
				return err
			}
			// Gets messages by user to be deleted //
			// - returns an error if err is not nil
			if messages, err := toDelete(session, message, args[3], i, true); err != nil {
				return err
			} else {
				// Deletes messages //
				// - returns an error (nil if non)
				return session.ChannelMessagesBulkDelete(message.ChannelID, messages)
			}
		}
	}
}

// Get messages from channel to delete //
// - returns an array of strings and an error (nil if non)
func toDelete(session *discordgo.Session, message *discordgo.Message, uName string, i int, admin bool) ([]string, error) {
	// Create list of message to delete //
	toDelete := make([]string, 0)

	// Check if user is an admin //
	if !admin {
		// User was not admin, get only messages create by them //
		// - returns list of message by user
		// - returns error (nil if non)
		return getUserMessages(session, message, message.Author.Username, i)

	}

	// Check if a username is given //
	if uName != "" {
		// Get messages by given user name //
		// - returns list of messages by given user
		// - returns an error (nil if non)
		return getUserMessages(session, message, uName, i)
	}

	// Get given number of messages //
	// - returns an error if err is not nil
	if messages, err := session.ChannelMessages(message.ChannelID, i+1, "", "", ""); err != nil {
		return nil, err
	} else {
		// Adds each message ID to toDelete list //
		for _, m := range messages {
			// Get time when message was created //
			// - returns an error if err is not nil
			if msTime, err := m.Timestamp.Parse(); err != nil {
				return nil, err
			} else {
				// Check if creation time of message is over 14 days //
				if !checkTime(msTime, time.Now()) {
					// Current message was over 14 day old //
					// - return current list of message to delete
					return toDelete, nil
				}
			}

			// Add current message to list of messages to delete //
			toDelete = append(toDelete, m.ID)
		}

		// Check if the list is zero //
		// - return error if no items were added to list
		if len(toDelete) == 0 {
			return nil, errors.New("there was no messages to delete")
		}

		// Return list of messages to delete //
		return toDelete, nil
	}
}

// Find user messages to delete //
func getUserMessages(session *discordgo.Session, message *discordgo.Message, uName string, i int) ([]string, error) {
	// Create list of message to delete //
	toDelete := make([]string, 0)

	// Save current message id //
	current := message.ID

	// Loop while toDelete list is less then the number of messages to be deleted //
	for len(toDelete) < i {
		// Get the given number of messages from the channel //
		// - returns an error if err is not nil
		if messages, err := session.ChannelMessages(message.ChannelID, i, current, "", ""); err != nil {
			return nil, err
		} else {
			// Returns an error if no messages to in the channel //
			if len(messages) == 0 && len(toDelete) == 0 {
				return nil, errors.New("there was no messages to delete with in given channel")
			}

			// Move current message pointer to the last message ID with in messages //
			current = messages[len(messages)-1].ID

			// for each message in messages add message to list if uName maches user ID //
			for _, m := range messages {
				// Get message creation time //
				// - returns an error if err is not nil
				if msTime, err := m.Timestamp.Parse(); err != nil {
					return nil, err
				} else {
					// Check if creation time is older then 14 days //
					// - return current list of message if older then 14 days
					if !checkTime(msTime, time.Now()) {
						return toDelete, nil
					}
				}
				// Gets the user info of the message //
				// - returns an error if err is nil
				if user, err := getMember(session, m); err != nil {
					return nil, err
				} else if user.Nick == uName {
					// user nickname matched uName, add message to list //
					toDelete = append(toDelete, m.ID)
				} else if m.Author.Username == uName {
					// User name matched uName, add message to list //
					toDelete = append(toDelete, m.ID)
				}
			}
		}
	}

	// Add the initial !message -delete call to the toDelete list //
	toDelete = append(toDelete, message.ID)

	// Return list of messages to delete //
	return toDelete, nil
}

// Clear all messages newer then 14 days from a channel //
func clearMessages(session *discordgo.Session, message *discordgo.Message, _ []string) error {
	// Check if user is an admin //
	// - return an error if err is not nil
	// - return if user is not an admin
	if admin, err := isAdmin(session, message); err != nil {
		return err
	} else if !admin {
		_, err := session.ChannelMessageSend(message.ChannelID, "You do not have the permissions to use this command.")
		return err
	}

	// Clear message Loop //
	// - runs toDelete to get the list of messages to delete in batches of 99
	// - loop ends if the list of messages returned from toDelete is 0 or it returns an error
	for ms, err := toDelete(session, message, "", 99, true); len(ms) != 0 || err != nil; ms, err = toDelete(session, message, "", 99, true) {
		// Deletes the current batch of 99 messages //
		// - returns an error if err is not nil
		if err := session.ChannelMessagesBulkDelete(message.ChannelID, ms); err != nil {
			return err
		}
	}

	// Clear was successful //
	return nil
}
