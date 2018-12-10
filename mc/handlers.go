package mc

import (
	"github.com/Tskana/CommanD-Bot/core"
	"strconv"
)

// TODO: update to handler more then one given user
func (m *MessageCommand) DeleteMessageHandler() error {
	switch len(m.Args) {
	case 0:
		ms, err := m.GetMessage()
		if err != nil {
			return err
		}

		return m.DeleteMessages(ms)
	case 1:
		n, err := strconv.Atoi(m.Args[0])
		if err != nil {
			ms, err := m.GetMessage(m.Args[0])
			if err != nil {
				return err
			}

			return m.DeleteMessages(ms)
		}

		ms, err := m.GetNMessages(n)
		if err != nil {
			return err
		}

		return m.DeleteMessages(ms...)
	case 2:
		n, err := strconv.Atoi(m.Args[0])
		if err != nil {
			return err
		}

		ms, err := m.GetNMessages(n, m.Args[1])
		if err != nil {
			return err
		}

		return m.DeleteMessages(ms...)
	default:
		return core.NewError("DeleteMessageHandler()", "to many arguments given to handler")
	}
}
