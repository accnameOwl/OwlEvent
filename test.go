package main

import (
	"fmt"
	"github.com/accnameowl/OwlEvent"
)

func main() {
	mainEventChannel := make(chan error)
	mainEventHandle := &EventHandle{
		Looping: true
	}
	go mainEventChannel <- mainEventHandle.Start()

	mainEventChannel.OnEvent("something", func() error {
		fmt.println("\"something\" was called...")
	}, async = true)
	mainEventChanne.OnEvent("somethingelse", func() error{
		fmt.println("\"somethingelse"\" was called...")
	})
	close(mainEventChannel)
}