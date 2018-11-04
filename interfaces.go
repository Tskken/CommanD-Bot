package CommanD_Bot

type CommandAction interface {
	RunCommand(*Root)error
}
