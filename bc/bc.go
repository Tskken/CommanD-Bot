package bc

import (
	"github.com/Tskana/CommanD-Bot/core"
)

type BotCommand struct {
	*core.Command

	BotOptions map[string]core.HandlerFunction
}

func (b *BotCommand) Init(command *core.Command) map[string]core.HandlerFunction {
	b.Command = command

	b.BotOptions = make(map[string]core.HandlerFunction)

	b.BotOptions["-changecommand"] = b.ChangeCommandHandler
	b.BotOptions["-cc"] = b.ChangeCommandHandler

	return b.BotOptions
}
