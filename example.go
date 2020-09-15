package owlevent

import (
	"fmt"
)

var (
	// Traffic ...
	// an example variable to handle a set of events.
	Traffic EventHandle
)

// TestExample ...
func TestExample() {
	// Attach events to Traffic
	Traffic.OnEvent("RunBy", func() {
		fmt.Println("A car ran by!")
	})
	Traffic.OnEvent("Parked", func() {
		fmt.Println("A car just parked!")
	})

}

// AddTraffic ...
func AddTraffic(eh *EventHandle) {
	CarOne := new(Car)

	// go-routine listeners
	didCarRunBy := make(chan bool)
	didCarPark := make(chan bool)

	go CarOne.RunBy(didCarRunBy)
	go CarOne.Park(didCarPark)

	if <-didCarRunBy {
		fmt.Println("The car running by was confirmed!")
	}
	if <-didCarPark {
		fmt.Println("The car parking was confirmed!")
	}
}

// ACarDoes ...
type ACarDoes interface {
	Park()
	RunBy()
}

// Car ...
type Car struct {
}

// RunBy ...
func (c *Car) RunBy(ch chan bool) {
	// Calls for Traffic event: RunBy
	// Outputs: A car ran by!
	Traffic.Call(ch, "RunBy")
}

// Park ...
func (c *Car) Park(ch chan bool) {
	// Calls for Traffic event: Parked
	// Output: A car just parked!
	Traffic.Call(ch, "Parked")
}
