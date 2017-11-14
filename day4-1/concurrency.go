package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)
	go func() {
		for {
			messages <- fmt.Sprintf("Time now %s\n", time.Now())
			time.Sleep(100 * time.Millisecond)
		}
	}()
	for {
		msg := <-messages
		fmt.Println(msg)
	}
}
