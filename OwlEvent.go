package OwlEvent

import (
	"fmt"
	"time"
)

// SleepTimer ...
const SleepTimer = 10

// EventHandle ...
type EventHandle struct {
	events  []Event
	Looping bool
}

// OnEvent ...
func (eventHandle *EventHandle) OnEvent(caller string, a func() error, b ...) {
	newEvent := &Event{
		id:       caller,
		isAsync:  async,
		function: a,
	}
	switch(b) {
	case async:
		newEvent.isAsync = true
	}
	eventHandle.events = append(eventHandle.events, *newEvent)
}

// Event ...
type Event struct {
	id           string
	isAsync      bool
	errorMessage string
	function     func() error
}

// Call ...
func (event *Event) Call(parent *EventHandle) error {
	e := event.function()
	return e
}

// AsyncCall ...
func (event *Event) AsyncCall(ch chan error) {
	ch <- event.function()
	close(ch)
}

// Start ...
func (eventHandle *EventHandle) Start() {
	for eventHandle.Looping {
		for _, e := range eventHandle.events {
			switch e.isAsync {
			case true:
				ch := make(chan error)
				go e.AsyncCall(ch)
				if ch != nil {
					fmt.Printf("%+v", ch)
				}
			case false:
				e.Call(eventHandle)
			}
		}
		time.Sleep(time.Second / SleepTimer)
	}
}
