package adv

import (
	"errors"
	"runtime/debug"
)

// EventHandler ...
type EventHandler interface {
	Id(ev *Event) (string, error)
	Callback(ev *Event) error
}

// Event ...
type Event struct {
	ID           string
	ErrorMessage error
}

// ! Handle Event slices!

// RemoveEventByIndex ...
// -> RemoveEventByIndex(indexToRemove, &EventList)
func RemoveEventByIndex(index int, el *[]Event) {
	(*el) = append((*el)[:index], (*el)[index+1:]...)
}

// AppendEvents ...
func AppendEvents(move *[]Event, to *[]Event) {
	(*to) = append((*to), (*move)...)
}

// Push ...
func Push(push *Event, to *[]Event) {
	(*to) = append((*to), *push)
}

// PopAndGetEvent ...
// Pop EventList[index], and return EventList[index]
//
// EventList gets shifted on Pop
func PopAndGetEvent(index int, el *[]Event) (event Event) {
	event, (*el) = (*el)[index], append((*el)[:index], (*el)[index+1:]...)
	return event
}
