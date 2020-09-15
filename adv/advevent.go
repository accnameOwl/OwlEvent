package adv

import (
	"github.com/accnameowl/owlevent"
	"sync"
)

// TODO : Impl Mutex lock... Examples:
//		Mutex lock/Unlock:
// 		-> https://tour.golang.org/concurrency/9
//		Defer panic and recover, instead of concurrent locking:
//		-> https://blog.golang.org/defer-panic-and-recover

// SafeAsyncEvents ...
type SafeAsyncEvents struct {
	events []Event
	mux    sync.Mutex
}

// RemoveEventByIndex ...
// -> RemoveEventByIndex(indexToRemove, &EventList)
func (sae *SafeAsyncEvents) RemoveEventByIndex(index int) {
	sae.mux.Lock()
	defer sae.mux.Unlock()
	sae.events = append(sae.events[:index], sae.events[index+1:]...)
}

// AppendEvents ...
func (sae *SafeAsyncEvents) AppendEvents(move *[]Event) {
	sae.mux.Lock()
	defer sae.mux.Unlock()
	sae.events = append(sae.events, (*move)...)
}

// Push ...
func (sae *SafeAsyncEvents) Push(push *Event) {
	sae.mux.Lock()
	defer sae.mux.Unlock()
	sae.events = append(sae.events, *push)
}

// PopAndGet ...
// Pop EventList[index], and return EventList[index]
//
// EventList gets shifted on Pop
func (sae *SafeAsyncEvents) PopAndGet(index int, el *[]Event) (event Event) {
	sae.mux.Lock()
	defer sae.mux.Unlock()
	event, sae.events = sae.events[index], append(sae.events[:index], sae.events[index+1:]...)
	return event
}
