package main

import (
	"fmt"
	"time"
)

func main() {

	channel1 := make(chan string)
	channel2 := make(chan string)
	channel3 := make(chan string)

	my_channel := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		channel3 <- "three"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		channel1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		channel2 <- "two"
	}()

	go func() {
		my_channel <- 2
	}()

	for i := 0; i < 4; i++ {
		select {
		case msg1 := <-channel1:
			fmt.Println("received", msg1)
		case msg2 := <-channel2:
			fmt.Println("received", msg2)
		case msg3 := <-my_channel:
			fmt.Println("received", msg3)
		case msg4 := <-channel3:
			fmt.Println("received", msg4)
		}
	}
}
