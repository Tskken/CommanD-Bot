package mc

import "github.com/Tskana/CommanD-Bot/core"

type MessageCommand struct {
	*core.Command

	// TODO: Implement command functions
	DeleteMessage core.HandlerFunction
}

func (m *MessageCommand) Init(command *core.Command) core.Commander {
	m.Command = command

	// TODO: Implement initializer
	m.DeleteMessage = m.DeleteMessageHandler
	return m
}

func (m *MessageCommand) Run() error {
	switch m.Option {
	case "-delete", "-del":
		return m.DeleteMessage()
	default:
		return nil
	}
}



