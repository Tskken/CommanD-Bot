package CommanD_Bot

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
	"time"
)

type Root struct {
	*discordgo.Session
	*discordgo.Message
}

func (r *Root) Run() error {
	return bot.GetCommand(r.CommandRoot()).RunCommand(r)
}

func (r *Root) Args() []string {
	return ToLower(strings.Fields(r.Content))
}

func (r *Root) CommandArgs() []string{
	// Parce input in to arguments and set them to lower case //
	return r.Args()[2:]
}

func (r *Root) CommandRoot()string {
	return r.Args()[0]
}

func (r *Root) CommandType()string {
	return r.Args()[1]
}

func (r *Root) MessageSend(content string) error {
	_, err := r.ChannelMessageSend(r.ChannelID, content)
	return err
}

// Create channel with given name an type //
// - Returns an error (nil if non)
func (r *Root) NewChannel(name, cType string) error {
	// Get guild to add channel to //
	// - Returns an error if err is not nil
	if guild, err := r.GetGuild(); err != nil {
		return err
	} else {
		// Check to make sure the type was correct //
		// - returns and error if it was not
		if cType != "text" && cType != "voice" {
			return errors.New("channel type was not ether text or voice")
		}

		// Create channel with given name and type in guild //
		// - Returns an error if err is not nil
		if _, err := r.GuildChannelCreate(guild.ID, name, cType); err != nil {
			return err
		}
	}

	return nil
}

// Get the channel to delete //
// - Returns a reference to a channel and an error (nil if non)
func (r *Root) GetChannelToDel(name string) (*discordgo.Channel, error) {
	// Get guild channel the channel exists in //
	// - returns an error if err is not nil
	if guild, err := r.GetGuild(); err != nil {
		return nil, err
	} else {
		// Gets all channels with in the guild //
		// - Returns an error if err is not nil
		if chs, err := r.GuildChannels(guild.ID); err != nil {
			return nil, err
		} else {
			// Check list of channels for given name to delete //
			for _, c := range chs {
				// Return channel if channel name = given name //
				if c.Name == name {
					return c, nil
				}
			}
		}
	}

	// Return error if channel was not found in guild //
	return nil, errors.New("could not find channel to delete")
}

// Scan a message and to classify it //
// - Returns an error if err is not nil
func (r *Root) Scan() error {
	msg := r.Args()

	g, err := r.GetGuild()
	if err != nil {
		return err
	}
	server := serverList[g.ID]

	// Check for bad words //
	// Count of number of bad words with in sentence //
	bWordCount := 0
	// Check each word with in KeyWordMap //
	// - Increment BWordCount if word value is true
	for _, ms := range msg {
		if _, ok := server.WordFilter[ms]; ok {
			bWordCount++
		}
	}

	//log.Println(float64(bWordCount) / float64(len(msg)))

	// Check average occurrence of bad words with in the message //
	// - Remove the message if average is over threshold
	if avr := float64(bWordCount) / float64(len(msg)); avr >= serverList[g.ID].wordFilterThresh {
		// Get the message to delete //
		// - Returns an error if err is not nil
		if a, err := r.GetMessages("", 0, true); err != nil {
			return err
		} else {
			// Delete message //
			// - Return an error if err is not nil
			if err := r.ChannelMessageDelete(r.ChannelID, a[0]); err != nil {
				return err
			}

			// Notify user of there mistake //
			// - Returns an error if err is not nil
			if _, err := r.ChannelMessageSend(r.ChannelID, r.Author.Mention()+" don't say that... that's not nice... bad! :point_up:"); err != nil {
				return err
			}

			return nil
		}
	} else {
		// Find the score of the message with in the Good and Spam categories //
		// - Returns an error if err is not nil
		if scores, inx, strict, err := filterClassifier.SafeProbScores(msg); err != nil {
			return err
		} else {
			//log.Println(scores)

			// If score of both classes are the same check with threshold //
			if !strict {
				// Goes through scores of each class //
				for i, score := range scores {
					switch i {
					// Skip good class value //
					case 0:
						break
						// Check Spam class value //
					case 1:
						// Check if value is over threshold //
						// - Delete if true
						if score >= serverList[g.ID].spamFilterThresh {
							// Get message to delete //
							// - Returns an error if err is not nil
							if a, err := r.GetMessages("", 0, true); err != nil {
								return err
							} else {
								// Delete message //
								// - Returns an error if err is not nil
								if err := r.ChannelMessageDelete(r.ChannelID, a[0]); err != nil {
									return err
								}
								// Notify member of there mistake //
								// - Returns an error if err is not nil
								if _, err := r.ChannelMessageSend(r.ChannelID, r.Author.Mention()+" Why must you spam... No spamming... bad! :point_up:"); err != nil {
									return err
								}
							}
						}
					}
				}
			} else {
				// One of the values was larger then the other //
				// - Good = 0
				// - Spam = 1
				switch inx {
				// Return if Good is the largest value //
				case 0:
					return nil
					// Delete messages if Spam is larger //
				case 1:
					// Get message to delete //
					// - Return an error if err is not nil
					if a, err := r.GetMessages( "", 0, true); err != nil {
						return err
					} else {
						// Delete message //
						// - Return an error ir err is not nil
						if err := r.ChannelMessageDelete(r.ChannelID, a[0]); err != nil {
							return err
						}
						// Notify member of there mistake //
						// - Returns an error if err is not nil
						if _, err := r.ChannelMessageSend(r.ChannelID, r.Author.Mention()+" Why must you spam... No spamming... bad! :point_up:"); err != nil {
							return err
						}
					}
					// Returns an error if inx was anything but 0 or 1 //
				default:
					return errors.New("inx was a value other then 0,1, or 2")
				}
			}
		}
	}

	return nil
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
		if len(mId.([]string)) > 99 {
			iterator := len(mId.([]string)) / 99
			iterator += 1
			for i := 0; i <= iterator; i++ {
				if i == iterator {
					newList := mId.([]string)[:len(mId.([]string))]
					return r.ChannelMessagesBulkDelete(r.ChannelID, newList)
				}
				newList := mId.([]string)[:100]
				mId.([]string) = mId.([]string)[100:]
				err := r.ChannelMessagesBulkDelete(r.ChannelID, newList)
				if err != nil {
					return err
				}
			}
		}
		return r.ChannelMessagesBulkDelete(r.ChannelID, mId.([]string))
	case string:
		return r.ChannelMessageDelete(r.ChannelID, mId.(string))
	default:
		return errors.New("mId is not of type []string or string")
	}

}

