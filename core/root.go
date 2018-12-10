package core

import (
	"errors"
	"github.com/bwmarrin/discordgo"
	"time"
)

type Root struct {
	*discordgo.Session
	*discordgo.Message
}

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

		if len(ms) == 0 {
			return "", errors.New("no user user found")
		}

		for _, id := range uID {
			if IsMentioned(ms[0].Author.Mention(), id) {
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

func (r *Root) GetNMessages(n int, uID... string) ([]string, error) {
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

func (r *Root) DeleteMessages(mID... string) error {
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
