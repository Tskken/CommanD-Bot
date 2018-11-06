package CommanD_Bot

type CommandFunction interface {
	RunCommand(*Root)error
}
