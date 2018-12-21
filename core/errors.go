package core

type BotErrors struct {
	location string
	info     string
}

func (e *BotErrors) Error() string {
	return e.location + ": " + e.info
}

func NewError(location, info string) *BotErrors {
	return &BotErrors{
		location: location,
		info:     info,
	}
}
