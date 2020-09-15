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
func RemoveEventByIndex(index int, el *[]Event) (err error) {
	if (index >= 0) || len((*el)) >= 0 {
		(*el) = append((*el)[:index], (*el)[index+1:]...)
	} else {
		err = errors.New("length of argument \"el *[]Event\" not established")
		debug.PrintStack()
	}
	return err
}

// AppendEvents ...
func AppendEvents(move *[]Event, to *[]Event) (err error) {
	if len((*move)) >= 0 {
		if len((*to)) >= 0 {
			(*to) = append((*to), (*move)...)
		} else {
			err = errors.New("Could not establish Event slice to append Event slice to")
			debug.PrintStack()
		}
	} else {
		err = errors.New("Could not establish Event slice to append")
		debug.PrintStack()
	}
	return err
}

// Push ...
func Push(push *Event, to *[]Event) (err error) {
	if len((*to)) >= 0 {
		(*to) = append((*to), *push)
	} else {
		err = errors.New("Could not establish Event slice to push Event to")
		debug.PrintStack()
	}
	return err
}

// PopAndGetEvent ...
// Pop EventList[index], and return EventList[index]
//
// EventList gets shifted on Pop
func PopAndGetEvent(index int, el *[]Event) (event Event, err error) {
	if index < len((*el)) {
		event, (*el) = (*el)[index], append((*el)[:index], (*el)[index+1:]...)
	} else {
		err = errors.New("index out of bounds of []Event argument")
		debug.PrintStack()
	}
	return event, err
}
