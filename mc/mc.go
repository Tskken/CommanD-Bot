package mc

import "github.com/Tskana/CommanD-Bot/core"

type MessageCommand struct {
	*core.Command

	MessageOptions map[string]core.HandlerFunction
}

func (m *MessageCommand) Init(command *core.Command) core.Commander {
	m.Command = command

	m.MessageOptions = make(map[string]core.HandlerFunction)

	m.MessageOptions["-delete"] = m.DeleteMessageHandler
	m.MessageOptions["-del"] = m.DeleteMessageHandler
	return m
}

func (m *MessageCommand) Run() error {
	fnc, ok := m.MessageOptions[m.Option]
	if !ok {
		return core.NewError("MessageCommand Run()", "unknown MessageCommand option given")
	}
	err := fnc()
	if err != nil {
		return err
	}

	return m.DeleteMessages(m.ID)
}
