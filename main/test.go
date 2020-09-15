package main

import (
	"fmt"
	"github.com/accnameowl/owlevent"
	"time"
)

func main() {

	// chOnSomething := make(chan bool)
	// chOnElse := make(chan bool)

	sEventChannel := make(chan bool)
	eEventChannel := make(chan bool)
	eventHandler := new(owlevent.EventHandle)

	// add multiple events to eventHandler
	//		-> SOMETHING
	eventHandler.OnEvent("ON_SOMETHING", func() {
		fmt.Println("ON_SOMETHING: 1")
	})
	eventHandler.OnEvent("ON_SOMETHING", func() {
		fmt.Println("ON_SOMETHING: 2")
	})
	eventHandler.OnEvent("ON_SOMETHING", func() {
		fmt.Println("ON_SOMETHING: 3")
	})
	eventHandler.OnEvent("ON_SOMETHING", func() {
		fmt.Println("ON_SOMETHING: 4")
	})
	eventHandler.OnEvent("ELSE", func() {
		fmt.Println("ELSE: 1")
	})
	eventHandler.OnEvent("ELSE", func() {
		fmt.Println("ELSE: 2")
	})
	eventHandler.OnEvent("ELSE", func() {
		fmt.Println("ELSE: 3")
	})
	eventHandler.OnEvent("ELSE", func() {
		fmt.Println("ELSE: 4")
	})

	go eventHandler.Start(sEventChannel, "ON_SOMETHING")
	time.Sleep(time.Millisecond * 100)
	go eventHandler.Start(eEventChannel, "ELSE")

	time.Sleep(time.Second * 2)
}
