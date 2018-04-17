package CommanD_Bot

import (
	"fmt"
	"log"
	"time"
)

// Error info struct //
type Error struct {
	id   string
	file string
	time time.Time
}

// Error function which interacts with error interface //
func (e *Error) Error() string {
	return fmt.Sprintf("%s, Error: %s in %s", e.time.Format(time.UnixDate), e.id, e.file)
}

// Created a new error //
func NewError(id string, file string) *Error {
	return &Error{id, file, time.Now()}
}

// logs given error //
func PrintError(error error) {
	log.Println(error)
}
