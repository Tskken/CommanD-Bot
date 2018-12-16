package mc

import (
	"github.com/Tskana/CommanD-Bot/core"
	"strconv"
)

// TODO: Posibly could use some optimizations
func (m *MessageCommand) DeleteMessageHandler() error {
	member, err := m.GetMember()
	if err != nil {
		return err
	}

	perm := core.Permissions{}
	ok := false

	for _, r := range member.Roles {
		g, err := m.GetGuild()
		if err != nil {
			return err
		}
		nRole := ""
		for _, role := range g.Roles {
			if role.ID == r {
				nRole = role.Name
			}
		}
		perm, ok = core.BotPermissions[nRole]
		if ok {
			break
		}
	}

	if !ok {
		return core.NewError("DeleteMessageHandler()", "no role given to user")
	}

	switch {
	case len(m.Args) == 0:
		for _, p := range perm.Permissions {
			switch p {
			case "*", DELALL:
				ms, err := m.GetMessage()
				if err != nil {
					return err
				}

				return m.DeleteMessages(ms...)
			case DELSELF:
				ms, err := m.GetMessage(m.Author.Mention())
				if err != nil {
					return err
				}

				return m.DeleteMessages(ms...)
			default:
				return core.NewError("DeleteMessageHandler()", "no supported permissions given to role")
			}
		}
	case len(m.Args) == 1:
		for _, p := range perm.Permissions {
			switch p {
			case "*", DELALL:
				n, err := strconv.Atoi(m.Args[0])
				if err != nil {
					ms, err := m.GetMessage(m.Args[0])
					if err != nil {
						return err
					}

					return m.DeleteMessages(ms...)
				}

				ms, err := m.GetNMessages(n)
				if err != nil {
					return err
				}

				return m.DeleteMessages(ms...)
			case DELSELF:
				n, err := strconv.Atoi(m.Args[0])
				if err != nil {
					return core.NewError("DeleteMessageHandler()", "insignificant permissions to use delete message for user")
				}

				ms, err := m.GetNMessages(n, m.Author.Mention())
				if err != nil {
					return err
				}

				return m.DeleteMessages(ms...)
			default:
				return core.NewError("DeleteMessageHandler()", "no supported permissions given to role")
			}
		}
	case len(m.Args) >= 2:
		for _, p := range perm.Permissions {
			switch p {
			case "*", DELALL:
				n, err := strconv.Atoi(m.Args[0])
				if err != nil {
					ms, err := m.GetMessage(m.Args...)
					if err != nil {
						return err
					}

					return m.DeleteMessages(ms...)
				}

				ms, err := m.GetNMessages(n, m.Args[1:]...)
				if err != nil {
					return err
				}

				return m.DeleteMessages(ms...)
			case DELSELF:
				return core.NewError("DeleteMessageHandler()", "insignificant permissions to use delete message for user")
			default:
				return core.NewError("DeleteMessageHandler()", "no supported permissions given to role")
			}
		}
	default:
		return core.NewError("DeleteMessageHandler()", "to many arguments given to handler")
	}
	return core.NewError("DeleteMessageHandler()", "something went wrong")
}
