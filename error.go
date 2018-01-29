package CommanD_Bot

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

func NewError(id string, file string)(e *Error){
	e.id = id
	e.time = time.Now()
	e.file = file
	return
}

func PrintError(error error){
	log.Println(error)
}