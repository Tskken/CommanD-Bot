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
	return fmt.Sprintf("%v, Error: %s in %s", e.time, e.id, e.file)
}

func NewError(id string, file string)*Error{
	return &Error{id, file, time.Now()}
}

func PrintError(error error){
	log.Println(error)
}