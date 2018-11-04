package CommanD_Bot

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

// Server info structure //
type server struct {
	WordFilter       map[string]bool      // server word filter map
	MuteList         map[string]time.Time // List of muted players
	BanTime          uint                 // server ban time length
	wordFilterThresh float64              // server world filter threshold
	spamFilterThresh float64              // server spam filter threshold
}

// List of all servers //
var serverList = make(map[string]*server)

var muteTimerList = make(map[string]*time.Timer)

// Creates new server //
// - returns a reference to a server and an error (nil if non)
func NewServer() (*server, error) {
	// Creates server structure //
	s := server{}

	// Set server default word filter //
	if err := s.CreateWordFilter(); err != nil {
		return nil, err
	}

	// TODO - Comment
	s.MuteList = make(map[string]time.Time)

	// Set default ban time //
	s.BanTime = 30

	// Set default word filter threshold //
	s.wordFilterThresh = 0.25

	// Set default spam filter threshold //
	s.spamFilterThresh = 0.75

	// Return a reference to server structure //
	return &s, nil
}

// Create word filter map //
// - returns an error (nil if non)
func (s *server) CreateWordFilter() error {
	// Create word filter map //
	s.WordFilter = make(map[string]bool)

	// Load default word filter from file //
	// - returns an error if err is not nil
	if err := s.LoadWordsFromFile(); err != nil {
		return err
	}
	return nil
}

// Change a value with in word filter //
// - returns an error (nil if non)
func (s *server) EditWordFilter(word string, action bool) error {
	if !action {
		// Remove word from filter //
		// - returns an error if word does not exist in list
		if _, ok := s.WordFilter[word]; !ok {
			return errors.New("can not delete a from the filter if the word does not exist")
		} else {
			// Remove wor from filter list //
			delete(s.WordFilter, word)
			return nil
		}
	} else {
		// Add word to filter //
		s.WordFilter[word] = action
		return nil
	}
}

// TODO - Comment
func (s *server) Mute(uId string, duration time.Duration) error {
	if t, ok := s.MuteList[uId]; ok {
		s.MuteList[uId] = t.Add(duration)
		timer, ok := muteTimerList[uId]
		if !ok {
			return errors.New("mute timer did not exist")
		}
		timer.Stop()
		timer.Reset(time.Until(t))
		return nil
	}

	s.MuteList[uId] = time.Now().Add(duration)
	timer := time.AfterFunc(duration, func() {
		err := s.UnMute(uId)
		if err != nil {
			log.Println(err)
		}
	})

	muteTimerList[uId] = timer

	return nil
}

// TODO - Comment
func (s *server) UnMute(uId string) error {
	if _, ok := s.MuteList[uId]; !ok {
		return errors.New("mute id no longer exists")
	} else {
		delete(s.MuteList, uId)
	}

	if _, ok := muteTimerList[uId]; !ok {
		return errors.New("mute timer no longer exists")
	} else {
		timer := muteTimerList[uId]
		timer.Stop()
		delete(muteTimerList, uId)
	}
	return nil
}

// TODO - Comment
func (s *server) IsMuted(uId string) (time.Time, bool) {
	t, ok := s.MuteList[uId]
	return t, ok
}

// Get admin role //
// - returns reference to ID and an error (nil if non)
func GetAdminRole(guild *discordgo.Guild) (*string, error) {
	// Look though guild roles //
	for _, role := range guild.Roles {
		// Check if role name is Admin
		if role.Name == "Admin" {
			// Return admin role id //
			return &role.ID, nil
		}
	}
	// Role was not found //
	// - return an error
	return nil, errors.New("role did not exist in guild")
}

// Check if the Admin role exist in guild //
// - return an error (nil if non)
func RoleCheck(session *discordgo.Session, guild *discordgo.Guild) error {
	// Try go get admin role //
	// - if error is not nil create admin role
	_, err := GetAdminRole(guild)
	if err != nil {
		// Create a new role with in the guild //
		// - returns an error if err is not nil
		if role, err := session.GuildRoleCreate(guild.ID); err != nil {
			return err
		} else {
			// Set the new roles name to Admin and permissions to admin //
			// - returns an error if err is not nil
			_, err = session.GuildRoleEdit(guild.ID, role.ID, "Admin", 16724736, true, 8, true)
			return err
		}
	}
	return nil
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
