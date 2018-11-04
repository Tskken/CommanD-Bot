package CommanD_Bot

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"strconv"
	"time"
	"strings"
)


type MessageCommands struct {
	commands map[string]func(*Root)error
}

func (m *MessageCommands) RunCommand(root *Root) error {
	return m.commands[root.CommandType()](root)
}

// TODO - Update Comments

// Set message command structure //
func LoadMessageCommand() *MessageCommands {
	// Create message command structure //
	m := MessageCommands{}

	// Create sub command map //
	m.commands = make(map[string]func(*Root) error)

	// Set message delete function //
	m.commands["-delete"] = DeleteMessageHandler
	m.commands["-del"] = DeleteMessageHandler

	// Set message clear function //
	m.commands["-clear"] = ClearMessages
	m.commands["-cl"] = ClearMessages

	// Return reference to command structure //
	return &m
}

// Create CommandInfo structure //
func LoadMessageCommandInfo() *commandInfo {
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
		"    **<User Name> <Number between 1 - 99>:** Deletes the last given number of messages by the given user (only can be used by admin)."

	// Set info for clear message function //
	m.commands["-clear"] = "**-clear** or **-cl**.\n" +
		"**Info**: Deletes all message with in a channel newer then 14 days."

	// Returns a reference to command info structure //
	return &m
}

// Message delete function //
// - returns an error (nil if non)
func DeleteMessageHandler(root *Root) error {
	// Gets user admin status //
	// - returns error if err is not nil
	// - true: is admin
	// - false: not admin
	admin, err := root.IsAdmin()
	if err != nil {
		return err
	}

	// Check the length of args for possible passed arguments //
	switch len(root.CommandArgs()) {
	// -del was called with no arguments //
	// Deletes the last message sent in server or by user if user is not an admin
	case 0:
		// Get messages to delete //
		// - returns an error if err is not nil
		if messages, err := root.GetMessages("", 1, admin); err != nil {
			return err
		} else {
			// Delete messages //
			// - returns error (nil if non)
			return root.DeleteMessage(messages)
		}
	// -del was called with either a number of messages or a user name //
	// - Number: deletes that number of messages in that channel or by that user if the user is not an admin
	// - UserName: deletes the last message by that given user name (only can be used by admins)
	case 1:
		// Try's to convert second argument to an integer //
		if i, err := strconv.Atoi(root.CommandArgs()[0]); err != nil {
			if !admin {
				// - returns an error (nil if non)
				return root.MessageSend("You are not an admin, you can not delete message by a specific user.")
			}
			// Argument two was a username //
			// - get last message sent by given user to delete
			// - returns an error if err is not nil
			if messages, err := root.GetMessages(root.CommandArgs()[0], 1, admin); err != nil {
				return err
			} else {
				// Delete given messages //
				// - returns an error (nil if non)
				return root.DeleteMessage(messages)
			}
		} else {
			// Argument two was a number //
			// - checks given number is between 1 and 99
			if i >= 100 || i <= 0 {
				// Send message to channel saying you do not have permission for this action //
				// - returns an error (nil if non)
				return root.MessageSend("You entered a number that I don't understand. Please enter a number between 1-99.")
			}

			// Gets last given number of messages to delete.  Gets messages by user if not admin //
			// - returns an error if err is not nil
			if messages, err := root.GetMessages("", i, admin); err != nil {
				return err
			} else {
				// Delete given messages //
				// - returns an error (nil if non)
				return root.DeleteMessage(messages)
			}
		}
	// -del was called with both a number to delete and a username //
	// - can not be called by non admin users
	default:
		// Check if the user is an admin //
		if !admin {
			// Send message saying user was not an admin //
			// - return an error (nil if non)
			return root.MessageSend("You can not use this root. If you want to multiple of your own messages just type !ms -del <number of messages>.")
		}

		// Try and convert the second argument to an integer //
		// - returns an error is err is not nil
		if i, err := strconv.Atoi(root.CommandArgs()[1]); err != nil {
			return err
		} else {
			// Checks number is with in range of 1 and 99 //
			if i >= 100 || i <= 0 {
				// Send message to channel for value being out of range //
				// - return an error (nil if non)
				return root.MessageSend("You entered a number that I don't understand. Please enter a number between 1-99.")
			}
			// Gets messages by user to be deleted //
			// - returns an error if err is not nil
			if messages, err := root.GetMessages(root.CommandArgs()[0], i, true); err != nil {
				return err
			} else {
				// Deletes messages //
				// - returns an error (nil if non)
				return root.DeleteMessage(messages)
			}
		}
	}
}

