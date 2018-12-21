package core

import (
	"strings"
)

type ParsedCommand struct {
	Command string
	Option  string
	Args    []string
}

func ParseMessage(message string) *ParsedCommand {
	inputArgs := strings.Fields(strings.ToLower(message))

	pc := new(ParsedCommand)

	if len(inputArgs) >= 2 {
		pc.Option = inputArgs[1]
	}

	if len(inputArgs) > 2 {
		pc.Args = inputArgs[2:]
	}

	pc.Command = inputArgs[0]

	return pc
}
