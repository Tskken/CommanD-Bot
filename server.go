package CommanD_Bot

import (
	"errors"
	"github.com/bwmarrin/discordgo"
)

type server struct {
	WordFilter       map[string]bool
	BanTime          uint
	wordFilterThresh float64
	spamFilterThresh float64
}

var serverList = make(map[string]*server)

func newServer() (*server, error) {
	s := server{}
	err := s.createWordFilter()
	if err != nil {
		return nil, err
	}
	s.setBanTimer(30)
	s.setWordFilterThreshold(0.25)
	s.setSpamFilterThreshold(0.75)

	return &s, nil
}

func (s *server) createWordFilter() error {
	s.WordFilter = make(map[string]bool)
	err := s.loadWordsFromFile()
	if err != nil {
		return err
	}
	return nil
}

func (s *server) editWordFilter(word string, val bool) error {
	if !val {
		if _, ok := s.WordFilter[word]; !ok {
			return errors.New("can not delete a from the filter if the word does not exist")
		} else {
			delete(s.WordFilter, word)
			if _, ok := s.WordFilter[word]; ok {
				return errors.New("word was not deleted")
			}
		}
	} else {
		s.WordFilter[word] = val
		if _, ok := s.WordFilter[word]; !ok {
			return errors.New("was not able to add word to map")
		}
	}
	return nil
}

func (s *server) setBanTimer(time uint) {
	s.BanTime = time
}

func (s *server) setWordFilterThreshold(thresh float64) {
	s.wordFilterThresh = thresh
}

func (s *server) setSpamFilterThreshold(thresh float64) {
	s.spamFilterThresh = thresh
}

func getAdminRole(guild *discordgo.Guild) (*string, error) {
	// Look though guild roles //
	for _, role := range guild.Roles {
		// Admin role exist and return the role ID //
		if role.Name == "Admin" {
			return &role.ID, nil
		}
	}
	return nil, errors.New("role did not exist in guild")
}

// Check if the Admin role is with in the guild and create it if not //
func roleCheck(session *discordgo.Session, guild *discordgo.Guild) error {
	_, err := getAdminRole(guild)
	if err != nil {
		// Create a new role with in the guild //
		// Returns an error if err is not nil
		if role, err := session.GuildRoleCreate(guild.ID); err != nil {
			return err
		} else {
			// Set the new roles name to Admin, permissions to admin //
			// Return the new roles ID
			// Returns an error if err is not nil
			_, err = session.GuildRoleEdit(guild.ID, role.ID, "Admin", 16724736, true, 8, true)
			return err
		}
	}
	return nil
}

// Check if a user is an admin //
func isAdmin(session *discordgo.Session, message *discordgo.Message) (bool, error) {
	// Get the guild the message was sent in //
	// Logs an error if err is not nil
	guild, err := getGuild(session, message)
	if err != nil {
		return false, err
	}

	// Gets the member that created the message from the guild //
	// Logs an error if err is not nil
	member, err := getMember(session, message)
	if err != nil {
		return false, err
	}

	// Gets the admin role ID from the guild //
	// Creates the role and returns the new ID if it does not exist
	// Logs an error if err is not nil
	if roleID, err := getAdminRole(guild); err != nil {
		return false, err
	} else {
		// Check members roles //
		for _, memRole := range member.Roles {
			// Member has admin role give them admin permissions //
			if memRole == *roleID {
				return true, err
			}
		}
		return false, err
	}
}

// Gets guild info for the given guild //
func getGuild(session *discordgo.Session, message *discordgo.Message) (*discordgo.Guild, error) {
	// Get the channel the messages is from with in the guild //
	// Returns an error if err is not nil
	if c, err := session.State.Channel(message.ChannelID); err != nil {
		return nil, err
	} else {
		// Returns the guild info //
		// Returns an error (nil if non | not nil if error) //
		return session.State.Guild(c.GuildID)
	}
}

// Gets a member from with in a guild //
func getMember(session *discordgo.Session, message *discordgo.Message) (*discordgo.Member, error) {
	// Gets the guild that the member should be in //
	// Returns an error if err is not nil
	if g, err := getGuild(session, message); err != nil {
		return nil, err
	} else {
		// Returns a member //
		// Returns an error (nil if non | not nil if error)
		return session.GuildMember(g.ID, message.Author.ID)
	}
}
