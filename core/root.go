package core

import (
	"github.com/bwmarrin/discordgo"
	"time"
)

type Root struct {
	*discordgo.Session
	*discordgo.Message
}

func (r *Root) GetMessage(uID ...string) ([]string, error) {
	if uID != nil {
		return r.getUserMessage(uID)
	}
	ms, err := r.getMessage()
	return []string{ms}, err
}

func (r *Root) getUserMessage(uID []string) (mID []string, err error) {
	for _, id := range uID {
		ms, err := r.ChannelMessages(r.ChannelID, 1, r.ID, "", "")
		for true {
			if err != nil {
				return nil, err
			}

			if len(ms) == 0 {
				return nil, NewError("getUserMessage()", "returned messages is zero")
			}

			if IsMentioned(ms[0].Author.Mention(), id) {
				mID = append(mID, ms[0].ID)
				break
			}

			ms, err = r.ChannelMessages(r.ChannelID, 1, ms[0].ID, "", "")
		}
	}
	return
}

func (r *Root) getMessage() (string, error) {
	ms, err := r.ChannelMessages(r.ChannelID, 1, r.ID, "", "")
	if err != nil {
		return "", err
	}

	return ms[0].ID, nil
}

func (r *Root) GetNMessages(n int, uID ...string) ([]string, error) {
	if uID != nil {
		return r.getNUserMessages(n, uID)
	}

	return r.getNMessages(n)
}

func (r *Root) getNUserMessages(n int, uID []string) (mID []string, err error) {
	current := r.ID

	for true {
		ms, err := r.ChannelMessages(r.ChannelID, n, current, "", "")
		if err != nil {
			return nil, err
		}

		for _, m := range ms {
			for _, id := range uID {
				if IsMentioned(m.Author.Mention(), id) {
					mID = append(mID, m.ID)
				}

				if len(mID) == n {
					return mID, nil
				}
			}
		}

		if len(ms) < n {
			return mID, nil
		}
	}

	return nil, NewError("getNUserMessages()", "something went wrong in getNUserMessages()")
}

func (r *Root) getNMessages(n int) (mID []string, err error) {
	ms, err := r.ChannelMessages(r.ChannelID, n, r.ID, "", "")
	if err != nil {
		return nil, err
	}

	for _, m := range ms {
		mID = append(mID, m.ID)
	}

	return
}

func (r *Root) DeleteMessages(mID ...string) error {
	mID, err := r.checkMessageTime(mID)
	if err != nil {
		return err
	}

	if len(mID) > 99 {
		for len(mID) > 0 {
			err := r.ChannelMessagesBulkDelete(r.ChannelID, mID[:100])
			if err != nil {
				return err
			}

			mID = mID[100:]
		}
		return nil
	}

	return r.ChannelMessagesBulkDelete(r.ChannelID, mID)
}

func (r *Root) checkMessageTime(mID []string) ([]string, error) {
	for i, id := range mID {
		ms, err := r.ChannelMessage(r.ChannelID, id)
		if err != nil {
			return nil, err
		}
		t, err := ms.Timestamp.Parse()
		if err != nil {
			return nil, err
		}
		if !CheckTime(t, time.Now()) {
			return mID[:i], nil
		}
	}
	return mID, nil
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
