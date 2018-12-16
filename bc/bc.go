package bc

import "github.com/Tskana/CommanD-Bot/core"

type BotCommand struct {
	*core.Command

	BotOptions map[string]core.HandlerFunction
}

func (b *BotCommand) Init(command *core.Command) core.Commander {
	b.Command = command

	b.BotOptions = make(map[string]core.HandlerFunction)

	b.BotOptions["-changecommand"] = b.ChangeCommandHandler
	b.BotOptions["-cc"] = b.ChangeCommandHandler

	return b
}

func (b *BotCommand) Run() error {
	fnc, ok := b.BotOptions[b.Option]
	if !ok {
		return core.NewError("BotCommand Run()", "unknown BotCommand option given")
	}

	err := fnc()
	if err != nil {
		return err
	}

	return b.DeleteMessages(b.ID)
}




