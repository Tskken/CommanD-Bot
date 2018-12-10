package core

import (
	"errors"
	"github.com/bwmarrin/discordgo"
)

type Root struct {
	*discordgo.Session
	*discordgo.Message
}

// TODO: optimize function
func (r *Root) GetMessage(uID... string) (string, error) {
	if uID != nil {
		return r.getUserMessage(uID)
	}

	return r.getMessage()
}

func (r *Root) getUserMessage(uID []string) (string, error) {
	ms, err := r.ChannelMessages(r.ChannelID, 1, r.ID, "", "")
	for true {
		if err != nil {
			return "", err
		}

		for _, id := range uID {
			if ms[0].Author.Mention() == id {
				return ms[0].ID, nil
			}
		}

		ms, err = r.ChannelMessages(r.ChannelID, 1, ms[0].ID, "", "")
	}

	return "", errors.New("something went wrong in getUserMessage()")
}

func (r *Root) getMessage() (string, error) {
	ms, err := r.ChannelMessages(r.ChannelID, 1, r.ID, "", "")
	if err != nil {
		return "", err
	}

	return ms[0].ID, nil
}

// TODO: optimize function
func (r *Root) GetNMessages(n int, uID... string) ([]string, error) {
	if n > 99 {
		return nil, errors.New("you can only delete up to 99 message at one time")
	}

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
				if id == m.Author.Mention() {
					mID = append(mID, m.ID)
				}

				if len(mID) == n {
					return
				}
			}
		}

		if len(ms) < n {
			return
		}
	}

	return nil, errors.New("something went wrong in getNUserMessages()")
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

