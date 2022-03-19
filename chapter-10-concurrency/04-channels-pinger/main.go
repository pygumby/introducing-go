package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		// Alternatively:
		// fmt.Println(<-c)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	// Using a channel like this synchronizes the two goroutines.
	// When `pinger` attempts to send a message on the channel, it will wait until `printer is
	// ready to receive the message (this is known as blocking).
	go pinger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
