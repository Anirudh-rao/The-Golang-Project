package channels

import "fmt"

func ChannelsInGos() {
	// This will make a new channel
	messages := make(chan string)

	go func() {
		// <- this will send a message to the channel
		messages <- "ping"
	}()

	// This will recive the channel message
	msg := <-messages
	fmt.Println(msg)
}