// Get messages from channel to delete //
// - returns an array of strings and an error (nil if non)
func (r *Root) GetMessages(uName string, i int, admin bool) ([]string, error) {
	// Check if user is an admin //
	if !admin {
		// User was not admin, get only messages create by them //
		// - returns list of message by user
		// - returns error (nil if non)
		member, err := r.GetMember()
		if err != nil {
			return nil, err
		}
		return r.GetMessagesId(member.User.Mention(), i)
	}

	return r.GetMessagesId(uName, i)
}

// Find user messages to delete //
func (r *Root) GetMessagesId(uName string, i int) ([]string, error) {
	// Create list of message to delete //
	toDelete := make([]string, 0)

	// Save current message id //
	current := r.ID

	// Loop while toDelete list is less then the number of messages to be deleted //
	for len(toDelete) < i {
		// Get the given number of messages from the channel //
		// - returns an error if err is not nil
		if messages, err := r.ChannelMessages(r.ChannelID, i, current, "", ""); err != nil {
			return nil, err
		} else {
			// Returns an error if no messages to in the channel //
			if len(messages) == 0 && len(toDelete) == 0 {
				if err := r.DeleteMessage(r.ID); err != nil {
					return nil, err
				}
				return nil, errors.New("there was no messages to delete with in given channel")
			} else if len(messages) == 0 && len(toDelete) != 0{
				toDelete = append(toDelete, r.ID)
				return toDelete, nil
			}

			// Move current message pointer to the last message ID with in messages //
			current = messages[len(messages)-1].ID

			// for each message in messages add message to list if uName matches user ID //
			for _, m := range messages {
				// Get message creation time //
				// - returns an error if err is not nil
				if ok, err := MessageTime(r.Message); err != nil {
					return nil, err
				} else if !ok {
					toDelete = append(toDelete, r.ID)
					return toDelete, nil
				}

				if uName != "" {
					// Gets the user info of the message //
					// - returns an error if err is nil
					if member, err := r.GetMember(); err != nil {
						return nil, err
					} else if GetMention(member, uName) {
						toDelete = append(toDelete, m.ID)
					}
				} else {
					toDelete = append(toDelete, m.ID)
				}


				if len(toDelete) == i {
					break
				}
			}
		}
	}

	// Add the initial !message -delete call to the toDelete list //
	toDelete = append(toDelete, r.ID)

	// Return list of messages to delete //
	return toDelete, nil
}

func (r *Root) DeleteMessage(mId interface{}) error {
	switch mId.(type) {
	case []string:
		return r.ChannelMessagesBulkDelete(r.ChannelID, mId.([]string))
	case string:
		return r.ChannelMessageDelete(r.ChannelID, mId.(string))
	default:
		return errors.New("mId is not of type []string or string")
	}

}

func MessageTime(message *discordgo.Message) (bool, error) {
	// Get message creation time //
	// - returns an error if err is not nil
	if msTime, err := message.Timestamp.Parse(); err != nil {
		return false, err
	} else {
		return CheckTime(msTime, time.Now()), nil
	}
}

func GetMention(member *discordgo.Member, mention string) bool {
	if mention[2] == '!' {
		mention = strings.Join(strings.Split(mention, "!"), "")
	}

	return member.User.Mention() == mention
}

// Clear all messages newer then 14 days from a channel //
func ClearMessages(root *Root) error {
	// Check if user is an admin //
	// - return an error if err is not nil
	// - return if user is not an admin
	if admin, err := root.IsAdmin(); err != nil {
		return err
	} else if !admin {
		return root.MessageSend("You do not have the permissions to use this command.")
	}

	// Clear message Loop //
	// - runs toDelete to get the list of messages to delete in batches of 99
	// - loop ends if the list of messages returned from toDelete is 0 or it returns an error
	for ms, err := root.GetMessages("", 99, true); len(ms) > 0; ms, err = root.GetMessages("", 99, true) {
		if err != nil {
			return err
		}

		// Deletes the current batch of 99 messages //
		// - returns an error if err is not nil
		if err := root.DeleteMessage(ms); err != nil {
			return err
		}
	}

	// Clear was successful //
	return nil
}

