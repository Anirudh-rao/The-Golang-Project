package channels

import "fmt"

func ChannelBuffering() {
	message := make(chan string, 2)
	message <- "buffered"
	message <- "channel"

	fmt.Println(<-message)
	fmt.Println(<-message)

}
