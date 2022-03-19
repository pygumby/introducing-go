package main

import (
	"fmt"
	"time"
)

// Now `pinger` is only allowed to send to `c`:
func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

// Now `printer` is only allowed to receive from `c`:
func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
