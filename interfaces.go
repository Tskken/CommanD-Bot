package CommanD_Bot

type CommandKey string

type CommandAction interface {
	RunCommand(RootCommand)error
}
