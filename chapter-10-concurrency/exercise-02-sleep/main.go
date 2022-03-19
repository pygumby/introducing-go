package main

import (
	"fmt"
	"time"
)

func Sleep(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	for {
		fmt.Println("Ping")
		Sleep(time.Second * 2)
	}
}
