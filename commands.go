package CommanD_Bot

type BotCommands struct {
	commands map[string]CommandFunction
}

func (b *BotCommands) GetCommand(cmd string) CommandFunction {
	return b.commands[cmd]
}

