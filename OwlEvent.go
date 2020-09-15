package owlevent

import (
	"fmt"
)

// EventHandle ...
type EventHandle struct {
	events []Event
}

// Event ...
type Event struct {
	id           string
	errorMessage error
	function     func()
}

// OnEvent ...
func (eh *EventHandle) OnEvent(caller string, a func()) {
	newEvent := &Event{
		id:       caller,
		function: a,
	}
	eh.events = append(eh.events, *newEvent)
}

// RemoveEventByIndex ...
// slice -> Delete without preserving order
func (eh *EventHandle) RemoveEventByIndex(i int) {
	eh.events = append(eh.events[:i], eh.events[i+1:]...)
}

// Call ...
func (eh *EventHandle) Call(ch chan bool, eventID string) {
	if len(eh.events) > 0 {
		for i := 0; i < len(eh.events); i++ {
			fmt.Println("Length of eventHandle.events: ", len(eh.events))
			value := eh.events[i]
			if value.id == eventID {
				value.function()
				// slice -> non-preserved order...
				eh.RemoveEventByIndex(i)
				// requires index decrementation
				i--
			} else {
				continue
			}
		}
		ch <- true
	}
	close(ch)
}
