package botErrors

import (
	"time"
	"fmt"
	"log"
)

type Error struct {
	id string
	file string
	time time.Time
}

func (e *Error) Error()string{
	return fmt.Sprintf("%s, Error: %s in %s", e.time.Format(time.UnixDate), e.id, e.file)
}

func NewError(id string, file string)*Error{
	return &Error{id, file, time.Now()}
}

func PrintError(error error){
	log.Println(error)
}