// TODO - Comment
func (r *Root) IsMuted() (bool, error) {
	guild, err := r.GetGuild()
	if err != nil {
		return false, err
	}
	server := serverList[guild.ID]

	member, err := r.GetMember()
	if err != nil {
		return false, err
	}

	if muteTime, muted := server.IsMuted(member.User.ID); muted {
		log.Println("is muted till " + time.Until(muteTime).Truncate(time.Second).String())
		if err := r.DeleteMessage(r.ID); err != nil {
			return false, err
		}
		return true, r.MessageSend(member.User.Mention()+" you are muted for "+time.Until(muteTime).Truncate(time.Second).String())
	}

	return false, nil
}

// Check if a user is an admin //
// - returns admin state boolean and an error (nil if non)
func (r *Root) IsAdmin() (bool, error) {
	// Get guild the message was sent in //
	// - returns an error if err is not nil
	guild, err := r.GetGuild()
	if err != nil {
		return false, err
	}

	// Get member that created the message //
	// - returns an error if err is not nil
	member, err := r.GetMember()
	if err != nil {
		return false, err
	}

	// Gets the admin role ID from the guild //
	// - returns an error if err is not nil
	if roleID, err := GetAdminRole(guild); err != nil {
		return false, err
	} else {
		// Check member roles //
		for _, memRole := range member.Roles {
			// check role //
			if memRole == *roleID {
				// User had admin role //
				// - return admin is true
				return true, err
			}
		}
		// User did not have admin role //
		// - return admin is false
		return false, err
	}
}

// Gets guild structure //
// - returns an error (nil if non)
func (r *Root) GetGuild() (*discordgo.Guild, error) {
	// Get the channel the message was created //
	// - returns an error if err is not nil
	if c, err := r.State.Channel(r.ChannelID); err != nil {
		return nil, err
	} else {
		// Gets guild from channel guild ID //
		// - returns a reference to guild structure and an error (nil if non)
		return r.State.Guild(c.GuildID)
	}
}

// Gets member structure //
// - returns an error (nil if non)
func (r *Root) GetMember() (*discordgo.Member, error) {
	// Gets the guild the message was created in //
	// - returns an error if err is not nil
	if g, err := r.GetGuild(); err != nil {
		return nil, err
	} else {
		// Get member from guild with message author ID //
		// - returns a reference to member structure and an error (nil if non)
		return r.GuildMember(g.ID, r.Author.ID)
	}
}
