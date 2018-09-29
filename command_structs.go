package CommanD_Bot

type MessageCommands struct {
	commands map[CommandKey]func(RootCommand)error
}

func (m *MessageCommands) RunCommand(command RootCommand) error {
	return m.commands[command.keys[1]](command)
}

type ChannelCommands struct {
	commands map[CommandKey]func(RootCommand)error
}

func (m *ChannelCommands) RunCommand(command RootCommand) error {
	return m.commands[command.keys[1]](command)
}

type GuildCommands struct {
	commands map[CommandKey]func(RootCommand)error
}

func (m *GuildCommands) RunCommand(command RootCommand) error {
	return m.commands[command.keys[1]](command)
}

type PlayerCommands struct {
	commands map[CommandKey]func(RootCommand)error
}

func (m *PlayerCommands) RunCommand(command RootCommand) error {
	return m.commands[command.keys[1]](command)
}

type UtilityCommands struct {
	commands map[CommandKey]func(RootCommand)error
}

func (m *UtilityCommands) RunCommand(command RootCommand) error {
	return m.commands[command.keys[1]](command)
}