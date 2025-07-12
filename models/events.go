package models

import (
	"time" 
)

type Event struct {
	ID          int
	Name        string `Binding:"required"`
	Description string `Binding:"required"`
	Location    string `Binding:"required"`
	DateTime    time.Time `Binding:"required"`
	UserID      int
}

var events = []Event {} // for now we will store it in slice of name event and later in database

func (e Event) Save(){
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}