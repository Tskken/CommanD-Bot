package messagecommands

import "github.com/Tskana/CommanD-Bot/core"

type MessageCommand struct {
	*core.Command

	// TODO: Implement command functions
}

func (m *MessageCommand) Init(command *core.Command) core.Commander {
	m.Command = command

	// TODO: Implement initializer
	return m
}

func (*MessageCommand) Run() error {
	panic("implement me")
}



