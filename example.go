package owlevent

import (
	"fmt"
)

var (
	// Traffic ...
	// an example variable to handle a set of events.
	Traffic EventHandle

	// ChTraffic ...
	// go channel which runs Traffic events concurrently
	ChTraffic chan bool
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
	CarOne.RunBy()
	CarOne.Park()
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
func (c *Car) RunBy() {
	// Calls for Traffic event: RunBy
	// Outputs: A car ran by!
	Traffic.Call(ChTraffic, "RunBy")
}

// Park ...
func (c *Car) Park() {
	// Calls for Traffic event: Parked
	// Output: A car just parked!
	Traffic.Call(ChTraffic, "Parked")
}
